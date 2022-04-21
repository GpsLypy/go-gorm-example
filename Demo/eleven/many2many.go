package main

import "gorm.io/gorm"

type Info struct {
	gorm.Model
	Money int
	DogID uint
}

type Dog struct {
	gorm.Model
	Name      string
	GirlGodID uint //舔狗心里有女神了(属于)
	//GirlGod   GirlGod
	Info    Info
	GirlGod []GirlGod `gorm:"many2many:dog_girl_god"`
}

type GirlGod struct {
	gorm.Model
	Name string
	Dog  []Dog `gorm:"many2many:dog_girl_god"` //女神拥有舔狗
}

func One2one() {
	i := Info{
		Money: 20000,
	}
	g := GirlGod{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "nv1",
	}
	g2 := GirlGod{
		Model: gorm.Model{
			ID: 2,
		},
		Name: "nv2",
	}
	d := Dog{
		// Model: gorm.Model{
		// 	ID: 1,
		// },
		Name: "wawa",
		Info: i,
		GirlGod: []GirlGod{
			g,
			g2,
		},
	}

	// d2 := Dog{
	// 	// Model: gorm.Model{
	// 	// 	ID: 2,
	// 	// },
	// 	Name: "wawa2",
	// }
	GLOBAL_DB.Create(&d)

}

func Find() {
	d := Dog{
		Model: gorm.Model{
			ID: 1,
		},
	}
	var girls []GirlGod
	GLOBAL_DB.Model(&d).Association("GrilGods").Find(&girls)
	GLOBAL_DB.Model(&d).Preload("Dogs").Association("GrilGods").Find(&girls)
	GLOBAL_DB.Model(&d).Preload("Dogs.Info").Association("GrilGods").Find(&girls)
	GLOBAL_DB.Model(&d).Preload("Dogs", func(db *gorm.DB) *gorm.DB {
		return db.Joins("Info").Where("money < ?", 10000)
	}).Association("GrilGods").Find(&girls)

}

func Add() {
	d := Dog{
		Model: gorm.Model{
			ID: 3,
		},
	}
	g := GirlGod{
		Model: gorm.Model{
			ID: 1,
		},
	}
	g2 := GirlGod{
		Model: gorm.Model{
			ID: 2,
		},
	}
	//创建一个舔狗模型，添女神
	GLOBAL_DB.Model(&d).Association("GirlGod").Append(&g, &g2)
	// GLOBAL_DB.Model(&g).Association("Dog").Delete(&g)
	// GLOBAL_DB.Model(&g).Association("Dog").Replace(&d, &d2)
	// GLOBAL_DB.Model(&g).Association("Dog").Clear()
}
