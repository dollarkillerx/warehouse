package utils

import (
	"encoding/json"
	"net/http"
)

type H map[string]interface{}

func JSON(w http.ResponseWriter, code int, r interface{}) error {
	w.WriteHeader(code)
	marshal, err := json.Marshal(r)
	if err != nil {
		return err
	}

	_, err = w.Write(marshal)
	return err
}

func String(w http.ResponseWriter, code int, r string) error {
	w.WriteHeader(code)
	_, err := w.Write([]byte(r))
	return err
}
