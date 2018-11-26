package listener

import (
	"regexp"

	"github.com/h3poteto/attendance_bot/models"
	"github.com/nlopes/slack"
)

func (l *Listener) MessageHandler(message *slack.MessageEvent, rtm *slack.RTM) error {
	start := regexp.MustCompile(`おはー`)
	if start.MatchString(message.Text) {
		if _, err := models.StartAttendance(message.User); err != nil {
			rtm.SendMessage(rtm.NewOutgoingMessage("失敗したよ！すでに出勤してない？", message.Channel))
			return err
		}
		rtm.SendMessage(rtm.NewOutgoingMessage("おはー 打刻したよー", message.Channel))
		return nil
	}
	end := regexp.MustCompile(`店じまい`)
	if end.MatchString(message.Text) {
		if _, err := models.FinishAttendance(message.User); err != nil {
			rtm.SendMessage(rtm.NewOutgoingMessage("失敗したよ！すでに退勤してない？", message.Channel))
			return err
		}
		rtm.SendMessage(rtm.NewOutgoingMessage("おつー 打刻したよー", message.Channel))
		return nil
	}
	return nil
}
