package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func Redirect(w http.ResponseWriter, msg string) {
	w.Write([]byte("<html>" + HeadRedirect))
	w.Write([]byte("<body>"))
	w.Write([]byte(Nav))
	w.Write([]byte("<dialog open style='margin-top: 100px;'>" + msg + "</dialog>"))
	w.Write([]byte("</body></html>"))
}

func reqlog(w http.ResponseWriter, r *http.Request) {
	log.Printf("msg=request||method=%s||url=%s||host=%s", r.Method, r.URL, r.Host)
}

func shutdown(w http.ResponseWriter, r *http.Request) {
	reqlog(w, r)
	if !Auth(w, r) {
		return
	}
	Redirect(w, "Webserv shutdown")
	stop <- "api"
}

func play(w http.ResponseWriter, r *http.Request) {
	reqlog(w, r)
	if !Auth(w, r) {
		return
	}
	fileName := strings.TrimPrefix(r.URL.Path, "/play/")
	filePath := fmt.Sprintf("%s/%s", *fileRoot, fileName)

	if !FileExist(filePath) {
		w.WriteHeader(http.StatusNotFound)
		Redirect(w, "Play error")
		return
	}

	w.Write([]byte("<html>" + Head + "<body>"))
	w.Write([]byte(Nav))
	w.Write([]byte("<div style='background-color: white; margin: 10 10px'>"))

	play := fmt.Sprintf("<video src='/download/%s' autoplay='autoplay' type='video/mp4' height='400px' width='100%%' controls='controls'></video></body>", fileName)
	w.Write([]byte(play))
	w.Write([]byte("</div></body></html>"))
}

func download(w http.ResponseWriter, r *http.Request) {
	reqlog(w, r)
	if !Auth(w, r) {
		return
	}
	if strings.Contains(r.URL.String(), "../") {
		Redirect(w, "Download error")
		return
	}

	// url without filename, render file list
	if r.URL.String() == "/download/" || r.URL.String() == "/" {
		files, err := ListDirAll(*fileRoot, "")
		if err != nil {
			Redirect(w, "Download error")
		}

		w.Write([]byte("<html>" + Head + "<body>"))
		w.Write([]byte(Nav))
		w.Write([]byte("<div style='background-color: white; margin: 10 10px'>"))
		w.Write([]byte("<table width='100%'>"))
		w.Write([]byte("<thead><th>File</th><th>Size</th><th>Modify time</th><th>Manage</th></thead><tbody>"))
		for _, file := range files {
			w.Write([]byte("<tr>"))
			line := ""
			if strings.HasSuffix(strings.ToLower(file.Name()), ".mov") ||
				strings.HasSuffix(strings.ToLower(file.Name()), ".mp4") {
				line = fmt.Sprintf(`<td><a href='/download/%s'>%s</a></td>
				<td>%s</td> <td>%s</td> <td style="text-align: center;"><div class='btn'><a href='/play/%s'>play</a></div>&nbsp;<div class='btn'><a href='/download/%s'>down</a></div>&nbsp;<div class='btn'><a href='/delete/%s'>del</a></div></td>`,
					file.Name(), file.Name(), UnitSize(file.Size()), file.ModTime().Format("2006-01-02 15:04:05"), file.Name(), file.Name(), file.Name())
			} else {
				line = fmt.Sprintf(`<td><a href='/download/%s'>%s</a></td>
				<td>%s</td> <td>%s</td> <td style="text-align: center;"><div class='btn'><a href='/download/%s'>down</a></div>&nbsp;<div class='btn'><a href='/delete/%s'>del</a></div></td>`,
					file.Name(), file.Name(), UnitSize(file.Size()), file.ModTime().Format("2006-01-02 15:04:05"), file.Name(), file.Name())
			}
			w.Write([]byte(line))
			w.Write([]byte("</tr>"))
		}
		w.Write([]byte("</tbody></table>"))
		w.Write([]byte("</div></body></html>"))
		return
	}

	fileName := strings.TrimPrefix(r.URL.Path, "/download/")
	filePath := fmt.Sprintf("%s/%s", *fileRoot, fileName)

	if !FileExist(filePath) {
		w.WriteHeader(http.StatusNotFound)
		Redirect(w, "Download error")
		return
	}
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		Redirect(w, "Download error")
		return
	}
	w.Write(buf)
	w.Header().Add("content-disposition", "attachment; filename=\""+fileName+"\"")

}

func del(w http.ResponseWriter, r *http.Request) {
	reqlog(w, r)
	if !Auth(w, r) {
		return
	}
	fileName := strings.TrimPrefix(r.URL.Path, "/delete/")
	if strings.Contains(fileName, "../") {
		Redirect(w, "Delete error")
		return
	}

	filePath := fmt.Sprintf("%s/%s", *fileRoot, fileName)

	if !FileExist(filePath) {
		Redirect(w, "Delete error")
		return
	}
	os.Remove(filePath)
	Redirect(w, "Delete success")
}

func upload(w http.ResponseWriter, r *http.Request) {
	reqlog(w, r)
	file, header, err := r.FormFile("upload_file")
	if err != nil {
		w.Write([]byte("<html>" + Head + "<body>"))
		w.Write([]byte(Nav))

		w.Write([]byte("<dialog open style='margin-top: 100px;'>"))
		w.Write([]byte("<div style='background-color: white;'>"))
		w.Write([]byte("<form action='/upload/' method='post' enctype='multipart/form-data'>"))
		w.Write([]byte("	<p><input type='file' name='upload_file'></p>"))
		w.Write([]byte("	<input type='submit' value='upload' />"))
		w.Write([]byte("</form>"))
		w.Write([]byte("</div></dialog>"))
		w.Write([]byte("</body></html>"))
		return
	}

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		Redirect(w, "Upload error")
		return
	}

	if strings.Contains(header.Filename, "../") {
		Redirect(w, "Upload error")
		return
	}

	savePath := fmt.Sprintf("%s/%s", *fileRoot, header.Filename)
	if FileExist(savePath) {
		Redirect(w, "Error, file has exist")
		return
	}
	err = ioutil.WriteFile(savePath, buf, 0666)
	if err != nil {
		Redirect(w, "Upload error")
		return
	}
	Redirect(w, "Upload success")
	return
}

func logout(w http.ResponseWriter, req *http.Request) {
	if *gUsername == "" {
		Redirect(w, "Not login yet")
		return
	}
	w.Header().Set("WWW-Authenticate", `Basic realm="Dotcoo User Login"`)
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("Logout"))
	return
}

func Auth(w http.ResponseWriter, req *http.Request) bool {
	// always return auth ok, when username is empty
	if *gUsername == "" {
		return true
	}

	auth := req.Header.Get("Authorization")
	if auth == "" {
		w.Header().Set("WWW-Authenticate", `Basic realm="Webserv User Login"`)
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}

	auths := strings.SplitN(auth, " ", 2)
	if len(auths) != 2 {
		w.Header().Set("WWW-Authenticate", `Basic realm="Webserv User Login"`)
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}
	authMethod := auths[0]
	authB64 := auths[1]
	switch authMethod {
	case "Basic":
		authstr, err := base64.StdEncoding.DecodeString(authB64)
		if err != nil {
			w.Header().Set("WWW-Authenticate", `Basic realm="Webserv User Login"`)
			w.WriteHeader(http.StatusUnauthorized)
			return false
		}
		userPwd := strings.SplitN(string(authstr), ":", 2)
		if len(userPwd) != 2 {
			return false
		}
		username := userPwd[0]
		password := userPwd[1]

		if username == *gUsername && password == *gPassword {
			return true
		} else {
			w.Header().Set("WWW-Authenticate", `Basic realm="Dotcoo User Login"`)
			w.WriteHeader(http.StatusUnauthorized)
			io.WriteString(w, "Unauthorized")
			return false
		}
	default:
		w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, "Unauthorized")
		io.WriteString(w, "Unauthorized")
		return false
	}
}
