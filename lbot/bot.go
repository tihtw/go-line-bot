package lbot

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func (b *Bot) SetConfig(c Config) {
	b.config = &c
}

func (b *Bot) SendTextMessage(mid, s string) error {
	if b.config == nil {
		return errors.New("Config have not been set")
	}
	if b.config.Debug {
		log.Println("Start to Set Message")
	}

	var payload Request
	payload.SetDefaults()
	payload.SetText(s)
	payload.AddTargetUser(mid)

	req, err := http.NewRequest("POST", b.config.ServerHost+"/v1/events", strings.NewReader(string(s)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-type", "application/json;charset=UTF-8]")
	req.Header.Set("X-Line-ChannelID", b.config.ChannelID)
	req.Header.Set("X-Line-ChannelSecret", b.config.ChannelSecret)
	req.Header.Set("X-Line-Trusted-User-With-ACL", b.config.MID)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if b.config.Debug {
		log.Println("Result: ", string(result))
	}

	return nil
}
