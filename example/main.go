package main

import (
	"fmt"
	"github.com/tihtw/go-line-bot/lbot"
	"log"
	"net/http"
	"os"
)

var bot lbot.Bot

func main() {
	fmt.Println("Server Start!!")
	var config lbot.Config
	config.SetDefaults()
	config.ChannelID = os.Getenv("LINE_CHANNEL_ID")
	config.ChannelSecret = os.Getenv("LINE_CHANNEL_SECRET")
	config.MID = os.Getenv("LINE_MID")
	config.Debug = true

	bot.SetConfig(config)
	http.HandleFunc("/", index)
	fmt.Println("Hello")
	fmt.Println(lbot.EventMessage)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("Received Message")
	req, err := lbot.ParseRequest(r)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("User %v: %v\n", req.Result[0].Content.From, req.Result[0].Content.Text)
	receivedMsg := req.Result[0].Content.Text

	log.Printf("%v\n", req)
	bot.SendTextMessage(req.Result[0].Content.From, receivedMsg)

}
