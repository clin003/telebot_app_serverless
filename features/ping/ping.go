package ping

import (
	"fmt"

	"github.com/clin003/tgbot_app_dev/features"

	tele "gopkg.in/telebot.v3"
)

func OnPing(c tele.Context) error {
	if !c.Message().Private() {
		// c.Delete()
		return c.Reply("pong")
	}
	text := fmt.Sprintf("Pong! %s%s @%s(%d)",
		c.Message().Sender.FirstName, c.Message().Sender.LastName, c.Message().Sender.Username, c.Message().Sender.ID)

	c.Delete()
	// return c.Reply(text)
	return c.Send(text)
}

func init() {
	features.RegisterFeature("/ping", OnPing)
}
