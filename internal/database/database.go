package database

import (
	"database/sql"
	entsql "entgo.io/ent/dialect/sql"
	"orse/ent"
	"orse/internal/config"
	"time"
)

func Open() (*ent.Client, error) {
	db, err := sql.Open(dsn())
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	// 从db变量中构造一个ent.Driver对象。
	drv := entsql.OpenDB("mysql", db)
	return ent.NewClient(ent.Driver(drv)), nil
}

func dsn()(driver string,dsn string)  {

	cfg := config.GetConfig("config/user/user.yaml")

	return  "mysql",cfg.Mysql.DataSource
}
