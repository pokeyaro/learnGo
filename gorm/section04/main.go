package main

import (
	"database/sql"
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

	// Migrate the schema
	_ = db.AutoMigrate(&User{})

	// Create with struct
	var users = []User{{Name: "三代目"}, {Name: "四代目"}, {Name: "五代目"}}
	db.CreateInBatches(users, 100)

	for _, user := range users {
		fmt.Println(user.ID)
	}

	// Create with map
	db.Model(&User{}).Create([]map[string]interface{}{
		{"Name": "佐助", "Age": 18},
		{"Name": "小樱", "Age": 17},
	})
}
