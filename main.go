package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Proxy struct {
	url    string
	target string
}

func readFile() []Proxy {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatalln(err.Error())
	}

	jsonMap := make(map[string]string)

	unmarshalError := json.Unmarshal(file, &jsonMap)
	if unmarshalError != nil {
		log.Fatalln(unmarshalError.Error())
	}

	var proxies = make([]Proxy, len(jsonMap))

	i := 0
	for k, v := range jsonMap {
		proxies[i] = Proxy{string(k), string(v)}
		i++
	}

	return proxies
}

func main() {
	proxies := readFile()

	for _, proxy := range proxies {
		fmt.Printf("url: %s target: %s \n", proxy.url, proxy.target)
	}
}
