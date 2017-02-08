package yukimikubot

import (
	"github.com/bwmarrin/discordgo"
)

func AppendHandlers(s *discordgo.Session) {

	s.AddHandler(setStatus)

}

func setStatus(s *discordgo.Session, event *discordgo.Ready) {
	_ = s.UpdateStatus
}
