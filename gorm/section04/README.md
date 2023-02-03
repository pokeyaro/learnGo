# CRUD —— Create

## 插入单条数据

```go
func main() {
    user := User{
        Name: "鸣人",
        Age:  18,
    }

    // 尽量通过数据的指针来创建
    result := db.Create(&user)

    fmt.Println(user.ID)             // 返回插入数据的主键
    fmt.Println(result.Error)        // 返回 error
    fmt.Println(result.RowsAffected) // 返回插入记录的条数
}
```

## 批量插入数据

- GORM 将生成单独一条 SQL 语句来插入所有数据

```go
func main() {
    var users = []User{{Name: "三代目"}, {Name: "四代目"}, {Name: "五代目"}}
    db.Create(&users)

    for _, user := range users {
        fmt.Println(user.ID)
    }
}
```

- 使用 CreateInBatches 分批创建时，可以指定每批的数量

```go
func main() {
    // 为什么不一次性提交所有，还要分批次呢？sql 语句是有长度限制的！当批量插入上万条记录的时候，是非常实用的！
    var users = []User{{name: "忍者_1"}, ...., {Name: "忍者_10000"}}

    // 数量为 100
    db.CreateInBatches(users, 100)
}
```

- 在初始化时，可以定义 CreateBatchSize 的个数

```go
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
    CreateBatchSize: 1000,
})
```

## 通过 Map 创建

- GORM 支持根据 map[string]interface{} 和 []map[string]interface{}{} 创建记录

```go
func main() {
    db.Model(&User{}).Create(map[string]interface{}{
        "Name": "鸣人", "Age": 18,
    })

    // batch insert from `[]map[string]interface{}{}`
    db.Model(&User{}).Create([]map[string]interface{}{
        {"Name": "佐助", "Age": 18},
        {"Name": "小樱", "Age": 17},
    })
}
```
