package main

import "gorm.io/gorm"

type TMG struct {
	ID   uint
	Name string
}

func TestTransaction() {
	GLOBAL_DB.AutoMigrate(&TMG{})
	GLOBAL_DB.Transaction(func(tx *gorm.DB) error {
		tx.Create(&TMG{Name: "汉子"})
		tx.Create(&TMG{Name: "威武"})
		tx.Create(&TMG{Name: "牛逼"})
		return nil
	})
}

func TestTransaction2() {
	GLOBAL_DB.AutoMigrate(&TMG{})
	tx := GLOBAL_DB.Begin()

	tx.Create(&TMG{Name: "汉子2"})
	tx.Create(&TMG{Name: "威武2"})
	tx.SavePoint("duo")
	tx.Create(&TMG{Name: "牛逼2"})
	tx.RollbackTo("duo")
	tx.Commit()
}
