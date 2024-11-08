package linebot

import (
	"fmt"
	"os"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

func (m botline) Replace(data string, rep map[string]interface{}) string {
	for name, value := range rep {
		val := fmt.Sprintf(`%v`, value)
		if strings.Contains(name, `avatar`) {
			val = strings.TrimSpace(val)
			if val == `` {
				value = `https://vannila-asset.s3.ap-southeast-1.amazonaws.com/images/logo/logo-vannila-by-finnx-primary.png`
			}
		}
		data = strings.Replace(data, "*|"+name+"|*", val, -1)
	}
	return data
}

func (b *botline) Message(target string, msg string, alt string) error {

	content := []byte(msg)
	var err error

	container, err := linebot.UnmarshalFlexMessageJSON(content)
	if err != nil {
		return err
	}

	// local debug
	if os.Getenv("APP_HOST") == "127.0.0.1" {
		target = os.Getenv("LINE_DEBUG_USER")
	}

	message := linebot.NewFlexMessage(alt, container)

	_, err = b.client.PushMessage(target, message).Do()
	if err != nil {
		return err
	}

	return nil
}
