package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

var token string
var buffer = make([][]byte, 0)

func main() {

	//Read file named token.txt in the same folder
	body, err := ioutil.ReadFile("token.txt")
	if err != nil {
		fmt.Println("error opening token file: ", err)
		if os.IsNotExist(err) {
			var f, e = os.Create("token.txt")
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

	dg, err := discordgo.New("Bot " + token)

	if err != nil {
		fmt.Println("Discord session creation failiure: ", err)
		return
	}

	err = dg.Open()

	if err != nil {
		log.Fatalln("Discord session could not be opened: ", err)
	}

}
