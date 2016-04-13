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
	Title     string  `json:"title"`
	Address   string  `json:"Address,omitempty"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type imageContent struct {
	OriginalContentUrl string `json:"originalContentUrl,omitempty"`
	PreviewImageUrl    string `json:"previewImageUrl,omitempty"`
}

type contentMetadata struct {
	Stkid    string `json:"STKID"`
	Stkpkgid string `json:"STKPKGID"`
	Stkver   string `json:"Stkver"`
}

type Sticker struct {
	Stkid    string
	Stkpkgid string
	Stkver   string
}

// When a user sends a message, the following data is sent to your server from the LINE platform.
type Message struct {

	// Identifier of the message.
	Id string `json:"id,omitempty"`
	// A numeric value indicating the type of message sent.
	ContentType int `json:"contentType,omitempty"`
	// MID of the user who sent the message.
	From mid `json:"from,omitempty"`
	// Time and date request created. Displayed as the amount of time passed since 0:00:00 on January 1, 1970. The unit is given in milliseconds.
	CreatedTime       int `json:"createdTime,omitempty"`
	parsedCreatedTime *time.Time

	// Array of user who will receive the message.
	To []mid `json:"to,omitempty"`
	// Type of user who will receive the message. (1: To user )
	ToType int `json:"toType,omitempty"`
	// Detailed information about the message
	ContentMetadata *contentMetadata `json:"contentMetadata,omitempty"`

	// Posted text to be delivered. Note: users can send a message which has max 10,000 characters.
	Text string `json:"text,omitempty"`
	imageContent
	// Location data. This property is defined if the text message sent contains location data.
	Location *Location `json:"location,omitempty"`
}

// The LINE platform sends operation requests to your BOT API server when users perform actions such as adding your official account as friend.
type Operation struct {

	// Revision number of operation
	Revision int `json:"revision,omitempty"`
	// Type of operation
	OpType int `json:"opType,omitempty"`
	// Array of MIDs
	Params []*string `json:"params,omitempty"`
}

type Content struct {
	Message
	Operation
}

type Result struct {
	// Fixed value "u2ddf2eb3c959e561f6c9fa2ea732e7eb8"
	From mid `json:from`
	// Fixed value "1341301815"
	FromChannel json.Number `json:"fromChannel"`
	// MID value granted by the BOT API server’s Channel
	To []mid `json:"to"`
	// Channel ID of the BOT API server
	ToChannel json.Number `json:"toChannel"`
	// Identifier used to show the type of data
	EventType eventString `json:eventType`
	// ID string to identify each event
	Id string `json:"id"`
	// Actual data relayed by the message
	Content Content `json:"content"`
}

type Request struct {

	// Array of target user. Max count: 150.
	To []mid `json:"to"`
	// 1383378250 Fixed value
	ToChannel int `json:"toChannel"`
	// "138311608800106203" Fixed value.
	EventType eventString `json:"eventType"`
	// Object that contains the message (varies according to message type).
	Content Content `json:"content"`
}

// Return object for Callback Request
type CallbackRequest struct {
	vaild  bool
	Result []Result `json:"result"`
}

type UserProfileResponse struct {
	// contacts
	Contacts []ProfileInfo `json:"contacts"`
	Count    int           `json:"count"`
	Total    int           `json:"total"`
	Start    int           `json:"start"`
	Display  int           `json:"display"`
}

type ProfileInfo struct {
	DisplayName   string `json:"displayName"`
	MID           mid    `json:"mid"`
	pictureUrl    string `json:"pictureUrl"`
	statusMessage string `json:"statusMessage"`
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

func ParseProfileResponse(r *http.Response) (*UserProfileResponse, error) {
	result := UserProfileResponse{}
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
	r.EventType = EventSendMessage
}

func (r *Request) AddTargetUser(m mid) error {
	if len(r.To) >= 150 {
		return ErrUserExceed
	}
	r.To = append(r.To, m)
	return nil
}

func (r *Request) SetText(text string) error {
	r.Content.ToType = ToTypeUser
	r.Content.ContentType = TextMessage
	r.Content.Text = text
	return nil
}

func (r *Request) SetImage(originalContentUrl, previewImageUrl string) error {
	r.Content.ToType = ToTypeUser
	r.Content.ContentType = ImageMessage
	r.Content.OriginalContentUrl = originalContentUrl
	r.Content.PreviewImageUrl = previewImageUrl
	return nil
}

func (r *Request) SetVideo(originalContentUrl, previewImageUrl string) error {
	r.Content.ToType = ToTypeUser
	r.Content.ContentType = VideoMessage
	r.Content.OriginalContentUrl = originalContentUrl
	r.Content.PreviewImageUrl = previewImageUrl
	return nil
}

func (r *Request) SetAudio(originalContentUrl, previewImageUrl string) error {
	r.Content.ToType = ToTypeUser
	r.Content.ContentType = AudioMessage
	r.Content.OriginalContentUrl = originalContentUrl
	r.Content.PreviewImageUrl = previewImageUrl
	return nil
}

func (r *Request) SetLocation(text, title string, latitude, longitude string) error {
	r.Content.ToType = ToTypeUser
	r.Content.ContentType = LocationMessage
	r.Content.Text = text
	r.Content.Location = new(Location)
	r.Content.Location.Title = title
	r.Content.Location.Latitude = latitude
	r.Content.Location.Longitude = longitude
	return nil
}

func (r *Request) SetSticker(s *Sticker) error {
	r.Content.ToType = ToTypeUser
	r.Content.ContentType = StickerMessage

	r.Content.ContentMetadata = new(contentMetadata)
	r.Content.ContentMetadata.Stkid = s.Stkid
	r.Content.ContentMetadata.Stkpkgid = s.Stkpkgid
	r.Content.ContentMetadata.Stkver = s.Stkver
	return nil
}
