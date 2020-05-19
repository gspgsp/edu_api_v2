package main

import (
	"edu_api_v2/config"
	"edu_api_v2/models"
	"fmt"
	_ "github.com/buaazp/fasthttprouter"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/valyala/fasthttp"
	"log"
)

func main() {

	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", config.Config.Mysql.DbUsername, config.Config.Mysql.DbPassword, config.Config.Mysql.DbHost, config.Config.Mysql.DbPort, config.Config.Mysql.DbDatabase)
	db, err := sqlx.Connect("mysql", dataSourceName)

	if err != nil {
		fmt.Printf("mysql 连接错误")
	}

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
	}
}
