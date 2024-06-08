package main

import (
	"fmt"
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
	args := Args{}
	args.ParseArgs()

	if args.help {
		fmt.Println("Usage: reverse-proxy <port> <target>")
		fmt.Println("Example: reverse-proxy :8080 http://www.google.com")
		os.Exit(0)
	}

	fmt.Printf("Proxying calls to %s%s to %s", args.port, args.path, args.target)

	setupReverseProxy(args.path, args.port, args.target)
}
