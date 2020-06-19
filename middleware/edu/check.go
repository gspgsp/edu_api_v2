package edu

import (
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
)

func CheckParam(ctx *fasthttp.RequestCtx) error {
	return errors.New("请登录先")
}
