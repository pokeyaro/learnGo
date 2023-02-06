package main

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

type User struct {
	ID    uint
	UUID  uuid.UUID
	Name  string
	Age   uint8
	Role  string
	Ctime time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// 添加ctime
	u.Ctime = time.Now()

	// 添加uuid
	u.UUID = uuid.New()

	// 判断role
	if u.Role == "admin" {
		return errors.New("invalid role")
	}
	return
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
		// 设置统一的表名前缀
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "t_",
		},
		Logger:          newLogger, // 设置logger
		CreateBatchSize: 1000,      // 设置单条 sql 最多支持批量创建的条目数
	})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移
	_ = db.AutoMigrate(&User{})

	db.Create(&User{
		Name: "鸣人",
		Age:  20,
		Role: "admin",
	})
}
