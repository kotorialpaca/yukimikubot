package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

//AppendHandlers function - Adds all the handlers previously declared somewhere in the package
func AppendHandlers(s *discordgo.Session) {

	s.AddHandler(setStatus)
	s.AddHandler(onMessageCreate)
	//s.AddHandler(onChannelCreate)
	s.AddHandler(onGuildCreate)

}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println(m.Message.Content)
	if strings.HasPrefix(m.Message.Content, Cfg.Prefix) {
		//Test Print Output Author -> Message
		fmt.Println(m.Author, " -> ", m.Message.Content)
		switch cmd := strings.Split(m.Message.Content, " "); strings.ToLower(cmd[0][1:]) {
		/*
			!event add
			!event modify
			!event remove
			!event list
		*/
		case "event":
			switch strings.ToLower(cmd[1]) {
			case "add":

			case "modify":

			case "remove":

			case "list":

			default:
				s.ChannelMessageSend(m.ChannelID, "Error - Invalid Command\nDISPLAY EVENT HELP HERE")
				return
			}
		}

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
