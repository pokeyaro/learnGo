# CRUD —— Update

## 保存所有字段

Save 会保存所有的字段，即使字段是零值

```go
func main() {
    var user User
	
    db.First(&user)
    user.Name = "漩涡鸣人"
    user.Age = 22

    // Save 是一个集 create 和 update 于一体的方法，如果检索到没有该条记录，则会进行创建！
    db.Save(&user)
}
```

## 更新单个列

```go
func main() {
    var user User

    // 条件更新
    db.Model(&User{}).Where("active = ?", true).Update("name", "鸣人")
    // UPDATE users SET name='鸣人', updated_at='2013-11-17 21:34:10' WHERE active=true;

    // 如果传入 user 实例，可以获取 User 的 ID 是 `111`
    db.Model(&user).Update("name", "鸣人")
    // UPDATE users SET name='鸣人', updated_at='2013-11-17 21:34:10' WHERE id=111;

    // 根据条件和 model 的值进行更新
    db.Model(&user).Where("active = ?", true).Update("name", "鸣人")
    // UPDATE users SET name='鸣人', updated_at='2013-11-17 21:34:10' WHERE id=111 AND active=true;
}
```

## 更新多个列

Updates 方法支持 struct 和 map[string]interface{} 参数。当使用 struct 更新时，默认情况下，GORM 只会更新非零值的字段

```go
func main() {
    var user User

    // 根据 `struct` 更新属性，只会更新非零值的字段
	db.Model(&user).Updates(User{Name: "鸣人", Age: 18, Active: false})
    // UPDATE users SET name='鸣人', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = 111;

    // 根据 `map` 更新属性
    db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
    // UPDATE users SET name='鸣人', age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=111;
}
```
