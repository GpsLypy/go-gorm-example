package main

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func TestFindOne() {
	//实例化
	//result:= map[string]interface{}{}
	//GLOBAL_DB.Model(&TestUser{}).First(&result)
	//主键检索
	var User TestUser
	//dbres := GLOBAL_DB.Model(&TestUser{}).First(&User, 14)
	//where条件
	//dbres := GLOBAL_DB.Where("name=?", "li").First(&User)
	dbres := GLOBAL_DB.First(&User, "name=?", "qm")

	fmt.Println(errors.Is(dbres.Error, gorm.ErrRecordNotFound))
	//GLOBAL_DB.Model(&TestUser{}).Take(&User)

	fmt.Println(User)

}

func TestFinds() {
	var User []TestUser
	GLOBAL_DB.Where("name<>?", "li").Find(&User)
	fmt.Println(User)
}
