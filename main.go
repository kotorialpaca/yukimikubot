package yukimikubot

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"encoding/json"

	"github.com/bwmarrin/discordgo"
)

var token string
var buffer = make([][]byte, 0)
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

	if len(body) != 59 {
		log.Fatalf("invalid bot token, please check your bot token and try again.")
	}
	// ** TOKEN HANDLING FINISHED

	// ** SETTINGS FILE HANDLING START
	raw, err := ioutil.ReadFile("settings.conf")
	if err != nil {
		fmt.Println("error opening config file: ", err)
	}
	err = json.Unmarshal(raw, Cfg)

	if err != nil {
		Cfg.Prefix = "!"
	}
	// ** HANDLE DISCORD SESSION START
	dg, err := discordgo.New("Bot " + token)

	if err != nil {
		fmt.Println("Discord session creation failiure: ", err)
		return
	}

	err = dg.Open()

	if err != nil {
		log.Fatalln("Discord session could not be opened: ", err)
	}

	fmt.Println("Yukimiku bot is now online.")

	<-make(chan struct{})
	return

}

type Config struct {
	Prefix string `json:"prefix"`
}
