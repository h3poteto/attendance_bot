package models

import (
	"github.com/h3poteto/attendance_bot/config"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type User struct {
	gorm.Model
	SlackID string `gorm:"unique:not null"`
}

func FindUser(SlackID string) (*User, error) {
	db := config.SharedDB()
	user := &User{
		SlackID: SlackID,
	}
	db.First(&user)
	if user.ID == 0 {
		return nil, errors.New("Could not find.")
	}
	return user, nil
}

func FindOrCreateUser(SlackID string) (*User, error) {
	user, err := FindUser(SlackID)
	if err != nil {
		// If User does not exist, we must create a user based on SlackID.
		user = &User{
			SlackID: SlackID,
		}
		db := config.SharedDB()
		db.Create(user)
		if db.NewRecord(*user) {
			return nil, errors.New("Could not create.")
		}
		return user, nil
	}
	return user, nil
}
