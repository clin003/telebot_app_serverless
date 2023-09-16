package info

import (
	"fmt"
	"strings"

	"github.com/clin003/tgbot_app_dev/features"

	tele "gopkg.in/telebot.v3"
)

func init() {
	features.RegisterFeature("/info", OnInfo)
}
func OnInfo(c tele.Context) error {
	if !c.Message().Private() {
		return nil
	}
	// c.Bot().ChatByID()
	payload := c.Message().Payload
	chatUsername := ""
	if strings.HasPrefix(payload, "@") {
		chatUsername = payload
	} else if strings.HasPrefix(payload, "https://t.me/") {
		if chatPath, isFound := strings.CutPrefix(payload, "https://t.me/"); isFound && len(chatPath) > 0 {
			if chatName, _, isFound := strings.Cut(chatPath, "/"); isFound && len(chatName) > 0 {
				chatUsername = "@" + chatName
			} else {
				chatUsername = "@" + chatPath
			}
		} else {
			return c.Reply("格式错误🙅:" + payload)
		}
	}

	chat, err := c.Bot().ChatByUsername(chatUsername)
	if err != nil {
		return c.Reply(err.Error())
	}

	text := fmt.Sprintf("%s id: %d", chat.Type, chat.ID)
	if chat.LinkedChatID != 0 {
		text = text + "\n" + fmt.Sprintf("连接群id: %d", chat.LinkedChatID)
	}
	text = text + "\n" + chat.Title
	if len(chat.Bio) > 0 {
		text = text + "\n" + "Bio: " + chat.Bio
	}
	if len(chat.Description) > 0 {
		text = text + "\n" + "Description: " + chat.Description
	}

	if len(chat.Username) > 0 {
		text = text + "\n" + "@" + chat.Username
	}
	return c.Reply(text)
}
