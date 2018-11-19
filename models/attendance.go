package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Attendance struct {
	gorm.Model
	UserID       int `gorm:"not null"`
	User         User
	AttendanceAt time.Time
	LeavingAt    time.Time
}
