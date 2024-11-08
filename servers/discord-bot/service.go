package discordbot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

const (
	GuildText          = discordgo.ChannelTypeGuildText
	DM                 = discordgo.ChannelTypeDM
	GuildVoice         = discordgo.ChannelTypeGuildVoice
	GroupDM            = discordgo.ChannelTypeGroupDM
	GuildCategory      = discordgo.ChannelTypeGuildCategory
	GuildNews          = discordgo.ChannelTypeGuildNews
	GuildStore         = discordgo.ChannelTypeGuildStore
	GuildNewsThread    = discordgo.ChannelTypeGuildNewsThread
	GuildPublicThread  = discordgo.ChannelTypeGuildPublicThread
	GuildPrivateThread = discordgo.ChannelTypeGuildPrivateThread
	GuildStageVoice    = discordgo.ChannelTypeGuildStageVoice
	GuildForum         = discordgo.ChannelTypeGuildForum
	GuildMedia         = discordgo.ChannelTypeGuildMedia
)

var TypeChannal = map[string]discordgo.ChannelType{
	`GuildText`:          GuildText,
	`DM`:                 DM,
	`GuildVoice`:         GuildVoice,
	`GroupDM`:            GroupDM,
	`GuildCategory`:      GuildCategory,
	`GuildNews`:          GuildNews,
	`GuildStore`:         GuildStore,
	`GuildNewsThread`:    GuildNewsThread,
	`GuildPublicThread`:  GuildPublicThread,
	`GuildPrivateThread`: GuildPrivateThread,
	`GuildStageVoice`:    GuildStageVoice,
	`GuildForum`:         GuildForum,
	`GuildMedia`:         GuildMedia,
}

type IDiscordService interface {
	GetAllChannels(t discordgo.ChannelType) ([]map[string]interface{}, error)
	ChannelMessageSend(channel, message string, options ...discordgo.RequestOption) (*discordgo.Message, error)
}

type discordService struct {
	s *discordgo.Session
}

var DiscordService IDiscordService

func newService(s *discordgo.Session) IDiscordService {
	DiscordService = &discordService{
		s: s,
	}
	return DiscordService
}

func (dc *discordService) GetAllChannels(t discordgo.ChannelType) ([]map[string]interface{}, error) {
	var chans []map[string]interface{}
	// Get all guilds (servers) the bot is a member of
	guilds := dc.s.State.Guilds

	// Iterate over each guild and print the channel IDs
	for _, guild := range guilds {
		gui, err := dc.s.Guild(guild.ID)
		if err != nil {
			fmt.Printf("Guild ID: %s\n , Name : %s\n", gui.ID, gui.Name)
			return nil, err
		}

		GChan, err := dc.s.GuildChannels(guild.ID)
		if err != nil {
			fmt.Printf("Guild Channels : %s", err.Error())
			return nil, err
		}

		for _, chanx := range GChan {
			if t == chanx.Type {
				chans = append(chans, map[string]interface{}{
					chanx.ID: chanx.Name,
				})
			}
		}
	}

	return chans, nil
}

func (dc *discordService) ChannelMessageSend(channel, message string, options ...discordgo.RequestOption) (*discordgo.Message, error) {
	ms, err := dc.s.ChannelMessageSend(channel, message, options...)
	if err != nil {
		return nil, err
	}

	return ms, nil
}
