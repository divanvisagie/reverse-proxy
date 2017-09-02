# reverse-proxy
A reverse proxy that reads a json file and does it's thing, and only that... ever


## Usage

To proxy requests to port 1337 to http://www.example.com

```sh
reverse-proxy :1337 http://www.example.com
```

If you want to proxy a particular path, like /api

```sh
reverse-proxy :1337 http://www.example.com /api
```