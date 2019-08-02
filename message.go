package dingtalk

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

const sendUrl = "https://oapi.dingtalk.com/robot/send?access_token="

/**
钉钉markdown支持的格式

标题
# 一级标题
## 二级标题
### 三级标题
#### 四级标题
##### 五级标题
###### 六级标题

引用
> A man who stands for nothing will fall for anything.

文字加粗、斜体
**bold**
*italic*

链接
[this is a link](http://name.com)

图片
![](http://name.com/pic.jpg)

无序列表
- item1
- item2

有序列表
1. item1
2. item2
*/

const EnvXArmoryIgnoreSendDingTalk = "X_ARMORY_IGNORE_SEND_DING_TALK"

type Util struct {
	AccessToken string
}

func (u *Util) SendText(t string) {
	if os.Getenv(EnvXArmoryIgnoreSendDingTalk) == "true" {
		return
	}
	msg := Message{
		AccessToken: u.AccessToken,
		Msgtype:     TEXT,
		Text: &TextMessage{
			Content: t,
		},
	}

	msg.Send()
}

func (msg *Message) Send() (*MessageResp, error) {
	marshal, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Post(sendUrl+msg.AccessToken, "application/json", bytes.NewReader(marshal))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	messageResp := MessageResp{}
	json.Unmarshal(body, &messageResp)
	return &messageResp, nil
}

type Message struct {
	AccessToken string             `json:"-"`
	Msgtype     MessageType        `json:"msgtype"`
	Text        *TextMessage       `json:"text"`
	Link        *LinkMessage       `json:"link"`
	Markdown    *MarkdownMessage   `json:"markdown"`
	ActionCard  *ActionCardMessage `json:"actionCard"`
	FeedCard    *FeedCardMessage   `json:"feedCard"`
	At          *MessageAt         `json:"at"`
}

type MessageType string

const (
	TEXT       MessageType = "text"
	LINK       MessageType = "link"
	MARKDOWN   MessageType = "markdown"
	ActionCard MessageType = "actionCard"
	FeedCard   MessageType = "feedCard"
)

type TextMessage struct {
	Content string `json:"content"`
}

type LinkMessage struct {
	Text       string `json:"text"`
	Title      string `json:"title"`
	PicUrl     string `json:"picUrl"`
	MessageUrl string `json:"messageUrl"`
}

type MarkdownMessage struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type ActionCardMessage struct {
	Title          string                   `json:"title"`
	Text           string                   `json:"text"`
	HideAvatar     string                   `json:"hideAvatar"`
	BtnOrientation string                   `json:"btnOrientation"`
	SingleTitle    string                   `json:"singleTitle"`
	SingleURL      string                   `json:"singleURL"`
	Btns           *[]*ActionCardMessageBtn `json:"btns"`
}

type FeedCardMessage struct {
	Links *[]*FeedCardLinkMessage `json:"links"`
}

type FeedCardLinkMessage struct {
	Title      string `json:"title"`
	MessageURL string `json:"messageURL"`
	PicURL     string `json:"picURL"`
}

type ActionCardMessageBtn struct {
	Title     string `json:"title"`
	ActionURL string `json:"actionURL"`
}

type MessageAt struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}

type MessageResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}
