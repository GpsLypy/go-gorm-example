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

//智能选择字段
func TestFindOne() {
	var u UserInfo

	dbres := GLOBAL_DB.Model(UserInfo{}).Where("name LIKE ?", "").Find(&u)

	fmt.Println(errors.Is(dbres.Error, gorm.ErrRecordNotFound))

	fmt.Println(u)

}
