package main

import (
	"os"

	"github.com/h3poteto/attendance_bot/listener"
)

func main() {
	token := os.Getenv("SLACK_TOKEN")
	l := listener.NewListener(token)
	l.Listen()
}
