package text

import (
	"fmt"

	"github.com/clin003/tgbot_app_dev/features"

	tele "gopkg.in/telebot.v3"
)

// Command: /start <PAYLOAD>
func OnText(c tele.Context) error {
	fmt.Println("OnText", 0)
	if c.Message().Private() ||
		c.Message().FromChannel() ||
		c.Message().IsReply() {
		return nil
	}
	fmt.Println("OnText", 1)
	if !(c.Message().OriginalChat != nil) || !(c.Message().SenderChat != nil) {
		return nil
	}
	fmt.Println("OnText", 2)
	if c.Message().OriginalChat.Type != tele.ChatChannel ||
		c.Message().SenderChat.Type != tele.ChatChannel ||
		!c.Message().FromGroup() {
		return nil
	}
	fmt.Println("OnText", 3)
	menu := &tele.ReplyMarkup{ResizeKeyboard: true}
	btn1 := menu.URL("薅羊毛📦", "https://t.me/haowu_push")
	btn2 := menu.URL("值得买🔥", "https://t.me/haowu_dw")
	menu.Reply(
		menu.Row(btn1),
		menu.Row(btn2),
	)
	selector := &tele.ReplyMarkup{}
	btnPrev := selector.URL("薅羊毛📦", "https://t.me/haowu_push")
	btnNext := selector.URL("值得买🔥", "https://t.me/haowu_dw")
	selector.Inline(
		selector.Row(btnPrev, btnNext),
	)
	c.Reply("评论区请友好👬发言selector", selector)
	c.Reply("评论区请友好👬发言menu", menu)
	menu.ResizeKeyboard = false
	return c.Reply("评论区请友好👬发言", menu)
}

func init() {
	features.RegisterFeature(tele.OnText, OnText)
}
