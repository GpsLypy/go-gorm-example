package main

import "gorm.io/gorm"

// type Dog struct {
// 	gorm.Model
// 	Name string
// }

// type GirlGod struct{
// 	gorm.Model
// 	Name string
// }

// func One2one(){
// 	GLOBAL_DB.AutoMigrate(&Dog{},&GirlGod{})
// }

//belong to
// type Dog struct {
// 	gorm.Model
// 	Name      string
// 	GirlGodID uint //舔狗心里有女神了(属于)
// 	GirlGod   GirlGod
// }

// type GirlGod struct {
// 	gorm.Model
// 	Name string
// 	//Dog Dog //女神拥有舔狗
// }

// func One2one() {
// 	GLOBAL_DB.AutoMigrate(&Dog{})
// }

// type Dog struct {
// 	gorm.Model
// 	Name      string
// 	GirlGodID uint //舔狗心里有女神了(属于)
// 	GirlGod   GirlGod
// }

// type GirlGod struct {
// 	gorm.Model
// 	Name string
// 	//Dog Dog //女神拥有舔狗
// }

// func One2one() {
// 	g := GirlGod{
// 		Model: gorm.Model{
// 			ID: 1,
// 		},
// 		Name: "喵喵",
// 	}
// 	d := Dog{
// 		Model: gorm.Model{
// 			ID: 1,
// 		},
// 		Name:    "奇奇",
// 		GirlGod: g,
// 	}

// 	GLOBAL_DB.Create(&d)
// 	//GLOBAL_DB.AutoMigrate(&Dog{})
// }

// type Dog struct {
// 	gorm.Model
// 	Name      string
// 	GirlGodID uint //舔狗心里有女神了(属于)
// }

// type GirlGod struct {
// 	gorm.Model
// 	Name string
// 	Dog  Dog //has one
// }

// func One2one() {
// 	// d := Dog{
// 	// 	Model: gorm.Model{
// 	// 		ID: 1,
// 	// 	},
// 	// 	Name:    "奇奇",
// 	// }
// 	// g := GirlGod{
// 	// 	Model: gorm.Model{
// 	// 		ID: 1,
// 	// 	},
// 	// 	Name: "喵喵",
// 	// 	Dog:d,
// 	// }

// 	//GLOBAL_DB.Create(&d)
// 	GLOBAL_DB.AutoMigrate(&GirlGod{},&Dog{})
// }

// func One2one() {
// 	// d := Dog{
// 	// 	Model: gorm.Model{
// 	// 		ID: 1,
// 	// 	},
// 	// 	Name:    "奇奇",
// 	// }
// 	g := GirlGod{
// 		Name: "喵喵",
// 	}

// 	//create 女神带出一只狗
// 	GLOBAL_DB.Create(&g)
// 	//GLOBAL_DB.AutoMigrate(&GirlGod{},&Dog{})
// }

// func One2one() {
// 	d := Dog{
// 		Name: "奇奇",
// 	}
// 	g := GirlGod{
// 		Name: "喵喵",
// 		Dog:  d,
// 	}

// 	//create 女神带出一只狗
// 	GLOBAL_DB.Create(&g)
// 	//GLOBAL_DB.AutoMigrate(&GirlGod{},&Dog{})
// }

// func One2one() {
// 	// d := Dog{
// 	// 	Name: "奇奇",
// 	// }
// 	// g := GirlGod{
// 	// 	Name: "喵喵",
// 	// 	Dog:  d,
// 	// }
// 	var girl GirlGod
// 	//create 女神带出一只狗
// 	GLOBAL_DB.First(&girl, 2)
// 	GLOBAL_DB.Preload("Dog").First(&girl, 2)

// 	//GLOBAL_DB.AutoMigrate(&GirlGod{},&Dog{})
// }

//belong to
// type Dog struct {
// 	gorm.Model
// 	Name      string
// 	GirlGodID uint //舔狗心里有女神了(属于)
// 	GirlGod   GirlGod
// }

// type GirlGod struct {
// 	gorm.Model
// 	Name string
// 	//Dog Dog //女神拥有舔狗
// }

// func One2one(){
// 	// d :=Dog{
// 	// Name: ,
// 	// }
// 	var dog Dog

// 	GLOBAL_DB.Preload("GirlGod").First(&dog,1)
// }

// func One2one() {
// 	d := Dog{
// 		Model: gorm.Model{
// 			ID: 1,
// 		},
// 	}
// 	g := GirlGod{
// 		Model: gorm.Model{
// 			ID: 1,
// 		},
// 	}
// g2 := GirlGod{
// 	Model: gorm.Model{
// 		ID: 2,
// 	},
//}
// 	//创建一个舔狗模型，添女神
// 	GLOBAL_DB.Model(&d).Association("GirlGod").Append(&g)
// 	GLOBAL_DB.Model(&d).Association("GirlGod").Delete(&g)
// 	GLOBAL_DB.Model(&d).Association("GirlGod").Replace(&g2)
// 	GLOBAL_DB.Model(&d).Association("GirlGod").Clear()
// }

//has one
type Dog struct {
	gorm.Model
	Name      string
	GirlGodID uint //舔狗心里有女神了(属于)
	//GirlGod   GirlGod
}

type GirlGod struct {
	gorm.Model
	Name string
	Dog  Dog //女神拥有舔狗
}

func One2one() {
	d := Dog{
		Model: gorm.Model{
			ID: 1,
		},
	}

	d2 := Dog{
		Model: gorm.Model{
			ID: 2,
		},
	}
	g := GirlGod{
		Model: gorm.Model{
			ID: 1,
		},
	}
	// g2 := GirlGod{
	// 	Model: gorm.Model{
	// 		ID: 2,
	// 	},
	// }
	//创建一个舔狗模型，添女神
	GLOBAL_DB.Model(&g).Association("Dog").Append(&d)
	GLOBAL_DB.Model(&g).Association("Dog").Delete(&g)
	GLOBAL_DB.Model(&g).Association("Dog").Replace(&d, &d2)
	GLOBAL_DB.Model(&g).Association("Dog").Clear()
}
