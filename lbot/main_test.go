package lbot

import (
	"net/http"
	"strings"
	"testing"
)

const (
	channelSecret = "00000000000000000000000000000000"
)

func TestParseRequest(t *testing.T) {
	payload := `
{"result":[
  {
    "from":"u2ddf2eb3c959e561f6c9fa2ea732e7eb8",
    "fromChannel":"1341301815",
    "to":["u0cc15697597f61dd8b01cea8b027050e"],
    "toChannel":1441301333,
    "eventType":"138311609000106303",
    "id":"ABCDEF-12345678901",
    "content": {
      "location":null,
      "id":"325708",
      "contentType":1,
      "from":"uff2aec188e58752ee1fb0f9507c6529a",
      "createdTime":1332394961610,
      "to":["u0a556cffd4da0dd89c94fb36e36e1cdc"],
      "toType":1,
      "contentMetadata":null,
      "text":"Hello, BOT API Server!"
    }
  }
]}`
	req, _ := http.NewRequest("POST", "", strings.NewReader(string(payload)))
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("User-Agent", "ChannelEventDispatcher/1.0")

	actual, err := ParseRequest(req)
	if err != nil {
		t.Error(err)
	}
	if actual.Result == nil {
		t.Error("Result == nil")
	}

	if len(actual.Result) != 1 {
		t.Error("Result != 1")
	}

	if actual.Result[0].From != "u2ddf2eb3c959e561f6c9fa2ea732e7eb8" {
		t.Error("actual.Result[0].From != u2ddf2eb3c959e561f6c9fa2ea732e7eb8")
	}

	// fmt.Printf("%s\n", out)

}
