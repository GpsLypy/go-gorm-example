package main

// type Jiazi struct {
// 	ID          uint
// 	Name        string
// 	Xiaofengche []Xiaofengche `gorm:"foreignKey:JiaziName;references:Name"`
// }

// type Xiaofengche struct {
// 	ID   uint
// 	Name string
// 	//JiaziID uint
// 	JiaziName string
// }

// func Polymorphic() {
// 	GLOBAL_DB.AutoMigrate(&Jiazi{}, &Xiaofengche{})
// 	GLOBAL_DB.Create(&Jiazi{Name: "jiazi", Xiaofengche: []Xiaofengche{
// 		{Name: "小风车"},
// 		{Name: "小风车1"},
// 	}})
// }
