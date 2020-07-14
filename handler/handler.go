package handler

import (
	"edu_api_v2/response"
	"github.com/valyala/fasthttp"
)

type MiddleWares func(ctx *fasthttp.RequestCtx) error

/**
非法请求处理
*/
func BadRequest(ctx *fasthttp.RequestCtx) {

}

/**
路由请求处理
*/
func Do(next fasthttp.RequestHandler, middleWares ...MiddleWares) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		if len(middleWares) > 0 {
			for _, md := range middleWares {
				if err := md(ctx); err != nil {
					response.New(ctx).SetMessage(err).JsonReturn()
					return
				}
			}
		}
		next(ctx)
	})
}
