package connect

import (
	"edu_api_v2/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init() {
	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", config.Config.Mysql.DbUsername, config.Config.Mysql.DbPassword, config.Config.Mysql.DbHost, config.Config.Mysql.DbPort, config.Config.Mysql.DbDatabase)
	db, err := sqlx.Connect("mysql", dataSourceName)

	if err == nil && db != nil {
		Db = db
	} else {
		fmt.Printf("mysql 连接错误")
	}
}

/**
获取表名
*/
func Table(prefix, table string) string {
	return prefix + table
}
