# 表的 struct 模型定义

## 约定

- GORM 倾向于约定优于配置，如果您遵循 GORM 的约定，您就可以少写的配置、代码。
  - 默认情况下，GORM 使用 ID 作为主键；
  - 使用结构体名的 '蛇形复数' 作为表名；
  - 使用字段名的 '蛇形' 作为列名；
  - 使用 CreatedAt、UpdatedAt 字段追踪创建、更新时间；

## 基础示例

```go
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
```

## gorm.Model

一个基础的模型，可以将它嵌入到您的结构体中，以包含这几个字段

```go
// gorm.Model 的定义
type Model struct {
    ID        uint           `gorm:"primaryKey"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
}
```

## 自定义 tag 标签

GORM 在声明 model 时，tag 是可选的，tag 名大小写不敏感，建议使用 camelCase 小驼峰命名风格

```go
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
```

原生 SQL 执行

```sql
CREATE TABLE `t_user_info` (
    `user_id` bigint unsigned AUTO_INCREMENT COMMENT '\'主键id\'',
    `user_name` varchar(50) NOT NULL UNIQUE COMMENT '\'用户名\'',
    `email` varchar(64) COMMENT '\'电子邮箱\'',
    `age` tinyint(8) unsigned DEFAULT 0 COMMENT '\'年龄\'',
    `birth_day` datetime COMMENT '\'出生日期\'',
    PRIMARY KEY (`user_id`),
    INDEX `idx_user_name` (`user_name`),
    INDEX `idx_email` (`email`)
);

INSERT INTO `t_user_info` (`user_name`,`email`,`age`,`birth_day`,`user_id`) VALUES ('鸣人','naruto@mail.com',0,'2023-02-03 16:35:45.478',107);
```
