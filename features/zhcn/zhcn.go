package zhcn

import (
	"github.com/clin003/tgbot_app_dev/features"
	tele "gopkg.in/telebot.v3"
)

func init() {
	features.RegisterFeature("/zhcn", OnZhCN)
}

// Command: /start <PAYLOAD>
func OnZhCN(c tele.Context) error {
	replyText := "点击下方按钮更换中文🔘"
	selector := &tele.ReplyMarkup{}
	btnList := make([]tele.Btn, 0)
	btnZhCN := selector.URL("中文包", "https://t.me/setlanguage/zhcncc")
	btnList = append(btnList, btnZhCN)

	selector.Inline(
		selector.Row(
			btnList...,
		),
	)
	return c.Reply(replyText, selector)
}
