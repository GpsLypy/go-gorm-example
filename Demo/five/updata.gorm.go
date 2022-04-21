package main

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type UserInfo struct {
	Name string
	Age  uint8
}

//@ update 只更新你选择的字段
//@ updates 更新所有字段，此时有两种形式，map，结构体0值不参与更新
//@ save 无论如何都更新所有的内容，包括0值
func TestUpdate() {
	// var TestUser TestUser

	//dbres := GLOBAL_DB.Model(&TestUser{}).Where("name LIKE ?", "l%").Update("name", "peng")
	//Save根据主键更新，无主键，自动插入新行
	//dbres := GLOBAL_DB.Model(&TestUser{}).Where("name LIKE ?", "l%").Save(&TestUser{Name:"qm1"})
	//那么要想做批量更新
	// var users []TestUser
	// dbres := GLOBAL_DB.Where("name = ?", "li").Find(&users)
	// for k := range users {
	// 	users[k].Age = 18
	// }
	// dbres.Save(&users)

	//map
	var user TestUser
	dbres := GLOBAL_DB.First(&user).Updates(map[string]interface{}{"Name": "", "Age": 0})

	//struct
	// var user TestUser
	// dbres := GLOBAL_DB.First(&user).Updates(TestUser{Name:"",Age:0})

	fmt.Println(errors.Is(dbres.Error, gorm.ErrRecordNotFound))

	//fmt.Println(u)

}
