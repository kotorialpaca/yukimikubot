package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func AppendHandlers(s *discordgo.Session) {

	s.AddHandler(setStatus)
	s.AddHandler(onMessageCreate)
	//s.AddHandler(onChannelCreate)
	s.AddHandler(onGuildCreate)

}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println(m.Message.Content)
	if strings.HasPrefix(m.Message.Content, Cfg.Prefix) {
		fmt.Println(m.Author, " -> ", m.Message.Content)
	}
	

}

func onGuildCreate(s *discordgo.Session, event *discordgo.GuildCreate) {
	if event.Guild.Unavailable {
		return
	}

	for _, channel := range event.Guild.Channels {
		if channel.ID == event.Guild.ID {
			s.ChannelMessageSend(channel.ID, Cfg.Name+" is now online.")
			return
		}
	}
}

func setStatus(s *discordgo.Session, event *discordgo.Ready) {

	_ = s.UpdateStatus(0, "Yuki no Hana")

}
