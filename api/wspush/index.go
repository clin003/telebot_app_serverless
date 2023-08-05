package wspush

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"gitee.com/lyhuilin/util"
	"github.com/clin003/tgbot_app_dev/common"
	"github.com/clin003/tgbot_app_dev/features"
	_ "github.com/clin003/tgbot_app_dev/main/distro/all"
	"github.com/clin003/tgbot_app_dev/utils"

	tele "gopkg.in/telebot.v3"
)

var (
	bot *tele.Bot
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		return
	}
	log.Println("r.URL.Path", r.URL.Path)
	// checkToken := r.Header.Get("Apiclient")
	// log.Println("checkToken", checkToken)
	// params := r.URL.Query()
	// userId, _ := strconv.Atoi(params.Get("to"))
	// msgText := params.Get("m")
	_, channelStr, ok := strings.Cut(r.URL.Path, "ws/push/")
	if len(channelStr) <= 0 || !ok {
		log.Println("收到非法推送:(r.URL.Path)", r.URL.Path)
		return
	}
	if ok := checkChannelSecretToken(channelStr); !ok {
		log.Println("收到非法推送:(token)", channelStr)
		return
	}

	body, err := io.ReadAll(r.Body)
	common.Must(err)
	// log.Println(string(body))
	// PushMsgData(channelStr, body)

	if err := PushMsgData(channelStr, body); err != nil {
		fmt.Fprintf(w, "err")
		return
	}
	fmt.Fprintf(w, "ok")
}
func checkChannelSecretToken(channelStr string) (retBool bool) {
	if len(channelStr) > 0 {
		liveTokenUUID := os.Getenv("BAICAI_WSPUSH_CHANNEL_" + channelStr)
		if channelStr == liveTokenUUID {
			retBool = true
		}
	}
	return
}

// 推送FeedMsg信息
func PushMsgData(botToken string, data []byte) error {
	// fmt.Printf("\n解析接收到的数据data：%s\n", data)
	var msg FeedRichMsgModel
	if err := json.Unmarshal(data, &msg); err != nil {
		log.Println(err, "解析数据失败:", string(data))
		return err
	}
	if len(msg.Image.PicURL) > 0 {
		//http://gchat.qpic.cn/gchatpic_new/0/0-0-036915B0DD5B458CE9989108B580E98D/0
		// https://c2cpicdw.qpic.cn/offpic_new/0/0BADBADBAD-610524638/0
		if strings.HasPrefix(msg.Image.PicURL, "http://gchat.qpic.cn") ||
			strings.Contains(msg.Image.PicURL, "qpic.cn/offpic_new") {
			// log.Debugf("图片采集失败：%s", msg.Image.PicURL)
			msg.Image.PicURL = ""
		}
	}
	if len(msg.Text.Content) > 0 {
		if err := feedKeyworldCheckEx(botToken, msg.Text.Content); err != nil {
			return err
		}
		msg.Text.Content = feedKeyworldReplaceEx(botToken, msg.Text.Content)
	}

	return pushMsgDataToTelegram(botToken, msg)
}

func feedKeyworldCheckEx(botToken, msgText string) error {
	feedKeyworldList := os.Getenv("BAICAI_WSPUSH_FEED_KEYWORLD_LIST_" + botToken)
	feedKeyworldFilter := os.Getenv("BAICAI_WSPUSH_FEED_KEYWORLD_FILTER_" + botToken)
	k, f, ok := utils.FeedKeyworldCheck(msgText, feedKeyworldFilter, feedKeyworldList)
	if !ok {
		return fmt.Errorf("关键词检查: 订阅词(%s) 屏蔽词(%s) 结果(%t) ", k, f, ok)
	}
	return nil
	// log.Printf("关键词检查: 订阅词(%s) 屏蔽词(%s) 结果(%t)", k, f, ok)
	// return ok
}

// 关键词替换
func feedKeyworldReplaceEx(botToken, msgText string) string {
	feedKeyworldReplace := os.Getenv("BAICAI_WSPUSH_FEED_KEYWORLD_REPLACE_" + botToken)
	return utils.FeedKeyworldReplace(msgText, feedKeyworldReplace)
}

func pushMsgDataToTelegram(botToken string, msg FeedRichMsgModel) error {
	// reciverId := msg.FeedChatId
	var reciverId int64
	reciverIdStr := os.Getenv("BAICAI_BOT_TELEGRAM_WSPUSH_FEED_CHAT_ID_" + botToken)
	if id, err := strconv.ParseInt(reciverIdStr, 10, 64); err != nil {
		// if msg.FeedChatId != 0 {
		// 	reciverId = msg.FeedChatId
		// } else {
		// 	return err
		// }
		return err
	} else {
		if id != 0 {
			reciverId = id
		} else {
			return errors.New("feed Chat ID is NULL,")
		}
	}
	switch msg.Msgtype {
	case "text":
		m := msg.Text.Content
		return SendMessage(reciverId, m, tele.ModeDefault)
	case "video":
		m := new(tele.Video)
		m.File = tele.FromURL(msg.Video.FileURL)
		if len(msg.Video.Caption) > 0 {
			m.Caption = msg.Video.Caption
		}
		if len(m.Caption) > 0 {
			if captionTmp, err := util.UrlRegMatchReplaceToTGHTML(m.Caption); err != nil {
			} else {
				m.Caption = captionTmp
			}
		}
		return SendMessage(reciverId, m, tele.ModeHTML)
	case "image":
		m := new(tele.Photo)
		m.File = tele.FromURL(msg.Image.PicURL)
		if len(msg.Image.Caption) > 0 {
			m.Caption = msg.Image.Caption
		}
		if len(m.Caption) > 0 {
			if captionTmp, err := util.UrlRegMatchReplaceToTGHTML(m.Caption); err != nil {
			} else {
				m.Caption = captionTmp
			}
		}
		return SendMessage(reciverId, m, tele.ModeHTML)
	case "rich":
		if len(msg.Image.PicURL) > 0 && strings.HasPrefix(msg.Image.PicURL, "http") {
			m := new(tele.Photo)
			m.File = tele.FromURL(msg.Image.PicURL)
			if len(msg.Image.Caption) > 0 {
				m.Caption = msg.Image.Caption
			} else if len(msg.Text.Content) > 0 {
				m.Caption = msg.Text.Content
			}
			if len(m.Caption) > 0 {
				if captionTmp, err := util.UrlRegMatchReplaceToTGHTML(m.Caption); err != nil {
				} else {
					m.Caption = captionTmp
				}
			}
			return SendMessage(reciverId, m, tele.ModeHTML)
		}
		if len(msg.Video.FileURL) > 0 && strings.HasPrefix(msg.Video.FileURL, "http") {
			m := new(tele.Video)
			m.File = tele.FromURL(msg.Video.FileURL)
			if len(msg.Video.Caption) > 0 {
				m.Caption = msg.Video.Caption
			} else if len(msg.Text.Content) > 0 {
				m.Caption = msg.Text.Content
			}
			if len(m.Caption) > 0 {
				if captionTmp, err := util.UrlRegMatchReplaceToTGHTML(m.Caption); err != nil {
				} else {
					m.Caption = captionTmp
				}
			}
			return SendMessage(reciverId, m, tele.ModeHTML)
		}
		if len(msg.Text.Content) > 0 {
			m := msg.Text.Content
			return SendMessage(reciverId, m, tele.ModeDefault)
		}
		return nil
	default:
		return errors.New("msg type is not support,")
	}
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

func SendMessage(reciverId int64, m interface{}, opts ...interface{}) error {
	reciver := &tele.User{
		ID: reciverId, //int64(reciverId),
	}

	if _, err := bot.Send(reciver, m, opts); err != nil {
		log.Printf("Send(%d,%v) Msg Error: %v", reciverId, m, err)
		return errors.New("send message failed,")
	}

	return nil
}

type FeedRichMsgModel struct {
	Msgtype string                `json:"msgtype"  form:"msgtype"` //rich text image video
	MsgID   string                `json:"msgID"  form:"msgID"`
	MsgTime string                `json:"msgTime"  form:"msgTime"`
	Text    FeedRichMsgTextModel  `json:"text"  form:"text"`
	Image   FeedRichMsgImageModel `json:"image"  form:"image"`
	Video   FeedRichMsgVideoModel `json:"video"  form:"video"`
	Link    string                `json:"link"  form:"link"`
	// (Optional)
	FeedChatId int64 `json:"feed_chat_id,omitempty"`
}
type FeedRichMsgTextModel struct {
	Content         string `json:"content"  form:"content"`
	ContentEx       string `json:"contentEx"  form:"contentEx"`
	ContentExPic    string `json:"contentExPic"  form:"contentExPic"`
	ContentMarkdown string `json:"contentMarkdown"  form:"contentMarkdown"`
}
type FeedRichMsgImageModel struct {
	PicURL   string `json:"picURL"  form:"picURL"`
	FilePath string `json:"filePath"  form:"filePath"`
	// (Optional)
	Caption string `json:"caption,omitempty"`
}
type FeedRichMsgVideoModel struct {
	FileURL  string `json:"fileURL"  form:"fileURL"`
	FilePath string `json:"filePath"  form:"filePath"`
	// (Optional)
	Caption string `json:"caption,omitempty"`
}

func (msg *FeedRichMsgModel) ToString() (res string) {
	res = fmt.Sprintf("msgID:%s,msgType:%s,msgTime:%s\n", msg.MsgID, msg.Msgtype, msg.MsgTime)
	if len(msg.Text.Content) > 0 {
		res = fmt.Sprintf("%s\n%s", res, msg.Text.Content)
	}
	if len(msg.Image.PicURL) > 0 {
		res = fmt.Sprintf("%s\n%s", res, msg.Image.PicURL)
	}
	if len(msg.Video.FileURL) > 0 {
		res = fmt.Sprintf("%s\n%s", res, msg.Video.FileURL)
	}

	return
}

// func feedKeyworldCheck(msgText string) (retText, retFilter string, retBool bool) {
// 	retBool = true

// 	keyworldFilter := utils.KeyworldListParse(getFeedKeyworldFilter())
// 	for _, v := range keyworldFilter {
// 		vc := v
// 		if strings.Contains(msgText, vc) {
// 			retFilter = vc
// 			retBool = false
// 			return
// 		}
// 	}

// 	keyworldList := utils.KeyworldListParse(getFeedKeyworldList())
// 	if len(keyworldList) <= 0 {
// 		retText = "无订阅词限定"
// 		retBool = true
// 		return
// 	} else {
// 		retBool = false
// 	}
// 	for _, v := range keyworldList {
// 		vc := v
// 		if strings.Contains(msgText, vc) {
// 			retText = vc
// 			retBool = true
// 			return
// 		}
// 	}
// 	return
// }
