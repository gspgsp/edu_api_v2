package main

import (
	"edu_api_v2/connect"
	"edu_api_v2/models"
	"fmt"
	_ "github.com/buaazp/fasthttprouter"

	_ "github.com/valyala/fasthttp"
	"log"
)

func main() {
	db := connect.Db
	if db != nil {
		rows, err := db.Queryx("select id, avatar from h_users where id = 4")

		if err != nil {
			fmt.Printf("查询错误:")
		} else {
			user := models.User{}
			for rows.Next() {
				err := rows.StructScan(&user)
				if err != nil {
					log.Fatalln(err)
				}
				fmt.Printf("%#v\n", user)
			}
		}
	} else {
		fmt.Printf("%s", "出错啦")
	}
}
