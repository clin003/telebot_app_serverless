package baicai

const (
	APP_ABOUT = "Baicai TG BOT."
	APP_SRC   = "https://github.com/clin003/telebot_app_serverless"
)

func Version() string {
	return APP_VERSION
}

func Usage() string {
	return APP_ABOUT + "\n" + "交流TG群: @baicai_dev" + "\n" + "源码托管: " + APP_SRC
}
func About() string {
	text := "Baicai v" + APP_VERSION + "\n" + Usage()
	return text
}
