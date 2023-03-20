package requests

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

func InitDB(DBFilePath string, models []interface{}) {

	db, err := gorm.Open(sqlite.Open(DBFilePath), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败")
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	err = db.AutoMigrate(
		models...,
	)
	if err != nil {
		return
	}
	fmt.Println("register table success")
	DB = db

}