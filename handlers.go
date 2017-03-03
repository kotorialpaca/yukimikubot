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

	ch, _ := s.Channel(m.ChannelID)
	creator, _ := s.State.Member(ch.GuildID, m.Author.ID)

	if creator.Nick == "" {
		creator.Nick = creator.User.Username
	}
	/*
		fmt.Println("guild id is:", ch.GuildID)
		fmt.Println("author id is: ", m.Author.ID)
		fmt.Println("creator nick is:", creator.Nick)
		fmt.Println("creator id is: ", creator.User.ID)
	*/
	if strings.HasPrefix(m.Message.Content, Cfg.Prefix) {
		//Test Print Output Author -> Message
		fmt.Println(m.Author, " -> ", m.Message.Content)
		switch cmd := strings.Split(m.Message.Content, " "); strings.ToLower(cmd[0][1:]) {
		/*
			!event add NAME DEFAULTGROUP(OPTIONAL)
			!event conf EVT_NAME start START_TIME
			!event conf EVT_NAME end END_TIME
			!event conf EVT_NAME desc DESCRIPTIONS
			!event conf EVT_NAME max MAXNUM
			!event remove
			!event list
		*/
		case "event":
			switch strings.ToLower(cmd[1]) {
			case "add":
				maxnum := 0
				if len(cmd) == 3 {
					maxnum = 100
				} else if len(cmd) < 3 {
					s.ChannelMessageSend(m.ChannelID, strings.Replace("<@!user>, Error - Insufficient number of arguments given", "user", creator.User.ID, -1))
					return
				}
				newEvt, err := controllers.NewEvent(m.Message.Content[11:], *creator, maxnum, ch.GuildID)
				//Implement DB Interaction
				if err != nil {
					s.ChannelMessageSend(m.ChannelID, strings.Replace("<@!user>, Error - Event Creation has failed", "user", creator.User.ID, -1))
					return
				}

				s.ChannelMessageSend(m.ChannelID, strings.Replace("<@!user>, Event has been successfully created!", "user", creator.User.ID, -1))
				s.ChannelMessageSend(m.ChannelID, newEvt.PrintPrettyString())
				return
			case "conf":
				switch strings.ToLower(cmd[len(cmd)-3]) {
				case "start":
					name := m.Message.Content[12 : strings.Index(m.Message.Content, "start")-1]
					/*
						fmt.Println("---- start ----")
						for _, test := range []byte(name[:19]) {
							fmt.Print(string(test))
						}
						fmt.Println("---- end ----")
						fmt.Println(strconv.Itoa(len(name)))
					*/
					evt, err := controllers.FindEvent(ch.GuildID, name)
					if err != nil {
						fmt.Println(err)
						s.ChannelMessageSend(m.ChannelID, "FILE CANNOT B FOUND SON")
						return
					}
					err = evt.SetStartDate(cmd[len(cmd)-2] + " " + cmd[len(cmd)-1])
					if err != nil {
						s.ChannelMessageSend(m.ChannelID, wrongDateWarning(creator))
						return
					}
					s.ChannelMessageSend(m.ChannelID, "Date for "+name+" has been successfully added!")
					return
				case "end":

					name := m.Message.Content[12:strings.Index(m.Message.Content, "end")]
					evt, err := controllers.FindEvent(ch.GuildID, name)
					err = evt.SetEndDate(cmd[len(cmd)-2] + " " + cmd[len(cmd)-1])
					if err != nil {
						s.ChannelMessageSend(m.ChannelID, wrongDateWarning(creator))
						return
					}
					return
				/*
					case "desc":
						name := m.Message.Content[12:strings.Index(m.Message.Content, "desc")]
						evt, err := controllers.FindEvent(ch.GuildID, name)
						if err != nil {
							s.ChannelMessageSend(m.ChannelID, "Event cannot be found")
							return
						}
						evt.SetDesc(m.Content[strings.Index(m.Message.Content, "desc")+6:])
						s.ChannelMessageSend(m.ChannelID, "Event Configuration Set - Description has been successfully modified!")
						return
				*/
				default:
					//if the command line doesnt contain desc
					if strings.Index(m.Message.Content, " desc ") == -1 {
						//if the command doesnt contain desc or max
						if strings.Index(m.Message.Content, " max ") == -1 {
							s.ChannelMessageSend(m.ChannelID, "Wrong parameter for EVENT CONF command")
							return
						} else {
							//if the command contains max
							name := m.Message.Content[12:strings.Index(m.Message.Content, "max")]
							evt, err := controllers.FindEvent(ch.GuildID, name)
							if err != nil {
								s.ChannelMessageSend(m.ChannelID, "Specified event could not be found")
							}
							c, er := strconv.Atoi(cmd[len(cmd)-1])
							if er != nil {
								s.ChannelMessageSend(m.ChannelID, "Invalid Parameter for Max Number given")
								return
							}
							evt.SetMax(c)
							s.ChannelMessageSend(m.ChannelID, "Max number has been successfully modified")
							return
						}
					} else {
						//if the command line contains desc
						name := m.Message.Content[12 : strings.Index(m.Message.Content, "desc")-1]
						fmt.Println(name)
						evt, err := controllers.FindEvent(ch.GuildID, name)
						if err != nil {
							s.ChannelMessageSend(m.ChannelID, "Event cannot be found")
							return
						}
						evt.SetDesc(m.Content[strings.Index(m.Message.Content, "desc")+6:])
						s.ChannelMessageSend(m.ChannelID, "Event Configuration Set - Description has been successfully modified!")
						return
					}
					//s.ChannelMessageSendEmbed
				}
			//case "remove":

			case "list":
				if len(cmd) != 2 {
					s.ChannelMessageSend(m.ChannelID, "Bro, too many param for event list")
					return
				}
				evtGrp := controllers.NewEventGroup(ch.GuildID)
				err := evtGrp.RetrieveEvents()

				if err != nil {
					s.ChannelMessageSend(m.ChannelID, "something wrong happened while reading event data")
					return
				}

				s.ChannelMessageSend(m.ChannelID, evtGrp.PrintPrettyEventStrings())
				return

			//case "signup":

			//default for event param
			default:
				s.ChannelMessageSend(m.ChannelID, "Error - Invalid Command\nDISPLAY EVENT HELP HERE")
				return
			}

		//default for prefix+command - unrecognizable command
		default:
			s.ChannelMessageSend(m.ChannelID, strings.Replace("Cannot recognize command, try #prxhelp", "#prx", Cfg.Prefix, -1))
			return
		}

	}

}

func wrongParamWarning(m *discordgo.Member) string {
	return strings.Replace("@user, Error - Invalid number of arguments given", "user", m.Nick, -1)
}

func wrongDateWarning(m *discordgo.Member) string {
	return strings.Replace("@user, Error - Invalid timestamp given", "user", m.Nick, -1)
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
	fmt.Println(Cfg.Status)
	err := s.UpdateStatus(0, Cfg.Status)
	if err != nil {
		fmt.Println(err)
	}

}

func printHelp() string {
	return ""
}
