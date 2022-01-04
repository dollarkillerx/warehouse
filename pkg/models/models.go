package models

type MetaData struct {
	BucketName    string                 `json:"bucket_name"`    // bucket name
	ObjectName    string                 `json:"object_name"`    // object name
	MD5Base64     string                 `json:"md5_base64"`     // md5
	ContentLength int64                  `json:"content_length"` // content length
	CreateTime    int64                  `json:"create_time"`    // create time
	Expires       int64                  `json:"expires"`        // 过期时间
	ContentType   string                 `json:"content_type"`   // content type
	MetaData      map[string]interface{} `json:"meta_data"`      // metadata
}

type Object struct {
	MetaData
	Data []byte `json:"data"`
}
