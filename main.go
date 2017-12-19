package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	fileRoot  *string = flag.String("d", "file_root", "root of dir")
	ip        *string = flag.String("h", "", "listen ip")
	port      *int    = flag.Int("p", 8080, "listen port")
	ttl       *int    = flag.Int("t", 600, "time to live in second")
	gUsername *string = flag.String("u", "", "username")
	gPassword *string = flag.String("P", "", "password")
	stop              = make(chan string)
)

func signalListener(stop chan string) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)
	select {
	case <-c:
		log.Println("receive stop signal")
		stop <- "signal"
	}
}

func httpListener(addr string, stop chan string) {
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Println(err)
		stop <- "port conflict"
	}
}

func ttlListener(stop chan string) {
	if *ttl == 0 {
		return
	}
	s := time.After(time.Second * time.Duration(*ttl))
	select {
	case <-s:
		log.Println("ttl is out")
		stop <- "ttl"
	}
}

func main() {
	flag.Parse()
	if !FileExist(*fileRoot) {
		os.MkdirAll(*fileRoot, 0755)
	}

	http.HandleFunc("/upload/", upload)
	http.HandleFunc("/download/", download)
	http.HandleFunc("/play/", play)
	http.HandleFunc("/delete/", del)
	http.HandleFunc("/shutdown/", shutdown)
	http.HandleFunc("/logout/", logout)
	http.HandleFunc("/", download)

	addr := fmt.Sprintf("%s:%d", *ip, *port)
	log.Printf("msg=http server start||ip=%s||port=%d||file_root=%s", *ip, *port, *fileRoot)

	go httpListener(addr, stop)
	go signalListener(stop)
	go ttlListener(stop)
	select {
	case msg := <-stop:
		log.Printf("msg=http server close||cause=%s", msg)
	}
}
