package main

import (
	_ "database/sql"
	"fmt"
	"time"

	"gorm.io/gorm"
)

//自定义Model
type Model struct {
	UUID uint      `gorm:"primaryKey"`
	Time time.Time `gorm:"column:timer"`
}

type TestUser struct {
	gorm.Model
	Name string `gorm:"default:qm"`
	Age  uint8  `gorm:"comment:年龄"`
}

func TestUserCreate() {
	GLOBAL_DB.AutoMigrate(&TestUser{})
}

func CreatedTest() {
	//dbres := GLOBAL_DB.Create(&TestUser{Name: "li", Age: 20})
	//只创建Name
	//dbres :=GLOBAL_DB.Select("Name").Create(&TestUser{Name:"liyupeng",Age:18})
	//创建除了Name之外的
	//dbres :=GLOBAL_DB.Omit("Name").Create(&TestUser{Name:"liyupeng",Age:18})
	dbres := GLOBAL_DB.Create(&[]TestUser{
		{Name: "li", Age: 20},
	})
	fmt.Println(dbres.Error, dbres.RowsAffected)
	if dbres.Error != nil {
		fmt.Println("erro create")
	} else {
		fmt.Println("success!")
	}
}
