package menu

import (
	"fmt"

	"github.com/clin003/tgbot_app_dev/features"

	tele "gopkg.in/telebot.v3"
)

// https://pkg.go.dev/gopkg.in/tucnak/telebot.v3#section-readme
var (
	// Universal markup builders.
	menu = &tele.ReplyMarkup{ResizeKeyboard: true}
	// selector = &tele.ReplyMarkup{}

	// Reply buttons.
	btnHelp     = menu.Text("ℹ Help")
	btnSettings = menu.Text("⚙ Settings")

	// Inline buttons.
	//
	// Pressing it will cause the client to
	// send the bot a callback.
	//
	// Make sure Unique stays unique as per button kind,
	// as it has to be for callback routing to work.
	//
	// btnPrev = selector.Data("⬅", "prev", "")
	// btnNext = selector.Data("➡", "next", "...")
)

// Command: /start <PAYLOAD>
func OnMenu(c tele.Context) error {
	if !c.Message().Private() {
		return nil
	}

	helloStr := fmt.Sprintf("Hello! %s(%d)", c.Message().Sender.Username, c.Message().Sender.ID)

	// c.Bot().Send(c.Sender(), helloStr, menu)
	return c.Send(helloStr, menu)
}

func init() {
	menu.Reply(
		menu.Row(btnHelp),
		menu.Row(btnSettings),
	)
	// selector.Inline(
	// 	selector.Row(btnPrev, btnNext),
	// )
	features.RegisterFeature("/menu", OnMenu)
}
