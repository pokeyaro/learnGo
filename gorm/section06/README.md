# CRUD —— Read (II)

## 概述

- GORM 的查询方式有以下 3 种

| 方式        | 说明                                                               |
|-----------|------------------------------------------------------------------|
| String 方式 | 最灵活的方式，接近原生 SQL 的写法，注意需要拼凑查询条件字段为 DB 中的字段名。                      |
| Struct 方式 | 最便捷的方式，查询条件直接使用 Struct 的字段名，屏蔽很多细节，但注意若条件中存在零值的情况，这种方式就比较**坑**了！ |
| Map 方式    | 最易于理解的方式，注意同样需要拼凑查询条件字段为 DB 中的字段名，没有坑点，可接受条件中存在零值的情况。            |

## 使用 String 条件

- 优点：拼接 SQL 十分灵活，最接近原生 SQL 写法
- 缺点：key 必须是 DB 的字段名，而非 ORM 中 Struct 的字段名

```go
func main() {
    var user User

    // Get first matched record
    db.Where("name = ?", "鸣人").First(&user)
    // SELECT * FROM users WHERE name = '鸣人' ORDER BY id LIMIT 1;

    // Get all matched records
    db.Where("name <> ?", "鸣人").Find(&users)
    // SELECT * FROM users WHERE name <> '鸣人';

    // IN
    db.Where("name IN ?", []string{"三代目", "四代目"}).Find(&users)
    // SELECT * FROM users WHERE name IN ('三代目','四代目');

    // LIKE
    db.Where("name LIKE ?", "%代%").Find(&users)
    // SELECT * FROM users WHERE name LIKE '%代%';

    // AND
    db.Where("name = ? AND age >= ?", "鸣人", "18").Find(&users)
    // SELECT * FROM users WHERE name = '鸣人' AND age >= 18;

    // Time
    db.Where("updated_at > ?", lastWeek).Find(&users)
    // SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';

    // BETWEEN
    db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)
    // SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';
}
```

## 使用 Struct 条件

- 优点：Key 是 ORM 中 Struct 的字段名，可屏蔽掉 DB 数据库中的字段名称
- 缺点：如果条件中存在零值的情况，则会忽略该相关的条件，不会构造到 SQL 中

```go
func main() {
    var user User

    res := db.Where(&User{Name: "佐助"}).First(&user)
    fmt.Println(res.RowsAffected)
	
    // 如果存在零值的情况
    db.Where(&User{Name: "佐助", Age: 0}).Find(&users)
    // SELECT * FROM users WHERE name = "佐助";
}
```

## 使用 Map 条件

- 优点：如果条件中有零值的情况，会拼凑其 SQL
- 缺点：Key 必须是 DB 的字段名，而非 ORM 中 Struct 的字段名

```go
func main() {
    var user User

    // 如果存在零值的情况
    db.Where(map[string]interface{}{"name": "佐助", "age": 0}).Find(&users)
    // SELECT * FROM users WHERE name = "佐助" AND age = 0;
}
```

## Or 条件

```go
func main() {
    var user User
	
    // String
    db.Where("name = ?", "鸣人").Or("name = ?", "佐助").Find(&users)
    // SELECT * FROM users WHERE name = '鸣人' OR name = '佐助';

    // Struct
    db.Where("name = '鸣人'").Or(User{Name: "佐助", Age: 20}).Find(&users)
    // SELECT * FROM users WHERE name = '鸣人' OR (name = '佐助' AND age = 20);

    // Map
    db.Where("name = '鸣人'").Or(map[string]interface{}{"name": "佐助", "age": 20}).Find(&users)
    // SELECT * FROM users WHERE name = '鸣人' OR (name = '佐助' AND age = 20);
}
```
