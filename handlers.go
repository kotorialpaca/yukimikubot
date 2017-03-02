package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/kotorialpaca/yukimikubot/controllers"
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
	ch, _ := s.Channel(m.ChannelID)
	creator, _ := s.GuildMember(ch.GuildID, m.Author.ID)
	if strings.HasPrefix(m.Message.Content, Cfg.Prefix) {
		//Test Print Output Author -> Message
		fmt.Println(m.Author, " -> ", m.Message.Content)
		switch cmd := strings.Split(m.Message.Content, " "); strings.ToLower(cmd[0][1:]) {
		/*
			!event add NAME START_DATE END_DATE MAX_MEM(OPTIONAL) DEFAULTGROUP(OPTIONAL)
			!event modify
			!event remove
			!event list
		*/
		case "event":
			switch strings.ToLower(cmd[1]) {
			case "add":
				maxnum := 0
				defa := false
				if len(cmd) == 5 {
					maxnum = 100
					defa = true
				} else if len(cmd) < 5 {
					s.ChannelMessageSend(m.ChannelID, strings.Replace("@user, Error - Insufficient number of arguments given", "user", creator.Nick, -1))
					return
				} else {
					max, err := strconv.Atoi(cmd[5])
					maxnum = max
					if err != nil {
						s.ChannelMessageSend(m.ChannelID, strings.Replace("@user, Error - Invalid Argument for Max Number", "user", creator.Nick, -1))
						return
					}
					d, err := strconv.ParseBool(cmd[6])
					defa = d
					if err != nil {
						s.ChannelMessageSend(m.ChannelID, strings.Replace("@user, Error - Invalid Argument for Default Group", "user", creator.Nick, -1))
						return
					}
				}
				newEvt, err := controllers.NewEvent(cmd[2], cmd[3], cmd[4], *creator, maxnum, defa, ch.GuildID)
				//Implement DB Interaction
				if err != nil {
					s.ChannelMessageSend(m.ChannelID, strings.Replace("@user, Error - Event Creation has failed", "user", creator.Nick, -1))
					return
				}

				s.ChannelMessageSend(m.ChannelID, strings.Replace("@user, Event has been successfully created!", "user", creator.Nick, -1))
				s.ChannelMessageSend(m.ChannelID, newEvt.PrintPrettyString())
				return
			case "modify":
				switch strings.ToLower(cmd[2]) {

				}
			case "remove":

			case "list":

			case "signup":

			//default for event param
			default:
				s.ChannelMessageSend(m.ChannelID, "Error - Invalid Command\nDISPLAY EVENT HELP HERE")
				return
			}
		default:
			s.ChannelMessageSend(m.ChannelID, strings.Replace("Cannot recognize command, try #prxhelp", "#prx", Cfg.Prefix, -1))
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
			s.ChannelMessageSend(channel.ID, Cfg.Name+" is now online.")
			return
		}
	}
}

func setStatus(s *discordgo.Session, event *discordgo.Ready) {

	_ = s.UpdateStatus(0, "Yuki no Hana")

}

func printHelp() string {
	return ""
}
