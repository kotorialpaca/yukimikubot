package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func AppendHandlers(s *discordgo.Session) {

	s.AddHandler(setStatus)
	s.AddHandler(onMessageCreate)
	s.AddHandler(onChannelCreate)
	s.AddHandler(onGuildCreate)

}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	fmt.Println("WTF")

}

func onChannelCreate(s *discordgo.Session, event *discordgo.GuildCreate) {
	if event.Guild.Unavailable {
		return
	}

	for _, channel := range event.Guild.Channels {
		if channel.ID == event.Guild.ID {
			_, _ = s.ChannelMessageSend(channel.ID, "Test is ready.")
			return
		}
	}
}

func onGuildCreate(s *discordgo.Session, event *discordgo.GuildCreate) {
	if event.Guild.Unavailable {
		return
	}

	for _, channel := range event.Guild.Channels {
		if channel.ID == event.Guild.ID {
			s.ChannelMessageSend(channel.ID, "**AIRHORN BOT READY FOR HORNING. TYPE `!AIRHORN` WHILE IN A VOICE CHANNEL TO ACTIVATE**")
			return
		}
	}
}

func setStatus(s *discordgo.Session, event *discordgo.Ready) {

	_ = s.UpdateStatus(0, "Yuki no Hana")

}
