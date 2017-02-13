package services

import (
	"strings"
	"time"

	"fmt"

	"strconv"

	"errors"

	"github.com/bwmarrin/discordgo"
)

//EventGroup object which groups all the Event objects
type EventGroup struct {
	Name   string
	Events []Event
}

//Event object to represent events for the Discord chat
type Event struct {
	Name      string
	Author    discordgo.Member
	Desc      string
	MaxMember int
	Groups    []Group
	StartDate time.Time
	EndDate   time.Time
}

//Group object for storing groups split within an Event object
type Group struct {
	Name      string
	Members   []discordgo.Member
	MaxMember int
}

//NewEvent returns a new Event object
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

//NewGroup returns a new Group object within Event
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
	outstr = strings.Replace(outstr, "CUR_MEM", strconv.Itoa(len(e.Members)), -1)
	outstr = strings.Replace(outstr, "MAX_MEM", strconv.Itoa(e.MaxMember), -1)

	return outstr
}

//GroupsToString returns string representation of the groups
func (e Event) GroupsToString() string {
	outstr := ""
	if len(e.Groups) != 0 {
		for _, b := range e.Groups {
			outstr = outstr + "**[" + b.Name + "]**\n"
			for _, v := range b.Members {
				outstr = outstr + "*" + v.Nick + "*\n"
			}
		}
	} else {
		outstr = "**[Members]**\n"
		for _, b := range e.Members {
			outstr = outstr + "**[" + b.Nick + "]**\n"
		}
	}

	return outstr
}

func (e Event) addGroupToEvent(gn string, max int, author discordgo.Member)  {
	newGroup := Group{
		Name: gn,
		MaxMember: max,
	}
	e.Groups = addToSlice(e.Groups, newGroup)
}

func (g Group) addMemberToGroup(m discordgo.Member) {
	g.Members = addToSlice(g.Members, m)
}

func addToSlice(slice []interface{}}, element interface{}) []interface{}}{
	n := len(slice)
	if n == cap(slice) {
		// Slice is full; must grow.
		// We double its size and add 1, so if the size is zero we still grow.
		newSlice := make([]{}interface}, len(slice), len(slice)+1)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}
