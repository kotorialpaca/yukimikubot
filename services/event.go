package services

import (
    "github.com/bwmarrin/discordgo"
	"time"
)

type EventGroup struct{
    Name string,
    Events []Event,
}

type Event struct{
    Name string,
    Desc string,
    Member []discordgo.User,
    Groups map[string]discordgo.User,
    StartDate time.Date,
    EndDate time.Date,
}

func New(n string, d string, sd string, ed string) Event, error {
    layout := "2006-01-02 03:04PM"
    //Start Time in Time variable
    stt, err := time.Parse(layout, sd)
    if err != nil {
        return err
    }
    ett, err := time.Parse(layout, ed)
    if err != nil {
        return err
    }
    return &Event{
        Name:n,
        Desc:d,
        StartDate:stt,
        EndDate:ett
    }
}

func (e Event) PrintPretty() string {

}