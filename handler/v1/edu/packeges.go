package edu

import (
	"github.com/valyala/fasthttp"
	"edu_api_v2/connect"
	"log"
	"edu_api_v2/models"
	"edu_api_v2/response"
)

func PackageList(ctx *fasthttp.RequestCtx) {

	rows, err := connect.Db.Queryx(`select att.id, att.title, att.subtitle, att.cover_picture, ifnull(learn_num, 0) + ifnull(buy_num, 0) as learn_num, avl.length  from h_edu_packages as att LEFT JOIN(
	select aa.package_id,sum(bb.length) as length
		from h_edu_package_course as aa LEFT JOIN h_edu_courses bb on aa.course_id=bb.id
		GROUP BY aa.package_id) as avl on att.id=avl.package_id`)

	if err != nil {
		log.Printf("查询错误:%s", err.Error())
		return
	}

	pkg := models.Package{}
	var packages []models.Package
	if rows == nil {
		log.Print("数据结果为空")
		return
	}

	for rows.Next() {
		err := rows.StructScan(&pkg)
		if err != nil {
			log.Printf("数据结构化错误:%s", err.Error())
		} else {
			packages = append(packages, pkg)
		}
	}

	response.New(ctx).SetData(packages).JsonReturn()
}
