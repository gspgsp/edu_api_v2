package edu

import (
	"github.com/valyala/fasthttp"
	"fmt"
	"github.com/iGoogle-ink/gopay"
)

/**
生成预订单
 */
func GenerateOrder(ctx *fasthttp.RequestCtx) {
	fmt.Println("这是是一个测试"+gopay.Version)
}
