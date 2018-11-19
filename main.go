package main

import (
	"os"

	"github.com/h3poteto/attendance_bot/config"
	"github.com/h3poteto/attendance_bot/listener"
	"github.com/h3poteto/attendance_bot/models"
)

func init() {
	db := config.SharedDB()
	db.AutoMigrate(&models.User{}, &models.Attendance{})
}

func main() {
	token := os.Getenv("SLACK_TOKEN")
	l := listener.NewListener(token)
	l.Listen()
}
