package listener

import (
	"github.com/nlopes/slack"
	log "github.com/sirupsen/logrus"
)

type Listener struct {
	client *slack.Client
}

func NewListener(token string) *Listener {
	client := slack.New(token)
	return &Listener{
		client,
	}
}

func (l *Listener) Listen() {
	rtm := l.client.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.HelloEvent:
			// Ignore hello
		case *slack.ConnectedEvent:
			log.Info("Infos:", ev.Info)
			log.Info("Connection counter:", ev.ConnectionCount)
			// Replace C2147483705 with your Channel ID
			rtm.SendMessage(rtm.NewOutgoingMessage("Hello world", "C2147483705"))

		case *slack.MessageEvent:
			err := l.MessageHandler(ev, rtm)
			if err != nil {
				log.Error(err)
			}

		case *slack.PresenceChangeEvent:
			log.Infof("Presence Change: %v\n", ev)

		case *slack.LatencyReport:
			// Ignore latency
		case *slack.RTMError:
			log.Errorf("Error: %s", ev.Error())

		case *slack.InvalidAuthEvent:
			log.Info("Invalid credentials")
			return

		default:

			// Ignore other events..
			// fmt.Printf("Unexpected: %v\n", msg.Data)
		}
	}
}
