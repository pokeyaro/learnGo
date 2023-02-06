# CRUD —— Delete

## 删除一条记录

删除一条记录时，删除对象需要指定主键，否则会触发 批量 Delete，例如：

```go
func main() {
    var user User
	
    // user 的 ID 是 `10`
    db.Delete(&user)
    // DELETE from users where id = 10;

    // 带额外条件的删除
    db.Where("name = ?", "鸣人").Delete(&user)
    // DELETE from users where id = 10 AND name = "鸣人";
}
```

## 根据主键删除

```go
func main() {
    var user User

    db.Delete(&User{}, 10)
    // DELETE FROM users WHERE id = 10;
	
    db.Delete(&users, []int{1,2,3})
    // DELETE FROM users WHERE id IN (1,2,3);
}
```

## 批量删除

```go
func main() {
    var user User

    db.Where("name LIKE ?", "%代目%").Delete(&User{})
    // DELETE from users where name LIKE "%代目%";

    db.Delete(&User{}, "name LIKE ?", "%代目%")
    // DELETE from users where name LIKE "%代目%";
}
```

## 软删除

- 如果您的模型包含了一个 gorm.deletedat 字段（gorm.Model 已经包含了该字段)，它将自动获得软删除的能力！
- 拥有软删除能力的模型调用 Delete 时，记录不会从数据库中被真正删除。但 GORM 会将 DeletedAt 置为当前时间， 并且你不能再通过普通的查询方法找到该记录。

```go
func main() {
    var user User

    // user 的 ID 是 `111`
    db.Delete(&user)
    // UPDATE users SET deleted_at="2013-10-29 10:23" WHERE id = 111;

    // 批量删除
    db.Where("age = ?", 20).Delete(&User{})
    // UPDATE users SET deleted_at="2013-10-29 10:23" WHERE age = 20;

    // 在查询时会忽略被软删除的记录
    db.Where("age = 20").Find(&user)
    // SELECT * FROM users WHERE age = 20 AND deleted_at IS NULL;
}
```

- 如果您不想引入 gorm.Model，您也可以这样启用软删除特性：

```go
type User struct {
    ID      int
    Deleted gorm.DeletedAt
    Name    string
}
```
