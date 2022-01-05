package warehouse_sdk

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/dollarkillerx/urllib"
	"github.com/dollarkillerx/warehouse/pkg/models"
	"github.com/pkg/errors"
	"github.com/vmihailenco/msgpack/v5"
)

type WareHouseSdk struct {
	addr      string
	accessKey string
	secretKey string
	timeout   time.Duration
}

func New(addr string, accessKey string, secretKey string, timeout time.Duration) *WareHouseSdk {
	if timeout < time.Second {
		timeout = time.Second * 10
	}
	return &WareHouseSdk{
		addr:      addr,
		accessKey: accessKey,
		secretKey: secretKey,
		timeout:   timeout,
	}
}

func (w *WareHouseSdk) auth(r *urllib.Urllib) *urllib.Urllib {
	return r.SetHeader("AccessKey", w.accessKey).SetHeader("SecretKey", w.secretKey).SetTimeout(w.timeout)
}

func (w *WareHouseSdk) Ping() error {
	code, rp, err := w.auth(urllib.Post(fmt.Sprintf("%s/v1/auth", w.addr))).ByteOriginal()
	if err != nil {
		return err
	}

	if code != 200 {
		return errors.New(string(rp))
	}

	return nil
}

func (w *WareHouseSdk) PutObject(bucketName, objectName string, data []byte, metaData *models.MetaData) error {
	var createTime = time.Now().Unix()
	var contentType = TypeByExtension(objectName)
	var mt map[string]interface{}
	var expires int64
	if metaData != nil {
		if metaData.CreateTime != 0 {
			createTime = metaData.CreateTime
		}
		if metaData.MetaData != nil {
			mt = metaData.MetaData
		}
		if metaData.Expires != 0 {
			expires = metaData.Expires
		}
	}

	obj := models.Object{
		MetaData: models.MetaData{
			BucketName:    bucketName,
			ObjectName:    objectName,
			MD5Base64:     md5Encode(data),
			ContentLength: int64(len(data)),
			CreateTime:    createTime,
			ContentType:   contentType,
			MetaData:      mt,
			Expires:       expires,
		},
		Data: data,
	}

	marshal, err := msgpack.Marshal(obj)
	if err != nil {
		return err
	}

	code, rp, err := w.auth(urllib.Post(fmt.Sprintf("%s/v1/put_object", w.addr))).SetBody(marshal).ByteOriginal()
	if err != nil {
		return err
	}

	if code != 200 {
		return errors.New(string(rp))
	}

	return nil
}

func (w *WareHouseSdk) GetObject(bucket, objectName string) (*models.Object, error) {
	var obj models.Object
	code, rp, err := w.auth(urllib.Get(fmt.Sprintf("%s/v1/get_object", w.addr))).Queries("file", objectName).Queries("bucket", bucket).ByteOriginal()
	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, errors.New(string(rp))
	}

	err = msgpack.Unmarshal(rp, &obj)
	if err != nil {
		return nil, err
	}

	return &obj, nil
}

func (w *WareHouseSdk) DelObject(bucket, objectName string) error {
	code, rp, err := w.auth(urllib.Post(fmt.Sprintf("%s/v1/del_object", w.addr))).Queries("file", objectName).Queries("bucket", bucket).ByteOriginal()
	if err != nil {
		return err
	}

	if code != 200 {
		return errors.New(string(rp))
	}
	return nil
}

func (w *WareHouseSdk) RemoveBucket(bucket string) error {
	code, rp, err := w.auth(urllib.Post(fmt.Sprintf("%s/v1/remove_bucket", w.addr))).Queries("bucket", bucket).ByteOriginal()
	if err != nil {
		return err
	}

	if code != 200 {
		return errors.New(string(rp))
	}
	return nil
}

func md5Encode(data []byte) string {
	md5Ctx := md5.New()
	md5Ctx.Write(data)
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
