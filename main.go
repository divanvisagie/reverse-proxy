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

// declare args struct to contain args
type args struct {
	port   string
	target string
	path   string
}

// create associated function on args struct
func (a *args) parseArgs() {
	//parse port
	if len(os.Args) < 2 {
		log.Fatalln("You need to pass in port and target parameters: reverse-proxy :8080 http:\\\\example.com")
	}

	a.port = os.Args[1]
	a.target = os.Args[2]
	a.path = "/"
}

func main() {
	path := "/"
	port := ":8080"
	target := "http://www.google.com"

	args := args{}
	args.parseArgs()

	fmt.Printf("Proxying calls to %s%s to %s", args.port, args.path, args.target)

	setupReverseProxy(path, port, target)
}
