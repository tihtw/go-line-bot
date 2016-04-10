package lbot

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// The LINE BOT
type Bot struct {
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
	Id string `json:id`
	// A numeric value indicating the type of message sent.
	ContentType int `json:contentType`
	// MID of the user who sent the message.
	From string `json:from`
	// Time and date request created. Displayed as the amount of time passed since 0:00:00 on January 1, 1970. The unit is given in milliseconds.
	CreatedTime       int `json:createdTime`
	parsedCreatedTime *time.Time

	// Array of user who will receive the message.
	To []string `json:to`
	// Type of user who will receive the message. (1: To user )
	ToType int `json:toType`
	// Detailed information about the message
	// ContentMetadata

	// Posted text to be delivered. Note: users can send a message which has max 10,000 characters.
	Text string `json:text`

	// Location data. This property is defined if the text message sent contains location data.
	Location *Location `json:location`
}

// The LINE platform sends operation requests to your BOT API server when users perform actions such as adding your official account as friend.
type Operation struct {

	// Revision number of operation
	Revision int `json:revision`
	// Type of operation
	OpType int `json:opType`
	// Array of MIDs
	Params []*string `json:params`
}

type Content struct {
	Message
	Operation
}

type Result struct {
	// Fixed value "u2ddf2eb3c959e561f6c9fa2ea732e7eb8"
	From string `json:from`
	// Fixed value "1341301815"
	FromChannel string `json:fromChannel`
	// MID value granted by the BOT API server’s Channel
	To []mid `json:to`
	// Channel ID of the BOT API server
	ToChannel int `json:toChannel`
	// Identifier used to show the type of data
	EventType eventString `json:eventType`
	// ID string to identify each event
	Id string `json:id`
	// Actual data relayed by the message
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
