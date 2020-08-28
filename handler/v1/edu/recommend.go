package edu

import (
	"github.com/valyala/fasthttp"
	"fmt"
	"log"
	"edu_api_v2/connect"
	"edu_api_v2/models"
	"edu_api_v2/response"
	"strconv"
	"errors"
)

var limit = "5"

type Tag struct {
	Id string `db:"id"json:"id"`
}

/**
推荐课程
 */
func GetRecommend(ctx *fasthttp.RequestCtx) {
	var (
		id = fmt.Sprintf("%s", ctx.QueryArgs().Peek("id"))
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
		recommends []models.Boutique
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
		if recommendData, err := defaultData(id, limit); err == nil {
			recommends = recommendData
		}
	} else {
		if recommendData, err := recommendData(tags, id); err == nil {
			recommends = recommendData
		}
	}

	response.New(ctx).SetData(recommends).JsonReturn()
}

/**
默认推荐数据
 */
func defaultData(id, limit string) ([]models.Boutique, error) {
	var (
		recommend  models.Boutique
		recommends []models.Boutique
	)

	rows, err := connect.Db.Queryx("select id, type, title, subtitle, difficulty_level, cover_picture, price, discount, discount_end_at, vip_level, vip_price, ifnull(learn_num, 0) + ifnull(buy_num, 0) as learn_count from h_edu_courses where type in ('free', 'boutique') and status = 'published' and is_recommended = 1 and id != " + id + " limit " + limit)

	if err != nil {
		log.Printf("查询错误:%s", err.Error())
		return nil, err
	}

	if rows == nil {
		log.Print("数据结果为空")
		return nil, errors.New("数据结果为空")
	}

	for rows.Next() {
		err := rows.StructScan(&recommend)
		if err != nil {
			log.Printf("数据结构化错误:%s", err.Error())
			return nil, err
		}

		recommends = append(recommends, recommend)
	}

	return recommends, nil
}

/**
获取推荐数据
 */
func recommendData(tags []Tag, id string) (recommends []models.Boutique, err error) {

	channel := make(chan models.Boutique, len(tags))
	isEnd := make(chan bool)
	times := 0

	for _, val := range tags {
		go tagCourses(channel, isEnd, val.Id, id)
	}

GetChannelData:
	for {
		select {
		case v, ok := <-channel:
			if ok {
				recommends = append(recommends, v)
			} else {
				log.Printf("当前channel")
			}
		case <-isEnd:
			times++
			if times == len(tags) {
				//超量
				val, _ := strconv.Atoi(limit)
				if len(recommends) >= val {
					recommends = recommends[0:val]
					break GetChannelData
				} else if len(recommends) > 0 {
					//补位
					if recommendData, err := defaultData(id, strconv.Itoa(val-len(recommends))); err == nil {
						recommends = append(recommends, recommendData...)
						break GetChannelData
					}
				} else {
					//默认
					if recommendData, err := defaultData(id, limit); err == nil {
						recommends = recommendData
						break GetChannelData
					}
				}
			}
		}
	}

	close(channel)
	close(isEnd)

	return recommends, nil
}

/**
推荐课程信息
 */
func tagCourses(channel chan models.Boutique, isEnd chan bool, tid, id string) {
	rows, err := connect.Db.Queryx("select c.id, c.type, c.title, c.subtitle, c.difficulty_level, c.cover_picture, c.price, c.discount, c.discount_end_at, c.vip_level, c.vip_price, ifnull(c.learn_num, 0) + ifnull(c.buy_num, 0) as learn_count from h_edu_courses as c left join h_taggables t on c.id = t.taggable_id where t.tag_id = " + tid + " and c.type in ('free', 'boutique') and c.status = 'published' and c.is_recommended = 0 and c.id != " + id + " group by c.id")

	if err != nil {
		log.Printf("查询错误:%s", err.Error())
		return
	}

	if rows == nil {
		log.Print("数据结果为空")
		return
	}

	var recommend models.Boutique
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
