package edu

import (
	"github.com/valyala/fasthttp"
	"log"
)

func RerurnUrl(ctx *fasthttp.RequestCtx)  {
	log.Printf("the return rquest body is:%s", ctx.UserAgent())
}
