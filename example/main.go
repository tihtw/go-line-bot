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

	profiles, _ := bot.GetUserProfile(req.Result[0].Content.From)
	if req.Result[0].EventType == lbot.EventOperation {

	} else {
		contentType := req.Result[0].Content.ContentType
		switch contentType {
		case lbot.TextMessage:
			log.Printf("User %v -> Bot: %v\n", profiles[0].DisplayName, req.Result[0].Content.Text)
			receivedMsg := req.Result[0].Content.Text
			bot.SendTextMessage(req.Result[0].Content.From, receivedMsg)
			log.Printf("Bot -> User %v: %v", profiles[0].DisplayName, receivedMsg)
		case lbot.ImageMessage:
			log.Printf("User %v -> Bot: Image id : %v\n", profiles[0].DisplayName, req.Result[0].Content.Id)
		case lbot.VideoMessage:
			log.Printf("User %v -> Bot: Video id : %v\n", profiles[0].DisplayName, req.Result[0].Content.Id)
		case lbot.AudioMessage:
			log.Printf("User %v -> Bot: Audio id : %v\n", profiles[0].DisplayName, req.Result[0].Content.Id)
		case lbot.LocationMessage:
			log.Printf("User %v -> Bot: Location %v %v (%v, %v)\n", profiles[0].DisplayName,
				req.Result[0].Content.Location.Title,
				req.Result[0].Content.Location.Address,
				req.Result[0].Content.Location.Latitude,
				req.Result[0].Content.Location.Longitude)
		case lbot.StickerMessage:
			log.Printf("User %v -> Bot: Stiker  %v\n", profiles[0].DisplayName,
				buildStickerUrl(
					lbot.Sticker{
						Stkver:   req.Result[0].Content.ContentMetadata.Stkver,
						Stkpkgid: req.Result[0].Content.ContentMetadata.Stkpkgid,
						Stkid:    req.Result[0].Content.ContentMetadata.Stkid}))

		default:
			log.Printf("Unsupport content type: %v from %v", contentType, req.Result[0].Content.From)

		}
	}

	// log.Printf("User %v: %v\n", req.Result[0].Content.From, req.Result[0].Content.Text)

	// log.Printf("%v: %v\n", profiles[0].DisplayName, req)

}

func buildStickerUrl(s lbot.Sticker) string {
	return "http://dl.stickershop.line.naver.jp/products/0/0/" + s.Stkver + "/" + s.Stkpkgid + "/PC/stickers/" + s.Stkid + ".png"
}
