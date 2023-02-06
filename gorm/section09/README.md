# 模型的常用钩子方法

## 自定义 DB 表名

> 注意：NamingStrategy 与 TableName 方法不能同时配置，TableName 优先级比较高！

- 默认表名是 Struct 名称的小写复数的形式，如 `User` 结构体的表名为 `users`

```go
type User struct {
    gorm.Model
    Name string
    Age  uint8
}
```

- 在 GORM 中可以通过给 Struct 添加 TableName 方法来自定义表名

```go
func (User) TableName() string {
    return "table_user_info"
}
```

- 统一给所有的表名前加上一个前缀，需要在 &gorm.Config 中配置

```go
import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/schema"
)

func main() {
    // 全局模式 
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
        // 设置统一的表名前缀
        NamingStrategy: schema.NamingStrategy{
            TablePrefix: "t_",
        },
    })
}
```

## 插入 hook 钩子

在添加一个新的条目时，自动加上 ctime 当前时间，uuid 值，判断当前角色等钩子行为

```go
type User struct {
    ID    uint
    UUID  string
    Name  string
    Age   uint8
    Role  string
    Ctime time.Time
}
```

```go
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
    // 添加ctime
    u.Ctime = time.Now()
	
    // 添加uuid
    u.UUID = uuid.New().String()
	
    // 判断role
    if u.Role == "admin" {
        return errors.New("invalid role")
    }
    return
}
```
