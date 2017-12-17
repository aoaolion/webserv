package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func reqlog(r *http.Request) {
	log.Printf("msg=request||method=%s||url=%s||host=%s", r.Method, r.URL, r.Host)
}

func closeServer(w http.ResponseWriter, r *http.Request) {
	reqlog(r)
	w.Write([]byte(Title))
	w.Write([]byte("webserv close"))
	stop <- "api"
}

func play(w http.ResponseWriter, r *http.Request) {
	reqlog(r)
	fileName := strings.TrimPrefix(r.URL.Path, "/play/")
	filePath := fmt.Sprintf("%s/%s", *fileRoot, fileName)

	if !FileExist(filePath) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("<head><meta http-equiv=refresh content='1;url=/'></head>"))
		w.Write([]byte(Title))
		w.Write([]byte("play error, auto redirect"))
		return
	}

	play := fmt.Sprintf("<body><video src='/download/%s' autoplay='autoplay' type='video/mp4' height='400px' width='100%%' controls='controls'></video></body>", fileName)
	w.Write([]byte(Title))
	w.Write([]byte(play))
}

func download(w http.ResponseWriter, r *http.Request) {
	reqlog(r)
	if strings.Contains(r.URL.String(), "../") {
		w.Write([]byte("<head><meta http-equiv=refresh content='1;url=/'></head>"))
		w.Write([]byte(Title))
		w.Write([]byte("download error, auto redirect"))
		return
	}

	// url without filename, render file list
	if r.URL.String() == "/download/" || r.URL.String() == "/" {
		files, err := ListDirAll(*fileRoot, "")
		if err != nil {
			w.Write([]byte("<head><meta http-equiv=refresh content='1;url=/'></head>"))
			w.Write([]byte(Title))
			w.Write([]byte("download error, auto redirect"))
		}
		w.Write([]byte(Title))
		w.Write([]byte("<table width='100%'>"))
		w.Write([]byte("<thead><th>file</th><th>size</th><th>modify time</th><th>manage</th></thead><tbody>"))
		for _, file := range files {
			w.Write([]byte("<tr>"))
			line := ""
			if strings.HasSuffix(strings.ToLower(file.Name()), ".mov") ||
				strings.HasSuffix(strings.ToLower(file.Name()), ".mp4") {
				line = fmt.Sprintf(`<td><a href='/download/%s'>%s</a></td>
				<td>%d</td> <td>%s</td> <td><a href='/play/%s'>play</a>&nbsp;<a href='/delete/%s'>delete</a></td>`,
					file.Name(), file.Name(), file.Size(), file.ModTime(), file.Name())
			} else {
				line = fmt.Sprintf(`<td><a href='/download/%s'>%s</a></td>
				<td>%d</td> <td>%s</td> <td><a href='/delete/%s'>delete</a></td>`,
					file.Name(), file.Name(), file.Size(), file.ModTime(), file.Name())
			}
			w.Write([]byte(line))
			w.Write([]byte("</tr>"))
		}
		w.Write([]byte("</tbody></table>"))
		return
	}

	fileName := strings.TrimPrefix(r.URL.Path, "/download/")
	filePath := fmt.Sprintf("%s/%s", *fileRoot, fileName)

	if !FileExist(filePath) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("<head><meta http-equiv=refresh content='1;url=/'></head>"))
		w.Write([]byte(Title))
		w.Write([]byte("download error, auto redirect"))
		return
	}
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		w.Write([]byte("<head><meta http-equiv=refresh content='1;url=/'></head>"))
		w.Write([]byte(Title))
		w.Write([]byte("download error, auto redirect"))
		return
	}
	w.Write(buf)
	w.Header().Add("content-disposition", "attachment; filename=\""+fileName+"\"")

}

func del(w http.ResponseWriter, r *http.Request) {
	reqlog(r)
	fileName := strings.TrimPrefix(r.URL.Path, "/delete/")
	if strings.Contains(fileName, "../") {
		w.Write([]byte("<head><meta http-equiv=refresh content='1;url=/'></head>"))
		w.Write([]byte(Title))
		w.Write([]byte("del error, auto redirect"))
		return
	}

	filePath := fmt.Sprintf("%s/%s", *fileRoot, fileName)

	if !FileExist(filePath) {
		w.Write([]byte("<head><meta http-equiv=refresh content='1;url=/'></head>"))
		w.Write([]byte(Title))
		w.Write([]byte("delete error, auto redirect"))
		return
	}
	os.Remove(filePath)
	w.Write([]byte(Title))
	w.Write([]byte("<head><meta http-equiv=refresh content='1;url=/download'></head>"))
	w.Write([]byte("delete success, auto redirect"))
}

func upload(w http.ResponseWriter, r *http.Request) {
	reqlog(r)
	file, header, err := r.FormFile("upload_file")
	if err != nil {
		w.Write([]byte(Title))
		w.Write([]byte("<div>"))
		w.Write([]byte("<form action='/upload/' method='post' enctype='multipart/form-data'>"))
		w.Write([]byte("	<p><input type='file' name='upload_file'></p>"))
		w.Write([]byte("	<input type='submit' value='upload' />"))
		w.Write([]byte("</form>"))
		w.Write([]byte("</div>"))
		return
	}

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		w.Write([]byte("<head><meta http-equiv=refresh content='1;url=/upload'></head>"))
		w.Write([]byte(Title))
		w.Write([]byte("upload error, auto redirect"))
		return
	}

	if strings.Contains(header.Filename, "../") {
		w.Write([]byte("<head><meta http-equiv=refresh content='1;url=/upload'></head>"))
		w.Write([]byte(Title))
		w.Write([]byte("upload error, auto redirect"))
		return
	}

	savePath := fmt.Sprintf("%s/%s", *fileRoot, header.Filename)
	if FileExist(savePath) {
		w.Write([]byte(Title))
		w.Write([]byte("<head><meta http-equiv=refresh content='1;url=/upload'></head>"))
		w.Write([]byte("file has exist, auto redirect"))
		return
	}
	ioutil.WriteFile(savePath, buf, 0666)
	w.Write([]byte(Title))
	w.Write([]byte("<head><meta http-equiv=refresh content='1;url=/download'></head>"))
	w.Write([]byte("upload success, auto redirect"))
	return
}
