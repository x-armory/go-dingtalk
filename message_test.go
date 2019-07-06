package dingtalk

import (
	"fmt"
	"testing"
)

func TestTextMessage(t *testing.T) {
	msg := Message{
		AccessToken: "cec09236b23441bffbdb022fbb479431fc0deb676156590f58ca66ee4cb8b93d",
		Msgtype:     TEXT,
		Text: &TextMessage{
			Content: "Text Message Content",
		},
	}

	msg.Send()
}

func TestLinkMessage(t *testing.T) {
	msg := Message{
		AccessToken: "cec09236b23441bffbdb022fbb479431fc0deb676156590f58ca66ee4cb8b93d",
		Msgtype:     LINK,
		Link: &LinkMessage{
			Title:      "title",
			Text:       "This is a link",
			PicUrl:     "https://gw.alipayobjects.com/zos/skylark-tools/public/files/84111bbeba74743d2771ed4f062d1f25.png",
			MessageUrl: "http://www.baidu.com",
		},
	}

	msg.Send()
}

func TestMarkdownMessage(t *testing.T) {
	msg := Message{
		AccessToken: "cec09236b23441bffbdb022fbb479431fc0deb676156590f58ca66ee4cb8b93d",
		Msgtype:     MARKDOWN,
		Markdown: &MarkdownMessage{
			Title: "title",
			Text: fmt.Sprintf("#### **%s %s %s**\n"+
				"![screenshot](%s)\n"+
				"###### 11点发布 [更多信息](http://www.thinkpage.cn/) \n",
				"title1y",
				"title2w",
				"titleON",
				"https://gw.alipayobjects.com/zos/skylark-tools/public/files/84111bbeba74743d2771ed4f062d1f25.png"),
		},
	}

	msg.Send()
}
func TestMultiActionCardMessage(t *testing.T) {
	msg := Message{
		AccessToken: "cec09236b23441bffbdb022fbb479431fc0deb676156590f58ca66ee4cb8b93d",
		Msgtype:     ActionCard,
		ActionCard: &ActionCardMessage{
			Title: "乔布斯 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身",
			Text: "![screenshot](https://gw.alipayobjects.com/zos/skylark-tools/public/files/84111bbeba74743d2771ed4f062d1f25.png)\n" +
				"### 乔布斯 20 年前想打造的苹果咖啡厅\n" +
				"Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划",
			HideAvatar:     "0",
			BtnOrientation: "0",
			Btns: &[]*ActionCardMessageBtn{
				{
					Title:     "内容不错",
					ActionURL: "https://www.com/",
				},
				{
					Title:     "内容不错",
					ActionURL: "https://www.com/",
				},
				{
					Title:     "内容不错",
					ActionURL: "https://www.com/",
				},
			},
		},
	}

	msg.Send()
}

func TestSingleActionCardMessage(t *testing.T) {
	msg := Message{
		AccessToken: "cec09236b23441bffbdb022fbb479431fc0deb676156590f58ca66ee4cb8b93d",
		Msgtype:     ActionCard,
		ActionCard: &ActionCardMessage{
			Title: "乔布斯 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身",
			Text: "![screenshot](https://gw.alipayobjects.com/zos/skylark-tools/public/files/84111bbeba74743d2771ed4f062d1f25.png)\n" +
				"### 乔布斯 20 年前想打造的苹果咖啡厅\n" +
				"Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划",
			HideAvatar:     "0",
			BtnOrientation: "0",
			SingleTitle:    "阅读全文",
			SingleURL:      "https://www.com/",
		},
	}

	msg.Send()
}


func TestFeedCardMessage(t *testing.T) {
	msg := Message{
		AccessToken: "cec09236b23441bffbdb022fbb479431fc0deb676156590f58ca66ee4cb8b93d",
		Msgtype:     FeedCard,
		FeedCard: &FeedCardMessage{
			Links: &[]*FeedCardLinkMessage{
				{
					Title:      "时代的火车向前开",
					MessageURL: "https://www.com",
					PicURL:     "https://gw.alipayobjects.com/zos/skylark-tools/public/files/84111bbeba74743d2771ed4f062d1f25.png",
				},
				{
					Title:      "时代的火车向前开2",
					MessageURL: "https://www.baidu.com",
					PicURL:     "https://gw.alipayobjects.com/zos/skylark-tools/public/files/84111bbeba74743d2771ed4f062d1f25.png",
				},
			},
		},
	}

	msg.Send()
}