package text

import (
	"fmt"
	"os"

	"github.com/clin003/tgbot_app_dev/features"
	tele "gopkg.in/telebot.v3"
)

var replyText, replyBtn1Text, replyBtn1Url, replyBtn2Text, replyBtn1Url string

func init() {
	features.RegisterFeature(tele.OnText, OnChannelLinkGroup)

	replyText = os.Getenv("BAICAI_BOT_REPLY_TEXT")
	replyBtn1Text = os.Getenv("BAICAI_BOT_REPLY_BTN1_TEXT")
	replyBtn1Url = os.Getenv("BAICAI_BOT_REPLY_BTN1_URL")
	replyBtn2Text = os.Getenv("BAICAI_BOT_REPLY_BTN2_TEXT")
	replyBtn1Url = os.Getenv("BAICAI_BOT_REPLY_BTN2_URL")
}

// Command: /start <PAYLOAD>
func OnChannelLinkGroup(c tele.Context) error {
	if len(replyText) <= 0 {
		return nil
	}
	// fmt.Println("OnChannelLinkGroup", 0)
	if c.Message().Private() ||
		c.Message().FromChannel() ||
		c.Message().IsReply() {
		return nil
	}
	// fmt.Println("OnChannelLinkGroup", 1)
	if !(c.Message().OriginalChat != nil) || !(c.Message().SenderChat != nil) {
		return nil
	}
	// fmt.Println("OnChannelLinkGroup", 2)
	if c.Message().OriginalChat.Type != tele.ChatChannel ||
		c.Message().SenderChat.Type != tele.ChatChannel ||
		!c.Message().FromGroup() {
		return nil
	}
	// fmt.Println("OnChannelLinkGroup", 3)
	// menu := &tele.ReplyMarkup{ResizeKeyboard: true}
	// btn1 := menu.URL("薅羊毛📦", "https://t.me/haowu_push")
	// btn2 := menu.URL("值得买🔥", "https://t.me/haowu_dw")
	// menu.Reply(
	// 	menu.Row(btn1),
	// 	menu.Row(btn2),
	// )
	selector := &tele.ReplyMarkup{}
	if len(replyBtn1Text) <= 0 {
		replyBtn1Text = "薅羊毛📦"
	}
	if len(replyBtn1Url) <= 0 {
		replyBtn1Url = "https://t.me/haowu_push"
	}
	if len(replyBtn2Text) <= 0 {
		replyBtn2Text = "值得买🔥"
	}
	if len(replyBtn2Url) <= 0 {
		replyBtn2Url = "https://t.me/haowu_dw"
	}
	btnPrev := selector.URL(replyBtn1Text, replyBtn1Url)
	btnNext := selector.URL(replyBtn2Text, replyBtn2Url)
	selector.Inline(
		selector.Row(
			btnPrev,
			btnNext,
		),
	)
	// c.Reply("评论区请友好👬发言selector", selector)
	// c.Reply("评论区请友好👬发言menu", menu)
	return c.Reply(replyText, selector)
}
