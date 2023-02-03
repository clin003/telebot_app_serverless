package main

import (
	"os"
	"time"

	"github.com/clin003/tgbot_app_dev/common"
	"github.com/clin003/tgbot_app_dev/features"
	_ "github.com/clin003/tgbot_app_dev/main/distro/all"

	tele "gopkg.in/telebot.v3"
)

func main() {
	b, err := tele.NewBot(tele.Settings{
		Token:  os.Getenv("BAICAI_BOT_TELEGRAM_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	common.Must(err)

	features.Handle(b)

	commands := []tele.Command{
		{
			Text:        "/id",
			Description: "Getid",
		},
		{
			Text:        "/ping",
			Description: "Ping",
		},
		{
			Text:        "/about",
			Description: "About",
		},
		{
			Text:        "/start",
			Description: "Start",
		},
	}
	b.SetCommands(commands)
	b.Start()
}
