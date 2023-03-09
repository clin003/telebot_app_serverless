package checkhealth

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/clin003/tgbot_app_dev/common"
	"github.com/clin003/tgbot_app_dev/features"
	_ "github.com/clin003/tgbot_app_dev/main/distro/all"

	tele "gopkg.in/telebot.v3"
)

var (
	bot *tele.Bot
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if err := autoSend(); err != nil {
		log.Println("autoSend err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		// return
	}
	log.Println(r.UserAgent())
}

func init() {
	var err error
	botToken := os.Getenv("BAICAI_BOT_TELEGRAM_TOKEN")
	bot, err = tele.NewBot(tele.Settings{
		Token:       botToken,
		Synchronous: true,
	})
	common.Must(err)

	features.Handle(bot)
}
func autoSend() error {
	reciverIdStr := os.Getenv("BAICAI_BOT_TELEGRAM_AUTOSEND_CHAT_ID")
	reciverId, err := strconv.ParseInt(reciverIdStr, 10, 64)
	if err != nil {
		return err
	}
	botVideoAPI := os.Getenv("BAICAI_BOT_VIDEO_API")
	if len(botVideoAPI) <= 0 || !strings.HasPrefix(botVideoAPI, "http") {
		err := fmt.Errorf("BAICAI_BOT_VIDEO_API Is NULL")
		return err
	}
	botVideoAPI = botVideoAPI + "?_t=" + fmt.Sprintf("%d", time.Now().UnixNano()/1e6) //13位时间戳
	log.Println(botVideoAPI)
	caption := os.Getenv("BAICAI_BOT_TELEGRAM_AUTOSEND_CAPTION")

	m := new(tele.Video)
	m.File = tele.FromURL(botVideoAPI)
	if len(caption) > 0 {
		m.Caption = caption
	}
	return SendMessage(reciverId, m)
}

func SendMessage(reciverId int64, m interface{}) error {
	reciver := &tele.User{
		ID: reciverId, //int64(reciverId),
	}

	if _, err := bot.Send(reciver, m); err != nil {
		log.Printf("Send(%d,%v) Msg Error: %v", reciverId, m, err)
		return errors.New("send message failed,")
	}

	return nil
}
