# CRUD —— Read

## 查询单个对象

GORM 提供了 First、Take、Last 方法，以便从数据库中检索单条记录

```go
func main() {
    var user User

    // 获取第一条记录（主键升序）
    db.First(&user)
    // SELECT * FROM users ORDER BY id LIMIT 1;

    // 获取一条记录，没有指定排序字段
    db.Take(&user)
    // SELECT * FROM users LIMIT 1;

    // 获取最后一条记录（主键降序）
    db.Last(&user)
    // SELECT * FROM users ORDER BY id DESC LIMIT 1;
}
```

当查询数据且没有找到记录时，它会返回 ErrRecordNotFound 错误

```go
func main() {
    var user User
    
    result := db.First(&user)
    
    fmt.Println(result.RowsAffected) // 返回找到的记录数
    fmt.Println(result.Error)        // returns error or nil

    // 检查 ErrRecordNotFound 错误
    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
        fmt.Println("未找到数据")
    }
}
```

## 根据主键检索

```go
func main() {
    var user User
    
    result := db.First(&user, 10)
    // SELECT * FROM `users` WHERE `users`.`id` = 10 ORDER BY `users`.`id` LIMIT 1

    result := db.First(&user, []int{1, 2, 3})
    // SELECT * FROM `users` WHERE `users`.`id` IN (1,2,3) ORDER BY `users`.`id` LIMIT 1
}
```

## 检索全部对象

```go
func main() {
    var users []User

    result := db.Find(&users)
    fmt.Println("总共记录：", result.RowsAffected)  // 相当于 len(users)

    for _, user := range users {
        fmt.Println(user.Name)
    }
}
```
