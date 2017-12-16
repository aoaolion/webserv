package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	Title = "<h1>webserv @aoaolion</h1><p><h3>https://github.com/aoaolion/webserv</h3></p>"
)

func FileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func closeServer(w http.ResponseWriter, r *http.Request) {
	stop <- "api"
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(Title))
	w.Write([]byte("<a href='upload'>>> upload</a><br>"))
	w.Write([]byte("<a href='download'><< download</a><br>"))
}

func upload(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("upload_file")

	if err != nil {
		w.Write([]byte(Title))
		w.Write([]byte("<div>"))
		w.Write([]byte("<form action='http://127.0.0.1:8080/upload' method='post' enctype='multipart/form-data'>"))
		w.Write([]byte("	<p><input type='file' name='upload_file'></p>"))
		w.Write([]byte("	<input type='submit' value='upload' />"))
		w.Write([]byte("</form>"))
		w.Write([]byte("</div>"))
		return
	}

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		w.Write([]byte("<head><meta http-equiv=refresh content='2;url=/upload'></head>"))
		w.Write([]byte(Title))
		w.Write([]byte("upload error, auto redirect"))
		return
	}

	if strings.Contains(header.Filename, "../") {
		w.Write([]byte("<head><meta http-equiv=refresh content='2;url=/upload'></head>"))
		w.Write([]byte(Title))
		w.Write([]byte("upload error, auto redirect"))
		return
	}

	savePath := fmt.Sprintf("%s/%s", *fileRoot, header.Filename)
	log.Printf("save file: %s", savePath)

	if FileExist(savePath) {
		w.Write([]byte(Title))
		w.Write([]byte("<head><meta http-equiv=refresh content='2;url=/upload'></head>"))
		w.Write([]byte("file has exist, auto redirect"))
		return
	}
	ioutil.WriteFile(savePath, buf, 0666)
	w.Write([]byte(Title))
	w.Write([]byte("<head><meta http-equiv=refresh content='2;url=/download'></head>"))
	w.Write([]byte("upload success, auto redirect"))
	return
}
