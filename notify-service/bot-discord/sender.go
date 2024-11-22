package botdiscord

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type DiscordSender struct {
	Target  string
	Message string
	Option  []discordgo.RequestOption
}

func (d DiscordSender) SendMessage() error {

	target := d.Target
	message := d.Message

	if strings.TrimSpace(target) == `` {
		return fmt.Errorf(`[Discord] target is empty`)
	}

	if strings.TrimSpace(message) == `` {
		return fmt.Errorf(`[Discord] message is empty`)
	}

	_, err := service.ChannelMessageSend(target, message, d.Option...)
	if err != nil {
		return err
	}

	fmt.Printf("Notify to discord channel [%s] success\n", target)

	return nil
}
