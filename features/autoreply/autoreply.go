package autoreply

import (
	"fmt"
	"os"
	"sync"

	"github.com/clin003/tgbot_app_dev/features"
	tele "gopkg.in/telebot.v3"
)

var replyText, replyBtn1Text, replyBtn1Url, replyBtn2Text, replyBtn2Url, replyBtn3Text, replyBtn3Url string
var syncMap sync.Map

func init() {
	features.RegisterFeature(tele.OnText, OnChannelLinkGroup)
	features.RegisterFeature(tele.OnPhoto, OnChannelLinkGroup)
	features.RegisterFeature(tele.OnVideo, OnChannelLinkGroup)
	features.RegisterFeature(tele.OnMedia, OnChannelLinkGroup)

	replyText = os.Getenv("BAICAI_BOT_REPLY_TEXT")
	replyBtn1Text = os.Getenv("BAICAI_BOT_REPLY_BTN1_TEXT")
	replyBtn1Url = os.Getenv("BAICAI_BOT_REPLY_BTN1_URL")
	replyBtn2Text = os.Getenv("BAICAI_BOT_REPLY_BTN2_TEXT")
	replyBtn2Url = os.Getenv("BAICAI_BOT_REPLY_BTN2_URL")
	replyBtn3Text = os.Getenv("BAICAI_BOT_REPLY_BTN3_TEXT")
	replyBtn3Url = os.Getenv("BAICAI_BOT_REPLY_BTN3_URL")
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
	msgId := ""
	if len(c.Message().AlbumID) > 0 {
		msgId = fmt.Sprintf("%d_%s", c.Message().Chat.ID, c.Message().AlbumID)
	} else { // c.Message().ID > 0
		msgId = msgId + fmt.Sprintf("%d_%d", c.Message().Chat.ID, c.Message().ID)
	}
	// else if c.Message().OriginalUnixtime > 0 {
	// 	msgId = msgId + fmt.Sprintf("%d_%d", c.Message().Chat.ID, c.Message().OriginalUnixtime)
	// }
	// else {
	// 	msgId = msgId + fmt.Sprintf("%d", c.Update().ID)
	// }
	if _, ok := syncMap.LoadOrStore(msgId, ""); ok {
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
	btnList := make([]tele.Btn, 0)
	if len(replyBtn1Text) <= 0 || len(replyBtn1Url) <= 0 {
		replyBtn1Text = "薅羊毛📦"
		replyBtn1Url = "https://t.me/haowu_push"
	}
	btnPrev := selector.URL(replyBtn1Text, replyBtn1Url)
	btnList = append(btnList, btnPrev)

	if len(replyBtn2Text) <= 0 || len(replyBtn2Url) <= 0 {
		replyBtn2Text = "值得买🔥"
		replyBtn2Url = "https://t.me/haowu_dw"
	}
	btnNext := selector.URL(replyBtn2Text, replyBtn2Url)
	btnList = append(btnList, btnNext)

	if len(replyBtn3Text) > 0 && len(replyBtn3Url) > 0 {
		// replyBtn3Text = "唯品会💃"
		// replyBtn3Url = "https://t.me/haowu_vip"
		btn3 := selector.URL(replyBtn3Text, replyBtn3Url)
		btnList = append(btnList, btn3)
	}

	selector.Inline(
		selector.Row(
			btnList...,
		// btnPrev,
		// btnNext,
		// btn3,
		),
	)
	// c.Reply("评论区请友好👬发言selector", selector)
	// c.Reply("评论区请友好👬发言menu", menu)
	return c.Reply(replyText, selector)
}
