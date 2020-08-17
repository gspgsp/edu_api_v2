package edu

import (
	"github.com/valyala/fasthttp"
	"strconv"
	"errors"
)

/**
课程详情信息验证
 */
func CourseDetailCheckParam(ctx *fasthttp.RequestCtx) error {
	if ctx.QueryArgs().Has("id") {
		id, err := strconv.Atoi(string(ctx.QueryArgs().Peek("id")))
		if err != nil {
			return errors.New("解析ID错误")
		}

		if id <= 0 {
			return errors.New("ID必须大于0")
		}

		return nil
	}

	return errors.New("ID参数必传")
}
