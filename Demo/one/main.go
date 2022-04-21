package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type User struct {
	Name string
}

type UserTwo struct {
	Name string
}

func main() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:Mysql123..@tcp(47.99.176.238:3306)/blog?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		// //DefaultStringSize: 256, // string 类型字段的默认长度
		// DisableDatetimePrecision: true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		// DontSupportRenameIndex: true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		// DontSupportRenameColumn: true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		// SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		SkipDefaultTransaction: false, //是否跳过默认事务
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",  //表名前缀`User`的表名应该是`t_users`
			SingularTable: false, //使用单数表名 `User`->`t_user`
		},
		DisableForeignKeyConstraintWhenMigrating: true, //不会建立物理外键，代码里自动体现外键关系
	})
	fmt.Println(db, err)

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxIdleTime(time.Hour)

	//1、自动迁移演示
	//_=db.AutoMigrate(&User{})
	//2、手动创建表
	M := db.Migrator()
	M.CreateTable(&User{})
	//3、检查表是否存在
	fmt.Println(M.HasTable(&User{}))
	//4、删除表
	M.DropTable(&User{})
	fmt.Println(M.HasTable(&User{}))

	M.CreateTable(&User{})
	M.RenameTable(&User{}, &UserTwo{})

}
