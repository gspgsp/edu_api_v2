package edu

import (
	"github.com/valyala/fasthttp"
	"edu_api_v2/connect"
	"log"
	"edu_api_v2/models"
	"edu_api_v2/response"
)

func BoutiqueList(ctx *fasthttp.RequestCtx) {

	rows, err := connect.Db.Queryx("select id, title, subtitle, difficulty_level, (ifnull(learn_num, 0) + ifnull(buy_num, 0)) as learn_count , price, discount, discount_end_at, vip_level, vip_price, cover_picture from h_edu_courses where status = 'published' order by updated_at desc, sort desc limit 5")

	if err != nil {
		log.Printf("查询错误:%s" + err.Error())
		return
	}

	boutique := models.Boutique{}
	var boutiques []models.Boutique

	if rows == nil {
		log.Print("数据结果为空")
		return
	}

	for rows.Next() {
		err := rows.StructScan(&boutique)

		if err != nil {
			log.Printf("数据结构化错误:%s", err.Error())
		} else {
			boutiques = append(boutiques, boutique)
		}
	}

	response.New(ctx).SetData(boutiques).JsonReturn()

}
