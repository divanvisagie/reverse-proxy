# reverse-proxy

[![Build Status](https://travis-ci.org/divanvisagie/reverse-proxy.svg?branch=master)](https://travis-ci.org/divanvisagie/reverse-proxy)

A simple command line reverse proxy

## Usage

To proxy requests to port 1337 to http://www.example.com

```sh
reverse-proxy :1337 http://www.example.com
```

If you want to proxy a particular path, like /api

```sh
reverse-proxy :1337 http://www.example.com /api
```