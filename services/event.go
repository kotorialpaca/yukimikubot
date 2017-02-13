package services

import (
	"strings"
	"time"

	"fmt"

	"github.com/bwmarrin/discordgo"
)

type EventGroup struct {
	Name   string
	Events []Event
}

type Event struct {
	Name      string
	Author    discordgo.Member
	Desc      string
	Members   []discordgo.Member
	MaxMember int
	Groups    []Group
	StartDate time.Time
	EndDate   time.Time
}

type Group struct {
	Name      string
	Members   []discordgo.Member
	MaxMember int
}

func NewEvent(n string, d string, sd string, ed string, u discordgo.Member, m int) *Event {
	layout := "2006-01-02 03:04PM"
	//Start Time in Time variable
	stt, err := time.Parse(layout, sd)
	if err != nil {
		fmt.Println("welp something went wrong while creating an event")
	}
	ett, err := time.Parse(layout, ed)
	if err != nil {
		fmt.Println("welp something went wrong while creating an event")
	}
	return &Event{
		Name:      n,
		Desc:      d,
		Author:    u,
		MaxMember: m,
		StartDate: stt,
		EndDate:   ett,
	}
}

func NewGroup(n string, max int) *Group {
	return &Group{
		Name:      n,
		MaxMember: max,
	}
}

/*
":star: **EVT_NAME** :star:\n
__Started by: AUTHOR__\n\n

Details: DESC\n\n

Members: CUR_MEM\n
Max Members: MAX_MEM\n\n

**[Group1]**\n
*bleh* \n
*dumbo* \n
*bob* \n\n

**[Group2]**
*john* \n
*doe* \n
*paul* \n

**[Group3]**
*lima*
*josh*"
*/
func (e Event) PrintPrettyString() string {
	outstr := ":star: **EVT_NAME** :star:\n__Started by: AUTHOR__\n\nDetails: DESC\n\nMembers: CUR_MEM\nMax Members: MAX_MEM\n\n"
	outstr = strings.Replace(outstr, "EVT_NAME", e.Name, -1)
	outstr = strings.Replace(outstr, "AUTHOR", e.Author.Nick, -1)
	outstr = strings.Replace(outstr, "DESC", e.Desc, -1)
	outstr = strings.Replace(outstr, "CUR_MEM", string(len(e.Members)), -1)
	outstr = strings.Replace(outstr, "MAX_MEM", string(e.MaxMember), -1)

	return outstr
}

func (e Event) GroupsToString() string {
	outstr := ""
	if len(e.Groups) != 0 {
		for a, b := range e.Groups {
			outstr = outstr + "**[" + b.Name + "]**\n"
			for i, v := range b.Members {
				outstr = outstr + "*" + v.Nick + "*\n"
			}
		}
	} else {
		outstr = "**[Members]**\n"
		for a, b := range e.Members {
			outstr = outstr + "**[" + b.Nick + "]**\n"
		}
	}

	return outstr
}
