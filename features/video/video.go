package video

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/clin003/tgbot_app_dev/features"

	tele "gopkg.in/telebot.v3"
)

func OnVideo(c tele.Context) error {
	botVideoAPI := os.Getenv("BAICAI_BOT_VIDEO_API")
	botVideoCaption := os.Getenv("BAICAI_BOT_VIDEO_CAPTION")
	botVideoAPI = botVideoAPI + "?_t=" + fmt.Sprintf("%d", time.Now().UnixNano()/1e6) //13位时间戳
	log.Println(botVideoAPI)

	m := new(tele.Video)
	m.File = tele.FromURL(botVideoAPI)

	// c.Delete()
	// return c.Reply(text)
	if c.Message().Private() {
		// c.Delete()
		// return c.Reply(m)
		if err := c.Reply(m); err != nil {
			botVideoAPI = botVideoAPI + "1"
			m := new(tele.Video)
			m.File = tele.FromURL(botVideoAPI)
			return c.Reply(m)
		} else {
			return nil
		}
	}
	c.Delete()
	if len(botVideoCaption) > 0 {
		m.Caption = botVideoCaption
	}
	if err := c.Send(m, tele.ModeMarkdown); err != nil {
		botVideoAPI = botVideoAPI + "1"
		m := new(tele.Video)
		m.File = tele.FromURL(botVideoAPI)
		return c.Send(m, tele.ModeMarkdown)
	} else {
		return nil
	}
}

func init() {
	features.RegisterFeature("/video", OnVideo)
}
