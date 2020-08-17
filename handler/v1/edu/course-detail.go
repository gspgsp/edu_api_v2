package edu

import (
	"github.com/valyala/fasthttp"
	"edu_api_v2/connect"
	"log"
	"edu_api_v2/models"
	"edu_api_v2/response"
	"fmt"
)

/**
获取课程详情
 */
func GetCourseDetail(ctx *fasthttp.RequestCtx) {
	var id = fmt.Sprintf("%s", ctx.QueryArgs().Peek("id"))

	rows, err := connect.Db.Queryx("select id, uuid, type, title, subtitle, price, vip_price, discount, discount_end_at, ifnull(learn_num, 0) + ifnull(buy_num, 0) as learn_count, length, rating, practical_rating, popular_rating, logic_rating, goals, audiences, summary from h_edu_courses where id = "+id)

	if err != nil {
		log.Printf("查询错误:%s", err.Error())
		return
	}

	courseDetail := models.CourseDetail{}
	var course_details []models.CourseDetail

	if rows == nil {
		log.Print("数据结果为空")
		return
	}

	for rows.Next() {
		err := rows.StructScan(&courseDetail)
		if err != nil {
			log.Printf("数据结构化错误:%s", err.Error())
			return
		}

		course_details = append(course_details, courseDetail)
	}

	response.New(ctx).SetData(course_details).JsonReturn()
}
