package main

import (
	"io/ioutil"
	"os"
	"strings"
)

func FileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func ListDirAll(dirPth, suffix string) ([]os.FileInfo, error) {
	suffix = strings.ToUpper(suffix)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	return dir, nil
}
