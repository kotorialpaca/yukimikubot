package controllers

import (
	"encoding/json"
	"strings"
	"time"

	"fmt"

	"strconv"

	"errors"

	"github.com/kotorialpaca/yukimikubot/helper"

	"github.com/boltdb/bolt"
	"github.com/bwmarrin/discordgo"
)

//EventGroup object which groups all the Event objects
type EventGroup struct {
	GuildID string  `json:"guild_id"`
	Events  []Event `json:"events"`
}

//Event object to represent events for the Discord chat
type Event struct {
	ID        uint64           `json:"id"`
	Name      string           `json:"name"`
	Author    discordgo.Member `json:"author"`
	Desc      string           `json:"desc"`
	MaxMember int              `json:"max_member"`
	Groups    []Group          `json:"groups"`
	StartDate time.Time        `json:"start_date"`
	EndDate   time.Time        `json:"end_date"`
	GuildID   string           `json:"guild_id"`
}

//Group object for storing groups split within an Event object
type Group struct {
	Name      string             `json:"name"`
	Members   []discordgo.Member `json:"members"`
	MaxMember int                `json:"max_members"`
}

//NewEvent returns a new Event object
func NewEvent(n string, sd string, ed string, u discordgo.Member, m int, def bool, gID string) (*Event, error) {
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
	newEvent := &Event{
		Name:      n,
		Author:    u,
		MaxMember: m,
		StartDate: stt,
		EndDate:   ett,
		GuildID:   gID,
	}
	if def {
		newEvent.AddGroupToEvent("DefaultGroup", 100, u)
	}

	if err != nil {
		return &Event{}, nil
	}

	err = newEvent.UpdateEvent()

	if err != nil {
		return nil, err
	}

	return newEvent, nil
}

func (e *Event) UpdateEvent() error {

	ndb, err := bolt.Open(e.GuildID, 0600, nil)
	if err != nil {
		return err
	}

	ndb.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("event"))

		if err != nil {
			return err
		}
		if e.ID == 0 {
			id, _ := b.NextSequence()
			e.ID = id
		}

		buf, err := json.Marshal(e)

		if err != nil {
			return err
		}

		return b.Put(helper.Itob(e.ID), buf)
	})

	return nil
}

func (eg *EventGroup) RetrieveEvents() error {
	ndb, err := bolt.Open(eg.GuildID, 0600, nil)
	if err != nil {
		return err
	}
	ndb.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("event"))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			n := len(eg.Events)
			if n == cap(eg.Events) {
				newSlice := make([]Event, len(eg.Events), len(eg.Events)+1)
				copy(newSlice, eg.Events)
				eg.Events = newSlice
			}
			eg.Events = eg.Events[0 : n+1]

			evt := &Event{}

			err := json.Unmarshal(v, &evt)

			if err != nil {
				return err
			}

			eg.Events[n] = *evt
		}

		return nil
	})

	return nil

}

//NewGroup returns a new Group object within Event
func NewGroup(n string, max int) *Group {
	return &Group{
		Name:      n,
		MaxMember: max,
	}
}

/*
PrintPrettyString will print prettified string of event detaills.
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
func (e *Event) PrintPrettyString() string {
	outstr := ":star: **EVT_NAME** :star:\n__Started by: AUTHOR__\n\nDetails: DESC\n\nMax Members: MAX_MEM\n\n"
	outstr = strings.Replace(outstr, "EVT_NAME", e.Name, -1)
	outstr = strings.Replace(outstr, "AUTHOR", e.Author.Nick, -1)
	outstr = strings.Replace(outstr, "DESC", e.Desc, -1)
	//outstr = strings.Replace(outstr, "CUR_MEM", strconv.Itoa(len(e.Members)), -1)
	outstr = strings.Replace(outstr, "MAX_MEM", strconv.Itoa(e.MaxMember), -1)

	outstr = outstr + e.GroupsToString()
	return outstr
}

//GroupsToString returns string representation of the groups
func (e *Event) GroupsToString() string {
	outstr := ""
	if len(e.Groups) != 0 {
		for _, b := range e.Groups {
			outstr = outstr + "**[" + b.Name + "]**\n"
			if len(b.Members) == 0 {
				outstr = outstr + "-- Empty Group --\n\n"
			}
			for _, v := range b.Members {
				outstr = outstr + "*" + v.Nick + "*\n"
			}
		}
	} else {
		outstr = "**[Members]**\n"
		for _, b := range e.Groups[0].Members {

			outstr = outstr + b.Nick + "\n"
		}
	}

	return outstr
}

func (e *Event) AddGroupToEvent(gn string, max int, author discordgo.Member) {
	newGroup := Group{
		Name:      gn,
		MaxMember: max,
	}
	n := len(e.Groups)
	if n == cap(e.Groups) {
		newG := make([]Group, len(e.Groups), len(e.Groups)+1)
		copy(newG, e.Groups)
		e.Groups = newG
	}
	e.Groups = e.Groups[0 : n+1]
	e.Groups[n] = newGroup

}

//AddMemberToGroup will add the member discordgo.Member to the group as an object
func (e *Event) AddMemberToGroup(gn string, m discordgo.Member) error {

	for key, value := range e.Groups {
		if strings.Compare(value.Name, gn) == 0 {
			if len(value.Members) == value.MaxMember {
				return errors.New("member full")
			}
			n := len(value.Members)
			if n == cap(value.Members) {
				newM := make([]discordgo.Member, len(value.Members), len(value.Members)+1)
				copy(newM, value.Members)
				value.Members = newM
			}
			value.Members = value.Members[0 : n+1]
			value.Members[n] = m
			e.Groups[key] = value

		}

	}

	return nil

}

func (e *Event) RemoveMemberFromGroup(gn string, m discordgo.Member) error {
	for key, value := range e.Groups {
		//If group name and the current iteration of the group matches then
		if strings.Compare(gn, value.Name) == 0 {
			for k, v := range value.Members {
				if strings.Compare(m.Nick, v.Nick) == 0 {
					value.Members = RemoveMemberFromGroup(value.Members, k)
					e.Groups[key] = value
					return nil
				}
			}
		}
	}
	return errors.New("specified group name or member couldnt not be found")

}

func (e *Event) RemoveGroup(gn string) error {

	for key, value := range e.Groups {
		if strings.Compare(value.Name, gn) == 0 {
			e.Groups = RemoveGroup(e.Groups, key)
			return nil
		}
	}

	return errors.New("specified group name could not be found")
}

//func GenerateNewMemberList(li []discordgo.Member)

func (e *Event) GetGroup(gn string) (Group, error) {
	for _, value := range e.Groups {
		if value.Name == gn {
			return value, nil
		}
	}

	return Group{}, errors.New("cannot find, many keks")
}

func RemoveMemberFromGroup(s []discordgo.Member, i int) []discordgo.Member {
	return append(s[:i], s[i+1:]...)
}

func RemoveGroup(s []Group, i int) []Group {
	return append(s[:i], s[i+1:]...)
}

/*
func (g *Group) AddMemberToGroup(m discordgo.Member) {

	n := len(g.Members)
	if n == cap(g.Members) {
		newM := make([]discordgo.Member, len(g.Members), len(g.Members)+1)
		copy(newM, g.Members)
		g.Members = newM
	}
	g.Members = g.Members[0 : n+1]
	g.Members[n] = m
}
*/
/*
func addToGroup(slice interface{}, element interface{}) interface{} {
	n := len(slice)
	if n == cap(slice) {
		// Slice is full; must grow.
		// We double its size and add 1, so if the size is zero we still grow.
		newSlice := make(interface{}, len(slice), len(slice)+1)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}
*/
