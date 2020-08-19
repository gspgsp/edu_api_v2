package edu

import (
	"github.com/valyala/fasthttp"
	"fmt"
	"edu_api_v2/connect"
	"log"
	"edu_api_v2/models"
	"edu_api_v2/response"
)

func GetCourseReview(ctx *fasthttp.RequestCtx) {
	var id = fmt.Sprintf("%s", ctx.QueryArgs().Peek("id"))

	rows, err := connect.Db.Queryx("select c.anonymous, c.rating, c.review, c.reviewed_at, c.reply, u.nickname from h_user_course as c left join h_users as u on c.user_id = u.id where course_id = " + id + " and c.status = 1 and c.reviewed = 1")

	if err != nil {
		log.Printf("查询错误:%s", err.Error())
		return
	}

	if rows == nil {
		log.Print("数据结果为空")
		return
	}

	var (
		review  models.UserReview
		reviews []models.UserReview
	)

	for rows.Next() {
		err := rows.StructScan(&review)
		if err != nil {
			log.Printf("数据结构化错误:%s", err.Error())
			return
		}

		reviews = append(reviews, review)
	}

	response.New(ctx).SetData(reviews).JsonReturn()
}
