package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// 数据库字段映射（标签可设置字段名、类型、默认值、索引、约束、备注等描述信息）
type UserInfo struct {
	UserID   uint       `gorm:"primaryKey;comment:'主键id'"`
	UserName string     `gorm:"column:user_name;type:varchar(50);index:idx_user_name;unique;not null;comment:'用户名'"`
	Email    string     `gorm:"column:email;type:varchar(64);index:idx_email;comment:'电子邮箱'"`
	Age      uint8      `gorm:"column:age;type:tinyint(8) unsigned;default:0;comment:'年龄'"`
	Birth    *time.Time `gorm:"column:birth_day;type:datetime;comment:'出生日期'"`
}

// 数据库表名映射
func (u UserInfo) TableName() string {
	return "t_user_info"
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
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	_ = db.AutoMigrate(&UserInfo{})

	// Create
	now := time.Now()
	db.Create(&UserInfo{
		UserID:   107,
		UserName: "鸣人",
		Email:    "naruto@mail.com",
		Birth:    &now,
	})
}
