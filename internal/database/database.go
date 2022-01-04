package database

import (
	"database/sql"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	"gopkg.in/ini.v1"
	"orse/ent"
	"os"
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
	cfg, err := ini.Load("configs/config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	dataUser := cfg.Section("database").Key("USER").String()
	dataPassword := cfg.Section("database").Key("PASSWORD").String()
	dataHost := cfg.Section("database").Key("HOST").String()
	dataName := cfg.Section("database").Key("NAME").String()
	dataType := cfg.Section("database").Key("TYPE").String()

	return  dataType,dataUser + ":" + dataPassword + "@tcp(" + dataHost + ")/" + dataName + "?parseTime=True"
}
