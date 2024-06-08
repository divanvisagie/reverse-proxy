package main

import (
	"log"
	"os"
)

// declare args struct to contain args
type Args struct {
	help   bool
	port   string
	target string
	path   string
}

// create associated function on args struct
func (a *Args) ParseArgs() {
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "-h" || os.Args[i] == "--help" {
			a.help = true
			return
		}
	}

	//parse port
	if len(os.Args) < 2 {
		log.Fatalln("You need to pass in port and target parameters: reverse-proxy :8080 http:\\\\example.com")
	}

	a.port = os.Args[1]
	a.target = os.Args[2]
	a.path = "/"
}
