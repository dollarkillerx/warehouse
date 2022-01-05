package warehouse_sdk

import (
	"fmt"
	"io/ioutil"
	"testing"
)

var img = "1EF47AB1E252EFF71AF37D8761B58F8E.jpg"

func TestSdk(t *testing.T) {
	sdk := New("http://127.0.0.1:8187", "123", "345", 0)
	err := sdk.Ping()
	if err != nil {
		panic(err)
	}

	file, err := ioutil.ReadFile(img)
	if err != nil {
		panic(err)
	}
	err = sdk.PutObject("img", img, file, nil)
	if err != nil {
		panic(err)
	}

	object, err := sdk.GetObject("img", img)
	if err != nil {
		panic(err)
	}

	fmt.Println(object.MetaData)
	fmt.Println(len(object.Data))

	//err = sdk.DelObject("img", img)
	//if err != nil {
	//	panic(err)
	//}

	err = sdk.RemoveBucket("img")
	if err != nil {
		panic(err)
	}
}
