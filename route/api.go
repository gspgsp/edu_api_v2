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
	router.GET("/api/v1/slides", handler.Do(edu2.SlidesList))
	router.GET("/api/v1/packages", handler.Do(edu2.PackageList))
	router.GET("/api/v1/boutiques", handler.Do(edu2.BoutiqueList))
	router.GET("/api/v1/frees", handler.Do(edu2.FreeList))
	router.GET("/api/v1/course-detail", handler.Do(edu2.GetCourseDetail, edu.CourseDetailCheckParam))
	router.GET("/api/v1/course-user", handler.Do(edu2.GetCourseUser, []handler.MiddleWares{edu.CourseDetailCheckParam, edu.CourseUserCheckParam}...))

	return router.Handler
}
