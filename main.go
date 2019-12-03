package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var version = "1.0.0"
var bind = flag.String("bind", "0.0.0.0:8000", "Listen address for HTTP server, e.g. 127.0.0.1:8000")
var name = flag.String("name", "simple-http-server", "The name of server, e.g. my server")

func init() {
	flag.Parse()
	log.Printf("Version: %s", version)
	log.Printf("%s server runing at: %s \n", *name, *bind)
	log.Println("use Ctrl + C to stop server.")
}

func main() {
	http.HandleFunc("/", rootPathHandler)
	log.Fatal(http.ListenAndServe(*bind, nil))
}

func rootPathHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s handle %s %s%s start!\n", *name, r.Method, r.Host, r.RequestURI)
	fmt.Fprintf(w, "Server name: %s %s\n\n", *name, version)

	fmt.Fprintf(w, "Method: %s \n", r.Method)
	fmt.Fprintf(w, "Host: %s \n", r.Host)
	fmt.Fprintf(w, "URL: %s \n", r.URL)
	fmt.Fprintf(w, "RequestURI: %s \n", r.RequestURI)
	fmt.Fprintf(w, "Proto: %s \n", r.Proto)
	fmt.Fprintf(w, "Referer: %s \n", r.Referer())
	fmt.Fprintf(w, "Header: %s \n", r.Header)
	fmt.Fprintf(w, "Form: %s \n", r.Form)
	fmt.Fprintf(w, "PostForm: %s \n\n", r.PostForm)

	fmt.Fprintf(w, "RemoteAddr: %s \n", r.RemoteAddr)
	fmt.Fprintf(w, "UserAgent: %s \n", r.UserAgent())

	log.Printf("handle %s %s%s success!\n", r.Method, r.Host, r.RequestURI)
}
