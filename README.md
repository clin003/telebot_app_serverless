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
# 自动发送对象id
BAICAI_BOT_TELEGRAM_AUTOSEND_CHAT_ID
# 自动发送尾巴内容
BAICAI_BOT_TELEGRAM_AUTOSEND_CAPTION
#video命令发送尾巴内容(私聊无尾巴)
BAICAI_BOT_VIDEO_CAPTION

# 兼容openAPI 推送验证 对应openAPI webhook的token
BAICAI_WSPUSH_CHANNEL_**=**
# 兼容openAPI 推送对象id
BAICAI_BOT_TELEGRAM_WSPUSH_FEED_CHAT_ID_**=
#订阅关键词列表(多个关键词，用竖线(|)分开)
BAICAI_WSPUSH_FEED_KEYWORLD_LIST_**=
#屏蔽关键词列表(多个关键词，用竖线(|)分开)
BAICAI_WSPUSH_FEED_KEYWORLD_FILTER_**=
#替换关键词列表(多个关键词，用竖线(|)分开)
BAICAI_WSPUSH_FEED_KEYWORLD_REPLACE_**=

#自动回复文本内容
BAICAI_BOT_REPLY_TEXT
#自动回复第一按钮描述
BAICAI_BOT_REPLY_BTN1_TEXT
#自动回复第一按钮URL
BAICAI_BOT_REPLY_BTN1_URL
#自动回复第二按钮描述
BAICAI_BOT_REPLY_BTN2_TEXT
#自动回复第二按钮URL
BAICAI_BOT_REPLY_BTN2_URL
```


##  交流群

交流TG群： @baicai_dev


## 赞赏
![赞赏白菜林，多少不重要，1元也是支持](https://cdn.jsdelivr.net/gh/clin003/cdn/assets/images/zanalipay.jpg)

![赞赏白菜林，多少不重要，1元也是支持](https://cdn.jsdelivr.net/gh/clin003/cdn/assets/images/zanweixin.jpg)