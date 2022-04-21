package main

import (
	"database/sql"
	"time"

	_ "gorm.io/gorm"
)

//自定义Model
type Model struct {
	UUID uint      `gorm:"primaryKey"`
	Time time.Time `gorm:"column:timer"`
}

type TestUser struct {
	//ID           uint
	// ActivatedAt  sql.NullTime
	// CreatedAt    time.Time
	// UpdatedAt    time.Time
	//gorm.Model

	Model        Model   `gorm:"embedded;embeddedPrefix:qm_"`
	Name         string  `gorm:"default:qm"`
	Email        *string `gorm:"not null"`
	Age          uint8   `gorm:"comment:年龄"`
	Birthday     *time.Time
	MemberNumber sql.NullString
}

func TestUserCreate() {
	GLOBAL_DB.AutoMigrate(&TestUser{})
}
