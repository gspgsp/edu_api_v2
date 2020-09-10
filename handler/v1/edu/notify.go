package edu

import (
	"github.com/valyala/fasthttp"
	"log"
)

func NotifyUrl(ctx *fasthttp.RequestCtx) {
	log.Printf("the rquest body is:%s", ctx.UserAgent())
}
