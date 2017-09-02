package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func setupReverseProxy(path string, port string, target string) {
	u, _ := url.Parse(target)
	http.Handle(path, httputil.NewSingleHostReverseProxy(u))

	http.ListenAndServe(port, nil)
}

func main() {
	path := "/"
	port := ":8080"
	target := "http://www.google.com"

	args := os.Args[1:]
	if len(args) < 1 {
		errorMessage := "You need to pass in port and target parameters: reverse-proxy :8080 http:\\\\example.com"
		log.Fatalln(errorMessage)
	}

	if len(args) == 1 {
		errorMessage := "You need to provide a target for your proxy"
		log.Fatalln(errorMessage)
	}

	if len(args) == 3 {
		path = args[2]
	}

	port = args[0]
	target = args[1]

	fmt.Printf("Proxying calls to %s%s to %s", port, path, target)

	setupReverseProxy(path, port, target)
}
