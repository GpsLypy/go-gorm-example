package main

type Jiazi struct {
	ID          uint
	Name        string
	Xiaofengche []Xiaofengche `gorm:"many2many:jiazi_fengche;foreignKey:Name;references:FCName"`
}

type Xiaofengche struct {
	ID     uint
	FCName string
	//JiaziID uint
	//JiaziName string
}

func Polymorphic2() {
	GLOBAL_DB.AutoMigrate(&Jiazi{}, &Xiaofengche{})
	GLOBAL_DB.Create(&Jiazi{Name: "jiazi", Xiaofengche: []Xiaofengche{
		{FCName: "小风车"},
		{FCName: "小风车1"},
	}})
}
