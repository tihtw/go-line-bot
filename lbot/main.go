package lbot

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	// "errors"
	"io/ioutil"
	"net/http"
	// "strings"
	"time"
)

// The LINE BOT
type Bot struct {
	// Configuration of this BOT
	config *Config
}

// mid, message? ID
type mid string

type eventString string

type Location struct {
	Title string `json:title`
}

// When a user sends a message, the following data is sent to your server from the LINE platform.
type Message struct {

	// Identifier of the message.
	Id string `json:id,omitempty`
	// A numeric value indicating the type of message sent.
	ContentType int `json:contentType,omitempty`
	// MID of the user who sent the message.
	From string `json:from,omitempty`
	// Time and date request created. Displayed as the amount of time passed since 0:00:00 on January 1, 1970. The unit is given in milliseconds.
	CreatedTime       int `json:createdTime,omitempty`
	parsedCreatedTime *time.Time

	// Array of user who will receive the message.
	To []string `json:to,omitempty`
	// Type of user who will receive the message. (1: To user )
	ToType int `json:toType,omitempty`
	// Detailed information about the message
	// ContentMetadata

	// Posted text to be delivered. Note: users can send a message which has max 10,000 characters.
	Text string `json:text,omitempty`

	// Location data. This property is defined if the text message sent contains location data.
	Location *Location `json:location,omitempty`
}

// The LINE platform sends operation requests to your BOT API server when users perform actions such as adding your official account as friend.
type Operation struct {

	// Revision number of operation
	Revision int `json:revision,omitempty`
	// Type of operation
	OpType int `json:opType,omitempty`
	// Array of MIDs
	Params []*string `json:params,omitempty`
}

type Content struct {
	Message
	Operation
}

type Result struct {
	// Fixed value "u2ddf2eb3c959e561f6c9fa2ea732e7eb8"
	From string `json:from`
	// Fixed value "1341301815"
	FromChannel json.Number `json:"fromChannel"`
	// MID value granted by the BOT API server’s Channel
	To []mid `json:to`
	// Channel ID of the BOT API server
	ToChannel json.Number `json:toChannel`
	// Identifier used to show the type of data
	EventType eventString `json:eventType`
	// ID string to identify each event
	Id string `json:id`
	// Actual data relayed by the message
	Content Content `json:content`
}

type Request struct {

	// Array of target user. Max count: 150.
	To []string `json:to`
	// 1383378250 Fixed value
	ToChannel int `json:toChannel`
	// "138311608800106203" Fixed value.
	EventType eventString `json:eventType`
	// Object that contains the message (varies according to message type).
	Content Content `json:content`
}

// Return object for Callback Request
type CallbackRequest struct {
	vaild  bool
	Result []Result `json:result`
}

func ParseRequest(r *http.Request) (*CallbackRequest, error) {
	result := CallbackRequest{}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// checkSignature reports whether messageMAC is a valid HMAC tag for message.
func CheckSignature(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}

// Set default fixed value for request
func (r *Request) SetDefaults() {
	r.ToChannel = DefaultToChannel
	r.EventType = EventMessage
}

func (r *Request) AddTargetUser(mid string) error {
	if len(r.To) >= 150 {
		return ErrUserExceed
	}
	r.To = append(r.To, mid)
	return nil
}

func (r *Request) SetText(text string) error {
	r.Content.ToType = ToTypeUser
	r.Content.ContentType = TextMessage
	r.Content.Text = text
	return nil
}
