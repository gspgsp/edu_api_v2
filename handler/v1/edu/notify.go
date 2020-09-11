package edu

import (
	"github.com/valyala/fasthttp"
	"log"
)

func NotifyUrl(ctx *fasthttp.RequestCtx) {
	log.Printf("the notify rquest body is:%s", string(ctx.PostArgs().Peek("trade_no")))
}
