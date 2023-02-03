# 表的迁移与 Crud 操作

## 自动迁移

```go
// Define table structure
type Product struct {
    gorm.Model
    Code  string
    Price uint
}

func main() {
    ...

    // Migrate the schema
    _ = db.AutoMigrate(&Product{})
}
```

原生 SQL 执行

```sql
SELECT DATABASE();

SELECT SCHEMA_NAME from Information_schema.SCHEMATA where SCHEMA_NAME LIKE 'test%' ORDER BY SCHEMA_NAME='test' DESC,SCHEMA_NAME limit 1;

SELECT count(*) FROM information_schema.tables WHERE table_schema = 'test' AND table_name = 'products' AND table_type = 'BASE TABLE';

CREATE TABLE `products` (
    `id` bigint unsigned AUTO_INCREMENT,
    `created_at` datetime(3) NULL,
    `updated_at` datetime(3) NULL,
    `deleted_at` datetime(3) NULL,
    `code` longtext,
    `price` bigint unsigned,
    PRIMARY KEY (`id`),
    INDEX `idx_products_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
```

## Crud 操作

- update：Updates 方法 struct 方式仅更新非零值的字段，若为零值，是不做更新操作的！map 方式可更新零值！
- delete：仅作逻辑删除，本质上 deleted_at 字段默认为空，若更新标记上时间，则算作被删除！

```go
func main() {
    ...

    // Create
    db.Create(&Product{Code: "D42", Price: 100})

    // Read
    var product Product
    db.First(&product, 1) // find product with integer primary key
    db.First(&product, "code = ?", "D42") // find product with code D42

    // Update - update product's price to 200
    db.Model(&product).Update("Price", 200)
    // Update - update multiple fields
    db.Model(&product).Updates(Product{Price: 400, Code: "C63"}) // non-zero fields
    db.Model(&product).Updates(map[string]interface{}{"Price": 0, "Code": ""})

    // Delete - delete product
    db.Delete(&product, 1)
}
```

原生 SQL 执行

```sql
-- 插入条目
INSERT INTO `products` (`created_at`,`updated_at`,`deleted_at`,`code`,`price`) VALUES ('2023-02-03 15:01:35.835','2023-02-03 15:01:35.835',NULL,'D42',100);

-- 查询第一条
SELECT * FROM `products` WHERE `products`.`id` = 1 AND `products`.`deleted_at` IS NULL ORDER BY `products`.`id` LIMIT 1;

-- 带where的查询
SELECT * 
FROM `products` 
WHERE code = 'D42' 
  AND `products`.`deleted_at` IS NULL 
  AND `products`.`id` = 1 
ORDER BY `products`.`id` 
LIMIT 1;

-- 更新单个字段
UPDATE `products` SET `price`=200,`updated_at`='2023-02-03 15:01:35.89' WHERE `products`.`deleted_at` IS NULL AND `id` = 1

-- 更新多个字段（该orm方式不能为零值）
UPDATE `products` SET `updated_at`='2023-02-03 15:01:35.92',`code`='C63',`price`=400 WHERE `products`.`deleted_at` IS NULL AND `id` = 1

-- 更新多个字段（该orm方式可为零值）
UPDATE `products` SET `code`='',`price`=0,`updated_at`='2023-02-03 15:01:35.946' WHERE `products`.`deleted_at` IS NULL AND `id` = 1

-- 逻辑删除
UPDATE `products` SET `deleted_at`='2023-02-03 15:01:35.981' WHERE `products`.`id` = 1 AND `products`.`id` = 1 AND `products`.`deleted_at` IS NULL;
```

## 解决 Updates 方法更新零值的问题

### 将 type 设置为指针 *type

```go
type Product struct {
    gorm.Model
    Code  *string
    Price *uint
}

func main() {
    ...

    // Create
    valInt := 100
    valStr := "D42"
    db.Create(&Product{
        Code: &valStr,
        Price: &valInt,
    })

    // Update - update product's price to 0, code to ""
    zeroInt := 0
    zeroStr := ""
    db.Model(&Product{ID:1}).Updates(Product{
        Price: &zeroInt, 
        Code: &zeroStr,
    })
}
```

### 使用 sql.NullType 来解决

```go
// 模型的字段类型定义为nulltype
type Product struct {
    gorm.Model
    Code  sql.NullString
    Price sql.NullInt64
}

func main() {
    ...

    // Create
    db.Create(&Product{
        Code: sql.NullString{String: "D42", Valid: true},
        Price: sql.NullInt64{Int64: 100, Valid: true},
    })

    // Update - update product's price to 0, code to ""
    var product Product
    db.Model(&product).Updates(Product{Price: 0, Code: ""})
}
```

原生 SQL 执行

```sql
-- 使用NullType可更新非零值
INSERT INTO `products` (`created_at`,`updated_at`,`deleted_at`,`code`,`price`) VALUES ('2023-02-03 14:52:17.93','2023-02-03 14:52:17.93',NULL,'D42',100);

-- 使用NullType可更新零值
UPDATE `products` SET `updated_at`='2023-02-03 14:52:17.954',`code`='',`price`=0 WHERE `products`.`deleted_at` IS NULL;
```
