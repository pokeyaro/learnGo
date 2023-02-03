# 数据库连接与配置

## 连接数据库

```go
import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var (
    ip     = "127.0.0.1"
    port   = 3306
    user   = "root"
    passwd = "123456"
    dbname = "test"
)

func main() {
    // refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, passwd, ip, port, dbname)
	
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    
    if err != nil {
        panic("failed to connect database")
    }
}
```

## 设置全局日志

允许在执行 orm 操作时，打印原生 sql

```go
import (
    "database/sql"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "log"
    "os"
    "time"
)

func main() {
    dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

    // Set the global logger
    newLogger := logger.New(
        log.New(os.Stdout, "\r\n", log.LstdFlags),  // io writer
        logger.Config{
            SlowThreshold:             time.Second, // Slow SQL threshold
            LogLevel:                  logger.Info, // Log level
            IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
            Colorful:                  true,        // Enable color
        },
    )

    // Globally mode
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
        Logger: newLogger,
    })
}
```
