package edu

import (
	"github.com/valyala/fasthttp"
	"edu_api_v2/connect"
	"log"
	"edu_api_v2/models"
	"edu_api_v2/response"
	"fmt"
	"strconv"
)

/**
获取课程详情
 */
func GetCourseDetail(ctx *fasthttp.RequestCtx) {
	var id = fmt.Sprintf("%s", ctx.QueryArgs().Peek("id"))

	rows, err := connect.Db.Queryx("select id, uuid, type, title, subtitle, price, vip_price, discount, discount_end_at, ifnull(learn_num, 0) + ifnull(buy_num, 0) as learn_count, length, rating, practical_rating, popular_rating, logic_rating, goals, audiences, summary from h_edu_courses where id = " + id)

	if err != nil {
		log.Printf("查询错误:%s", err.Error())
		return
	}

	courseDetail := models.CourseDetail{}

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
	}

	response.New(ctx).SetData(courseDetail).JsonReturn()
}

/**
获取课程用户
 */
func GetCourseUser(ctx *fasthttp.RequestCtx) {
	var (
		id = fmt.Sprintf("%s", ctx.QueryArgs().Peek("id"))
	)
	limit, _ := strconv.Atoi(fmt.Sprintf("%s", ctx.QueryArgs().Peek("limit")))

	var user_ids []string
	//随机获取一个讲师信息
	rows, err := connect.Db.Queryx("select user_id from h_edu_course_user where course_id = " + id)
	if err != nil {
		log.Printf("查询错误:%s", err.Error())
		return
	}

	if rows == nil {
		log.Print("数据结果为空")
		return
	}
	var user_id string
	for rows.Next() {
		err := rows.Scan(&user_id)
		if err != nil {
			log.Printf("数据结构化错误:%s", err.Error())
			return
		}

		user_ids = append(user_ids, user_id)
	}

	if len(user_ids) == 0 {
		log.Print("未找到相关讲师信息")
		return
	}

	courseUser := make(chan *models.User, len(user_ids))

	for _, uId := range user_ids {
		go user_info(uId, courseUser)
	}

	var users []*models.User
	for i := 0; i < len(user_ids); i++ {
		v := <-courseUser
		users = append(users, v)
	}

	if len(users) > 0 && limit == 1 {
		users = users[0:1]
	} else {
		if cap(users) <= limit {
			users = users[0:cap(users)]
		} else {
			users = users[0:limit]
		}
	}

	response.New(ctx).SetData(users).JsonReturn()
}

/**
获取用户信息
 */
func user_info(uId string, courseUser chan *models.User) {
	row := connect.Db.QueryRow("select id, avatar, nickname, title, about from h_users where id = " + uId)
	user := &models.User{}
	var (
		id       int64
		avatar   string
		nickname string
		title    int64
		about    string
	)

	err := row.Scan(&id, &avatar, &nickname, &title, &about)
	if err != nil {
		log.Printf("用户数据结构化错误:%s", err.Error())
		return
	}

	user.Id = &id
	user.Avatar = &avatar
	user.Nickname = &nickname
	user.Title = &title
	user.About = &about

	defer func() {
		courseUser <- user
	}()

	return
}
