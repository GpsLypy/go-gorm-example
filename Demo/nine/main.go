package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "gorm.io/gorm/schema"
)

var GLOBAL_DB *gorm.DB

func main() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:Mysql123..@tcp(47.99.176.238:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local", // DSN data source name
	}), &gorm.Config{
		SkipDefaultTransaction: false, //是否跳过默认事务
		// NamingStrategy: schema.NamingStrategy{
		// 	TablePrefix:   "t_",  //表名前缀`User`的表名应该是`t_users`
		// 	SingularTable: false, //使用单数表名 `User`->`t_user`
		// },
		DisableForeignKeyConstraintWhenMigrating: true, //不会建立物理外键，代码里自动体现外键关系
	})
	fmt.Println(111)
	fmt.Println(db, err)

	sqlDB, err := db.DB()
	fmt.Println(err)
	GLOBAL_DB = db
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxIdleTime(time.Hour)
	fmt.Println(222)
	//TestUserCreate()
	//CreatedTest()
	//TestUserCreate()
	CreatedTest()
}
