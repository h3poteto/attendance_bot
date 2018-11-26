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
	leaving := AlreadyLeaving(user)
	if !leaving {
		return nil, errors.New("Still working")
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

func FinishAttendance(SlackID string) (*Attendance, error) {
	user, err := FindOrCreateUser(SlackID)
	if err != nil {
		return nil, err
	}
	attendance, err := FindLastAttendance(user)
	if err != nil {
		return nil, err
	}
	if attendance.LeavingAt != nil {
		return nil, errors.New("Already leaving")
	}
	timestamp := time.Now()
	attendance.LeavingAt = &timestamp
	db := config.SharedDB()
	db.Save(attendance)
	return attendance, nil
}

func FindLastAttendance(user *User) (*Attendance, error) {
	attendance := Attendance{}
	db := config.SharedDB()
	db.Where("user_id = ?", user.ID).Order("attendance_at desc").First(&attendance)
	if attendance.ID == 0 {
		return nil, errors.New("Could not find attendance.")
	}
	return &attendance, nil
}

func AlreadyLeaving(user *User) bool {
	attendance, err := FindLastAttendance(user)
	if err != nil {
		return true
	}
	if attendance.LeavingAt == nil {
		return false
	}
	return true
}
