package edu

import (
	"github.com/valyala/fasthttp"
	"edu_api_v2/connect"
	"log"
	"edu_api_v2/models"
	"edu_api_v2/response"
)

func FreeList(ctx *fasthttp.RequestCtx) {

	rows, err := connect.Db.Queryx("select id, title, subtitle, difficulty_level, (ifnull(learn_num, 0) + ifnull(buy_num, 0)) as learn_count, cover_picture from h_edu_courses where status = 'published' order by updated_at desc, sort desc limit 4")
	if err != nil {
		log.Printf("查询错误:%s", err.Error())
		return
	}

	free := models.Free{}
	var frees []models.Free
	if rows == nil {
		log.Print("数据结果为空")
		return
	}

	for rows.Next() {
		err := rows.StructScan(&free)
		if err != nil {
			log.Printf("数据结构化错误:%s", err.Error())
			return
		} else {
			frees = append(frees, free)
		}
	}

	response.New(ctx).SetData(frees).JsonReturn()
}
