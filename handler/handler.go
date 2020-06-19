package handler

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"time"
)

type MiddleWares func(ctx *fasthttp.RequestCtx) error

type ResponseBody struct {
	Code    int
	Message string
	Ctx     *fasthttp.RequestCtx `json:"-"` //尝试将ctx写到struct里面
}

/**
非法请求处理
*/
func BadRequest(ctx *fasthttp.RequestCtx) {

}

/**
路由请求处理
*/
func Do(middleWares ...MiddleWares) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		if len(middleWares) > 0 {
			for _, md := range middleWares {
				if err := md(ctx); err != nil {

					// 直接ctx返回
					ctx.Response.Reset()
					ctx.Response.Header.SetStatusCode(fasthttp.StatusUnprocessableEntity)

					//将ctx写到struct里
					res := ResponseBody{422, err.Error(), ctx}
					js, _ := json.Marshal(res)

					res.Ctx.Response.SetBody(js)
					res.Ctx.Response.Header.SetContentType("application/json;charset=UTF-8")
					res.Ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")

					// ctx响应以后，继续处理的逻辑
					go func() {
						time.Sleep(10 * time.Second)
						fmt.Println("hello i'm coming...")
					}()

					// ctx响应参数设置
					//ctx.Response.SetBody(js)
					//ctx.Response.Header.SetContentType("application/json;charset=UTF-8")
					//ctx.Response.Header.Set("Connection", "keep-alive")
					//ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")

					//ctx.Write([]byte("干的漂亮"))

					//return
				}
			}
		}
	})
}
