package main

import (
	_ "errors"
	_ "fmt"

	_ "gorm.io/gorm"
)

func TestDelete() {
	var user TestUser
	//软删除
	GLOBAL_DB.Where("name =?", "").Delete(&user)
	//硬删除
	GLOBAL_DB.Unscoped().Where("name = ?", "").Delete(&user)

	GLOBAL_DB.Raw("select * FROM test_users WHERE name = ?", "li").Scan(&user)
}
