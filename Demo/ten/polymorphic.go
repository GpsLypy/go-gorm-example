package main

type Jiazi struct {
	ID          uint
	Name        string
	Xiaofengche Xiaofengche `gorm:"polymorphic:Owner;polymorphicValue:dudu"`
}

type Yujie struct {
	ID          uint
	Name        string
	Xiaofengche Xiaofengche `gorm:"polymorphic:Owner;polymorphicValue:huhu"`
}

type Xiaofengche struct {
	ID        uint
	Name      string
	OwnerType string
	OwnerID   uint
}

func Polymorphic() {
	GLOBAL_DB.AutoMigrate(&Jiazi{}, &Yujie{}, &Xiaofengche{})
}

func Create() {
	GLOBAL_DB.Create(&Jiazi{Name: "jiazi", Xiaofengche: Xiaofengche{
		Name: "小风车",
	}})

	GLOBAL_DB.Create(&Yujie{Name: "yujie", Xiaofengche: Xiaofengche{
		Name: "大风车",
	}})
}
