package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/vmihailenco/msgpack/v5"
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

func MsgPack(w http.ResponseWriter, code int, r interface{}) error {
	w.WriteHeader(code)
	marshal, err := msgpack.Marshal(r)
	if err != nil {
		return err
	}

	_, err = w.Write(marshal)
	return err
}

func FromMsgPack(r *http.Request, b interface{}) error {
	if r.Body != nil {
		defer r.Body.Close()
	}

	all, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	fmt.Println(len(all))
	err = msgpack.Unmarshal(all, b)
	if err != nil {
		return err
	}

	return err
}
