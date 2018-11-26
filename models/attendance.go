package models

import (
	"time"

	"github.com/h3poteto/attendance_bot/config"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type Attendance struct {
	gorm.Model
	UserID       int `gorm:"not null"`
	User         User
	AttendanceAt *time.Time
	LeavingAt    *time.Time
}

func StartAttendance(SlackID string) (*Attendance, error) {
	user, err := FindOrCreateUser(SlackID)
	if err != nil {
		return nil, err
	}
	timestamp := time.Now()
	attendance := &Attendance{
		User:         *user,
		AttendanceAt: &timestamp,
		LeavingAt:    nil,
	}
	db := config.SharedDB()
	db.Create(attendance)
	if db.NewRecord(*attendance) {
		return nil, errors.New("Could not create attendance record.")
	}
	return attendance, nil
}
