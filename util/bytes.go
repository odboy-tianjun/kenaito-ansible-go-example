package util

import (
	"io"
	"io/ioutil"
)

func ReaderToString(reader io.Reader) (string, error) {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
