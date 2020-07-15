package route

import (
	"edu_api_v2/handler"
	edu2 "edu_api_v2/handler/v1/edu"
	"edu_api_v2/middleware/edu"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func Handler() fasthttp.RequestHandler {
	router := fasthttprouter.New()
	router.NotFound = handler.BadRequest
	router.MethodNotAllowed = handler.BadRequest
	router.GET("/v1/slides", handler.Do(edu2.SlidesList, edu.CheckParam))

	return router.Handler
}