package response

import (
	"edu_api_v2/message"
	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

/**
初始化响应对象
*/
func New(ctx *fasthttp.RequestCtx) *response {
	return &response{200, "success", []interface{}{}, ctx}
}

func (r *response) SetMessage(err error) *response {
	r.Message = message.String(err)
	r.Code = message.Code(err)
	return r
}

func (r *response) SetData(data interface{}) *response {
	r.Data = data
	return r
}

/**
api响应对象
*/
type response struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Data    interface{}          `json:"data"`
	Ctx     *fasthttp.RequestCtx `json:"-"`
}
