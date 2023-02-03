package about

import (
	"github.com/clin003/tgbot_app_dev/baicai"
	"github.com/clin003/tgbot_app_dev/features"

	tele "gopkg.in/telebot.v3"
)

func OnProcess(c tele.Context) error {
	text := baicai.About()
	return c.Reply(text)
}
func OnVersion(c tele.Context) error {
	text := baicai.Version()
	return c.Reply(text)
}

func init() {
	features.RegisterFeature("/about", OnProcess)
	features.RegisterFeature("/version", OnVersion)
}
