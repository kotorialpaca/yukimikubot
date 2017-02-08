package yukimikubot

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func AppendHandlers(s *discordgo.Session) {

	s.AddHandler(setStatus)

}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if strings.HasPrefix(m.Content, Cfg.Prefix) {

	}

}

func setStatus(s *discordgo.Session, event *discordgo.Ready) {

	_ = s.UpdateStatus(0, "Yuki no Hana")

}
