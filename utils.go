package main

import (
	"fmt"
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

func UnitSize(size int64) string {
	ret := ""
	switch {
	case 0 <= size && size < 1024:
		ret = fmt.Sprintf("%d Byte", size)
	case 1024 <= size && size < 1024*1024:
		ret = fmt.Sprintf("%.2f KB", float32(size)/1024)
	case 1024*1024 <= size && size < 1024*1024*1024:
		ret = fmt.Sprintf("%.2f MB", float32(size)/1024/1024)
	case 1024*1024*1024 <= size && size < 1024*1024*1024*1024:
		ret = fmt.Sprintf("%.2f GB", float32(size)/1024/1024/1024)
	}
	return ret
}
