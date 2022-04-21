package main

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	_ "gorm.io/gorm"
)

type Info struct {
}

type CInfo struct {
	Name string
	Age  int
}

func (c CInfo) Value() (driver.Value, error) {
	str, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	return string(str), nil
}

func (c *CInfo) Scan(value interface{}) error {
	str, ok := value.([]byte)
	if !ok {
		return errors.New("不匹配的数据类型")
	}
	json.Unmarshal(str, c)
	return nil
}

type CUser struct {
	ID   uint
	Info CInfo
}

func Customize() {
	GLOBAL_DB.AutoMigrate(&CUser{})
	//GLOBAL_DB.Create(&CUser{Info: CInfo{Name: "haha", Age: 20}})
	var user CUser
	GLOBAL_DB.First(&user)
}
