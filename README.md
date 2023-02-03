A serverless Telegram bot

# Baicai TGBot

A Telegram bot

## 部署到 Vercel

[Vercel Deploy](https://vercel.com/import/project?template=https://github.com/clin003/telebot_app_serverless)

##  环境变量

```bash
# 机器人token
BAICAI_BOT_TELEGRAM_TOKEN
# 自定义 注册webhook url 地址（每次构建的时候注册这个地址，可留空）
BAICAI_BOT_TELEGRAM_WEBHOOK_URL
# 视频接口地址
BAICAI_BOT_VIDEO_API
# Notify 推送验证token
BAICAI_NOTIFY_CHECKTOKEN
# 自动发送对象id
BAICAI_BOT_TELEGRAM_AUTOSEND_CHAT_ID
# 自动发送尾巴内容
BAICAI_BOT_TELEGRAM_AUTOSEND_CAPTION

# 兼容openAPI 推送验证 对应openAPI webhook的token
BAICAI_WSPUSH_CHANNEL_**=**
# 兼容openAPI 推送对象id
BAICAI_BOT_TELEGRAM_WSPUSH_FEED_CHAT_ID_**=
```


##  交流群

交流TG群： @baicai_dev

##  感谢

