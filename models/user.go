package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	SlackID int `gorm:"unique:not null"`
}
