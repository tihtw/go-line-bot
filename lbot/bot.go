package lbot

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func (b *Bot) SetConfig(c Config) {
	b.config = &c
}

func (b *Bot) SendTextMessage(m mid, s string) error {
	if b.config == nil {
		return errors.New("Config have not been set")
	}
	if b.config.Debug {
		log.Println("Start to Set Message")
	}

	var payload Request
	payload.SetDefaults()
	payload.SetText(s)
	payload.AddTargetUser(mid(m))

	out, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	if b.config.Debug {
		log.Println("Output json: " + string(out))
	}

	req, err := http.NewRequest("POST", b.config.ServerHost+"/v1/events", strings.NewReader(string(out)))
	if err != nil {
		return err
	}

	b.addAuthHeader(req)
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

func (b *Bot) GetUserProfile(m mid) ([]ProfileInfo, error) {
	if b.config == nil {
		return nil, errors.New("Config have not been set")
	}

	if b.config.Debug {
		log.Println("Start to Set Message")
	}

	req, err := http.NewRequest("GET", b.config.ServerHost+"/v1/profiles?mids="+string(m), nil)
	if err != nil {
		return nil, err
	}

	b.addAuthHeader(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if b.config.Debug {
		log.Println("Result: ", string(result))
	}

	return nil, nil

}

func (b *Bot) addAuthHeader(r *http.Request) {

	r.Header.Set("Content-type", "application/json; charset=UTF-8")
	r.Header.Set("X-Line-ChannelID", b.config.ChannelID)
	r.Header.Set("X-Line-ChannelSecret", b.config.ChannelSecret)
	r.Header.Set("X-Line-Trusted-User-With-ACL", b.config.MID)

}
