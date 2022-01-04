package utils

import (
	"os"
	"path"
	"strings"
)

func MakeDir(objectName string) error {
	split := strings.Split(objectName, "/")
	split = split[:len(split)-1]

	return os.MkdirAll(path.Join(split...), 00700)
}
