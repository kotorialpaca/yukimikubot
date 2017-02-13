package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/kotorialpaca/yukimikubot/services"
)

func main() {
	m := discordgo.Member{
		GuildID:  "1289307128937",
		JoinedAt: "dkm kek",
		Nick:     "top kek user",
		Deaf:     false,
		Mute:     false,
	}
	evt := services.NewEvent("Test Event", "This is a test event!", "2017-02-01 04:00PM", "2017-02-01 05:00PM", m, 10)
	fmt.Println(evt.PrintPrettyString())
}
