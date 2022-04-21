package main

import (
	"fmt"

	"gorm.io/gorm"
)

type Info struct {
	gorm.Model
	Money int
	DogID uint
}

//has one
type Dog struct {
	gorm.Model
	Name      string
	GirlGodID uint //舔狗心里有女神了(属于)
	//GirlGod   GirlGod
	Info Info
}

type GirlGod struct {
	gorm.Model
	Name string
	Dogs []Dog //女神拥有舔狗
}

func One2one() {
	//GLOBAL_DB.AutoMigrate(&Dog{},&GirlGod{})
	// d := Dog{
	// 	Model: gorm.Model{
	// 		ID: 1,
	// 	},
	// 	Name: "WAWA",
	// }

	// d2 := Dog{
	// 	Model: gorm.Model{
	// 		ID: 2,
	// 	},
	// 	Name: "WAWA2",
	// }
	// g := GirlGod{
	// 	Model: gorm.Model{
	// 		ID: 1,
	// 	},
	// 	Name: "QIQI",
	// 	Dogs: []Dog{d, d2},
	// }

	//GLOBAL_DB.Create(&g)

	// var girl GirlGod
	// GLOBAL_DB.Preload("Dogs").First(&girl)
	// //GLOBAL_DB.Preload("Dogs","name = ?","QIQI").First(&girl)
	// GLOBAL_DB.Preload("Dogs", func(db *gorm.DB) *gorm.DB {
	// 	return db.Where("name = ?", "WAWA")
	// }).First(&girl)

	// fmt.Println(girl)

	//GLOBAL_DB.AutoMigrate(&Info{})

	var girl GirlGod
	//GLOBAL_DB.Preload("Dogs.Info").Preload("Dogs").First(&girl)
	// GLOBAL_DB.Preload("Dogs.Info").First(&girl)
	// GLOBAL_DB.Preload("Dogs.Info").Preload("Dogs","name = ?","WAWA").First(&girl)
	// GLOBAL_DB.Preload("Dogs.Info","money>100").Preload("Dogs","name = ?","WAWA").First(&girl)
	// WRONG :GLOBAL_DB.Preload("Dogs.Info","name = ?","WAWA").First(&girl)

	GLOBAL_DB.Preload("Dogs", func(db *gorm.DB) *gorm.DB {
		return db.Joins("Info").Where("money >200")
	}).First(&girl)

	fmt.Println(girl)

}
