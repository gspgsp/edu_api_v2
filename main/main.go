package main

import (
	"edu_api_v2/route"
	_ "github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"log"
)

func main() {

	server := &fasthttp.Server{
		Handler:          route.Handler(),
		Name:             "API_SERVER",
		DisableKeepalive: false,
		LogAllErrors:     true,
	}

	if err := server.ListenAndServe("127.0.0.1:80"); err != nil {
		log.Fatalln(err)
	}
}
