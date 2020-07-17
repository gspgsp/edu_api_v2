package edu

import (
	"edu_api_v2/connect"
	"edu_api_v2/models"
	"edu_api_v2/response"
	"github.com/valyala/fasthttp"
	"log"
)

var (
	db = connect.Db
)

func SlidesList(ctx *fasthttp.RequestCtx) {
	rows, err := db.Queryx("select id, port, title, url, carousel, ifnull(sort, '0') as sort from h_slides where status = 1 limit 4")
	if err != nil {
		log.Printf("查询错误:%s", err.Error())
		return
	}

	slide := models.Slide{}
	var slides []models.Slide
	if rows == nil {
		log.Print("数据结果为空")
		return
	}

	for rows.Next() {
		err := rows.StructScan(&slide)
		if err != nil {
			log.Printf("数据结构化错误:%s", err.Error())
		} else {
			slides = append(slides, slide)
		}
	}

	response.New(ctx).SetData(slides).JsonReturn()
}
