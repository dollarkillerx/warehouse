package server

import (
	"github.com/dollarkillerx/warehouse/internal/config"
	"github.com/dollarkillerx/warehouse/pkg/models"
	"github.com/dollarkillerx/warehouse/pkg/utils"
	"github.com/vmihailenco/msgpack/v5"

	"io/ioutil"
	"net/http"
	"os"
	"path"
)

// ApiPutObject 当前设计只适合小文件
func (s *Server) ApiPutObject(w http.ResponseWriter, r *http.Request) {
	var obj models.Object
	err := utils.FromMsgPack(r, &obj)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	if obj.BucketName == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	metaData, err := msgpack.Marshal(obj.MetaData)
	if err != nil {
		utils.Logger.Error(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	metaDataPath := path.Join(config.GetStoragePath(), obj.BucketName, "metadata", obj.ObjectName)
	dataPath := path.Join(config.GetStoragePath(), obj.BucketName, "data", obj.ObjectName)

	err = utils.MakeDir(metaDataPath)
	if err != nil {
		utils.Logger.Error(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	err = utils.MakeDir(dataPath)
	if err != nil {
		utils.Logger.Error(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	err = ioutil.WriteFile(metaDataPath, metaData, 00600)
	if err != nil {
		utils.Logger.Error(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	err = ioutil.WriteFile(dataPath, obj.Data, 00600)
	if err != nil {
		utils.Logger.Error(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	utils.String(w, 200, "success")
}

// Download 当前设计只适合小文件
func (s *Server) Download(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Query().Get("file")
	if file == "" {
		http.Error(w, http.StatusText(404), 404)
		return
	}

	bucket := r.URL.Query().Get("bucket")
	if bucket == "" {
		http.Error(w, http.StatusText(404), 404)
		return
	}

	dataPath := path.Join(config.GetStoragePath(), bucket, "data", file)
	readFile, err := ioutil.ReadFile(dataPath)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}

	metaDataPath := path.Join(config.GetStoragePath(), bucket, "metadata", file)
	metaData, err := ioutil.ReadFile(metaDataPath)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	var metadata models.MetaData
	err = msgpack.Unmarshal(metaData, &metadata)
	if err != nil {
		utils.Logger.Error(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", metadata.ContentType)
	w.Header().Set("Content-Disposition", "attachment; filename="+dataPath)
	w.Write(readFile)
}

// ApiDelete ApiDelete
func (s *Server) ApiDelete(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Query().Get("file")
	if file == "" {
		http.Error(w, http.StatusText(404), 404)
		return
	}

	bucket := r.URL.Query().Get("bucket")
	if bucket == "" {
		http.Error(w, http.StatusText(404), 404)
		return
	}

	dataPath := path.Join(config.GetStoragePath(), bucket, "data", file)
	metaDataPath := path.Join(config.GetStoragePath(), bucket, "metadata", file)

	os.Remove(dataPath)
	os.Remove(metaDataPath)

	utils.String(w, 200, "success")
}

// ApiGetObject ApiGetObject
func (s *Server) ApiGetObject(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Query().Get("file")
	if file == "" {
		http.Error(w, http.StatusText(404), 404)
		return
	}

	bucket := r.URL.Query().Get("bucket")
	if bucket == "" {
		http.Error(w, http.StatusText(404), 404)
		return
	}

	dataPath := path.Join(config.GetStoragePath(), bucket, "data", file)
	readFile, err := ioutil.ReadFile(dataPath)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}

	metaDataPath := path.Join(config.GetStoragePath(), bucket, "metadata", file)
	metaData, err := ioutil.ReadFile(metaDataPath)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	var metadata models.MetaData
	err = msgpack.Unmarshal(metaData, &metadata)
	if err != nil {
		utils.Logger.Error(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	utils.MsgPack(w, 200, models.Object{
		MetaData: metadata,
		Data:     readFile,
	})
}

// ApiRemoveBucket ApiRemoveBucket
func (s *Server) ApiRemoveBucket(w http.ResponseWriter, r *http.Request) {
	bucket := r.URL.Query().Get("bucket")
	if bucket == "" {
		http.Error(w, http.StatusText(404), 404)
		return
	}

	bucketPath := path.Join(config.GetStoragePath(), bucket)

	os.Remove(bucketPath)

	utils.String(w, 200, "success")
}
