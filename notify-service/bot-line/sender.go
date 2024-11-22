package botline

import (
	"fmt"
	"strings"
)

type LineSender struct {
	Target   string
	Message  interface{} // string for text, map[string]interface for [flex,bubble] message
	Alt      *string     // for [flex,bubble] message
	Template *string     // for [flex,bubble] message
}

func (l LineSender) SendMessage() error {

	target := l.Target
	message := l.Message

	if strings.TrimSpace(target) == `` {
		return fmt.Errorf(`[Line] target is empty`)
	}

	switch v := message.(type) {
	case string:
		if err := textMessage(target, v); err != nil {
			return err
		}
	case map[string]interface{}:
		if l.Alt == nil || strings.TrimSpace(*l.Alt) == `` {
			return fmt.Errorf(`[Line] alt is empty`)
		}
		if l.Template == nil || strings.TrimSpace(*l.Template) == `` {
			return fmt.Errorf(`[Line] template is empty`)
		}
		if err := bubbleMessage(target, v, *l.Template, *l.Alt); err != nil {
			return err
		}
	default:
		return fmt.Errorf("type should be map[string]interface or string")
	}

	return nil
}

func textMessage(target, message string) error {

	return lineBot.TextMessage(target, message)
}

func bubbleMessage(target string, message map[string]interface{}, template, alt string) error {

	msg := replace(template, message)

	return lineBot.BubbleMessage(target, msg, alt)
}
