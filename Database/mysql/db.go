package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

const (
	dsn = "root:8520@tcp(127.0.0.1:3306)/studygroup?charset=utf8mb4&parseTime=True&loc=Local"
)

var (
	DBsql *gorm.DB
)

// 初始化数据库连接
func init() {
	sql, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,  // string 类型字段的默认长度
		SkipInitializeWithVersion: true, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db, err := sql.DB()
	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxOpenConns(10)
	DBsql = sql
}
