package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"

	"strings"

	"github.com/bwmarrin/discordgo"
)

var token string
var buffer = make([][]byte, 0)

//Cfg is the config of the bot defined on settings.conf
var Cfg Config

func main() {
	// ** DISCORD BOT TOKEN HANDLING START **
	body, err := ioutil.ReadFile("token.conf")
	if err != nil {
		fmt.Println("error opening token file: ", err)
		if os.IsNotExist(err) {
			var f, e = os.Create("token.conf")
			if e != nil {
				log.Fatalf("failed to create file")
			}
			defer f.Close()
			log.Fatalf("created token file, please input bot token in the token file")
		}
	}

	token = string(body)
	fmt.Println(token)

	if len(body) != 59 {
		log.Fatalf("invalid bot token, please check your bot token and try again.")
	}
	// ** TOKEN HANDLING FINISHED

	// ** SETTINGS FILE HANDLING START
	raw, err := ioutil.ReadFile("settings.conf")
	if err != nil {
		fmt.Println("error opening config file: ", err)
	}
	err = json.Unmarshal(raw, &Cfg)
	if err != nil {
		fmt.Println("error has occurred during unmarshal process, ", err)
	}

	// Video Caching Enabled
	if Cfg.CacheEnabled {
		fmt.Println("playback caching enabled")
		if _, err := os.Stat("cache"); err != nil {
			os.Mkdir("cache", 770)
			fmt.Println("cache directory created")
		} else {
			fmt.Println("cache directory exists")
		}
	} else {
		fmt.Println("playback caching disabled")
	}
	// ** HANDLE DISCORD SESSION START
	dg, err := discordgo.New("Bot " + token)

	if err != nil {
		fmt.Println("Discord session creation failiure: ", err)
		return
	}

	err = dg.Open()
	// Append Handlers
	AppendHandlers(dg)
	if err != nil {
		log.Fatalln("Discord session could not be opened: ", err)
	}

	fmt.Println(Cfg.Name + " is now online.")
	u, err := dg.User("@me")
	fmt.Println(u.ID)
	if err != nil {
		fmt.Println("error getting account details, ", err)
	}
	fmt.Println("To join this bot to the server, go to ")

	url := strings.Replace(Cfg.OauthLink, "CID", u.ID, -1)
	fmt.Println(url)

	//Establish gochannel with os.Signal, watch for ^C
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Printf("Closing Discord Bot %s\n", Cfg.Name)
		_ = dg.Close()
		os.Exit(0)
	}()

	<-make(chan struct{})
	return

}

//Config structure of the bot defined on settings.conf
type Config struct {
	Name         string `json:"name"`
	Prefix       string `json:"prefix"`
	CacheEnabled bool   `json:"cache_enabled"`
	Status       string `json:"status"`
	OauthLink    string `json:"oauth_link"`
}
