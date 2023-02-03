package getid

import (
	"fmt"

	"github.com/clin003/tgbot_app_dev/features"

	tele "gopkg.in/telebot.v3"
)

func OnGetID(c tele.Context) error {
	text := ""
	text = fmt.Sprintf("@%s(%d)",
		c.Message().Sender.Username, c.Message().Sender.ID,
	)
	if c.Message().FromGroup() {
		if len(c.Message().Chat.Username) > 0 {
			text = fmt.Sprintf("%s\n%s @%s(%d)", text,
				c.Message().Chat.Title, c.Message().Chat.Username, c.Message().Chat.ID,
			)
		} else {
			text = fmt.Sprintf("%s\n%s(%d)", text,
				c.Message().Chat.Title, c.Message().Chat.ID,
			)
		}
	}

	if !c.Message().Private() {
		c.Bot().Send(c.Sender(), text)
		return c.Reply(text)
	}
	return c.Send(text)
}

func init() {
	features.RegisterFeature("/id", OnGetID)
}
