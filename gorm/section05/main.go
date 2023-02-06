package main

import (
	"database/sql"
	"errors"
	"fmt"
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

	// 获取第一条记录（主键升序）
	//db.First(&user)

	// 获取一条记录，没有指定排序字段
	//db.Take(&user)

	// 获取最后一条记录（主键降序）
	//db.Last(&user)

	// 检索主键
	res1 := db.First(&user, []int{1, 2, 3})
	fmt.Println(res1.Error) // returns error or nil
	// 检查 ErrRecordNotFound 错误
	retBool := errors.Is(res1.Error, gorm.ErrRecordNotFound)
	fmt.Println(retBool)

	// 检索全部对象
	var users []User
	res2 := db.Find(&users)
	fmt.Println("总共记录：", res2.RowsAffected)
	for _, user := range users {
		fmt.Println(user.Name)
	}
}
