package models

import (
	"Pjue-Pjue/assets_daily_clear_web/config"
	"database/sql"
	"fmt"
	"github.com/go-gorp/gorp"
	// mysql
	_ "github.com/go-sql-driver/mysql"
)

var (
	// DB 数据操作对象
	DB *gorp.DbMap
)

func init() {
	conf := config.GetConfiguration()
	// user:password@tcp(127.0.0.1:3306)/hello
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		conf.DbUser,
		conf.DbPassword,
		conf.DbHost,
		conf.DbPort,
		conf.DbName,
	)

	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	// construct a gorp DbMap
	DB = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}

	// 注册实体
	DB.AddTableWithName(Account{}, "accounts").SetKeys(false, "id")
	DB.AddTableWithName(DailyAsset{}, "daily_asset_usd_valuate").SetKeys(false, "id")

}
