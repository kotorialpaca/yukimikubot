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
	m1 := discordgo.Member{
		GuildID:  "1289307128937",
		JoinedAt: "dkm kek",
		Nick:     "user1",
		Deaf:     false,
		Mute:     false,
	}
	m2 := discordgo.Member{
		GuildID:  "1289307128937",
		JoinedAt: "dkm kek",
		Nick:     "user2",
		Deaf:     false,
		Mute:     false,
	}
	evt := services.NewEvent("Test Event", "This is a test event!", "2017-02-01 04:00PM", "2017-02-01 05:00PM", m, 10)
	evt.AddGroupToEvent("DEEPS", 4, m)
	err := evt.AddMemberToGroup("DEEPS", m1)
	if err != nil {
		fmt.Println("idk y it no wurk, ", err)
	}
	err = evt.AddMemberToGroup("DEEPS", m2)
	if err != nil {
		fmt.Println("idk y it no wurk, ", err)
	}
	for _, value := range evt.Groups[0].Members {
		fmt.Println(value.Nick)
	}

	fmt.Println(evt.PrintPrettyString())
}
