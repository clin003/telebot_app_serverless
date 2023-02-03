package start

import (
	"fmt"

	"github.com/clin003/tgbot_app_dev/features"

	tele "gopkg.in/telebot.v3"
)

// Command: /start <PAYLOAD>
func Onstart(c tele.Context) error {
	if !c.Message().Private() {
		return nil
	}

	helloStr := fmt.Sprintf("Hello! %s(%d)", c.Message().Sender.Username, c.Message().Sender.ID)
	return c.Send(helloStr)
}

func init() {
	features.RegisterFeature("/start", Onstart)
}
