package main

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

	// Set the global logger
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)

	// Globally mode
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:          newLogger, // 设置logger
		CreateBatchSize: 1000,      // 设置单条 sql 最多支持批量创建的条目数
	})
	if err != nil {
		panic("failed to connect database")
	}

	var user User

	// 使用save更新
	db.First(&user)
	user.Name = "漩涡鸣人"
	user.Age = 22
	db.Save(&user) // Save 是一个集 create 和 update 于一体的方法，如果检索到没有该条记录，则会进行创建！

	// 使用update更新
	db.Model(&User{}).Where("age = ?", 22).Update("name", "鸣人")
}
