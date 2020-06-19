package main

import (
	"edu_api_v2/route"
	_ "github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"log"
)

func main() {

	server := &fasthttp.Server{
		Handler:          route.Handler(),
		Name:             "API_SERVER",
		DisableKeepalive: false,
		LogAllErrors:     true,
	}

	if err := server.ListenAndServe("localhost:80"); err != nil {
		log.Fatalln(err)
	}

	//server := fasthttp.Server{Handler:"",Name:"API-SERVER", DisableKeepalive:false,LogAllErrors:true}
	//
	//if err := server.ListenAndServe("localhost"); err != nil {
	//	log.Fatalln(err)
	//}

	//db := connect.Db
	//if db != nil {
	//	rows, err := db.Queryx("select id, avatar from h_users where id = 4")
	//
	//	if err != nil {
	//		fmt.Printf("查询错误:")
	//	} else {
	//		user := models.User{}
	//		for rows.Next() {
	//			err := rows.StructScan(&user)
	//			if err != nil {
	//				log.Fatalln(err)
	//			}
	//			fmt.Printf("%#v\n", user)
	//		}
	//	}
	//}else {
	//	fmt.Printf("%s", "出错啦")
	//}
}
