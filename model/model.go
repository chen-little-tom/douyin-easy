// Package model /*
package model

import (
	"douyin-easy/config"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Model struct {
	CreateAt time.Time `json:"createAt" gorm:"column:create_at"` // 创建时间
	UpdateAt time.Time `json:"updateAt" gorm:"column:update_at"` // 更新时间
}

var DB *gorm.DB

func InitDB() {
	app := config.GlobalConfig
	host, port, dbname := app.Database.Host, app.Database.Port, app.Database.Dbname
	username, password := app.Database.Username, app.Database.Password

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tb_", // 配置数据表前缀
			SingularTable: true,  // 配置数据表单数
		},
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Printf("connect to database is failure, err: %s\n", err)
		log.Fatal("Application stop")
	}
	DB = db
}
