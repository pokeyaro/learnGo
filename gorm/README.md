# Gorm 快速入门

## 什么是 ORM

- ORM 全称是：Object Relational Mapping（对象关系映射），其主要作用是在编程中，把面向对象的概念跟数据库中的表的概念对应起来。
- 举例来说，定义一个对象，对应着一张表；这个对象的实例，对应着表里的一条数据。

## 常用 ORM

https://github.com/go-gorm/gorm ✅ [官方文档](https://gorm.io/zh_CN/docs/)

https://github.com/jmoiron/sqlx

https://github.com/ent/ent

## 写在前面

> 学习要以 sql 为主，orm 为辅！多花精力在 sql 上，远比 orm 要重要的多！

优点：
- 增加代码的可维护性与提高开发效率
- 屏蔽 sql 细节，可以自动对实体 Entity 对象与数据库中的 Table 进行字段与属性的映射，不用写原生 sql
- 屏蔽各种数据库之间的语法差异

缺点：
- orm 会牺牲程序的执行效率
- 太过依赖 orm 会导致 sql 理解不够
