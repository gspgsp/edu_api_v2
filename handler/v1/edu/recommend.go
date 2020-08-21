package edu

import (
	"github.com/valyala/fasthttp"
	"fmt"
	"log"
	"edu_api_v2/connect"
	"edu_api_v2/models"
	"edu_api_v2/response"
	models2 "edu_api/models"
)

type Tag struct {
	Id string `db:"id"json:"id"`
}

/**
推荐课程
 */
func GetRecommend(ctx *fasthttp.RequestCtx) {
	var (
		id    = fmt.Sprintf("%s", ctx.QueryArgs().Peek("id"))
		limit = "5"
	)

	rows, err := connect.Db.Queryx("select t.id from h_tags t where exists (select 1 from h_taggables t_g inner join h_edu_courses c on c.id = t_g.taggable_id where t_g.tag_id = t.id and c.id = " + id + " and c.status = 'published')")

	if err != nil {
		log.Printf("查询错误:%s", err.Error())
		return
	}

	if rows == nil {
		log.Print("数据结果为空")
		return
	}

	var (
		tagId      Tag
		tags       []Tag
		recommend  models.Course
		recommends []models.Course
	)

	for rows.Next() {
		err := rows.StructScan(&tagId)
		if err != nil {
			log.Printf("数据结构化错误:%s", err.Error())
			return
		}

		tags = append(tags, tagId)
	}

	if len(tags) == 0 {
		rows, err := connect.Db.Queryx("select id, type, title, subtitle, difficulty_level, cover_picture from h_edu_courses where type in ('free', 'boutique') and status = 'published' and is_recommended = 1 and id != " + id + " limit " + limit)

		if err != nil {
			log.Printf("查询错误:%s", err.Error())
			return
		}

		if rows == nil {
			log.Print("数据结果为空")
			return
		}

		for rows.Next() {
			err := rows.StructScan(&recommend)
			if err != nil {
				log.Printf("数据结构化错误:%s", err.Error())
				return
			}

			recommends = append(recommends, recommend)
		}
	} else {
		//推荐功能

	}

	log.Printf("id is:%v", tags)
	response.New(ctx).SetData(recommends).JsonReturn()
}

func recommends(tags []Tag, id string) (recommends []models2.Recommend, err error) {

	channel := make(chan models2.Recommend, len(tags))
	isEnd := make(chan bool)
	endNumber := 0

	for _, val := range tags {
		go tagCourses(channel, isEnd, val.Id, id)
	}

GetChannelData:
	for  {
		select {
			case v, ok := <- channel:
				if ok {
					recommends = append(recommends, v)
				}else {
					log.Printf("当前channel")
				}
		case <-isEnd:
			endNumber++
			break GetChannelData
		}
	}



}

func tagCourses(channel chan models2.Recommend, isEnd chan bool, tid, id string) {

	rows, err := connect.Db.Queryx("select c.id, c.type, c.title, c.subtitle, c.difficulty_level, c.cover_picture from h_edu_courses as c left join h_taggables t on c.id = t.taggable_id where t.tag_id = " + tid + " and c.type in ('free', 'boutique') and c.status = 'published' and c.is_recommended = 1 and c.id != " + id +" group by c.id")

	if err != nil {
		log.Printf("查询错误:%s", err.Error())
		return
	}

	if rows == nil {
		log.Print("数据结果为空")
		return
	}

	var recommend models2.Recommend
	for rows.Next() {
		err := rows.StructScan(&recommend)
		if err != nil {
			log.Printf("数据结构化错误:%s", err.Error())
			return
		}

		channel <- recommend
	}

	defer func() {
		isEnd <- true
	}()
}
