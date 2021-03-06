package lbot

import (
	"encoding/json"
	"fmt"
	// "io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
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

	// payload = `null`
	req, _ := http.NewRequest("POST", "", strings.NewReader(string(payload)))
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("User-Agent", "ChannelEventDispatcher/1.0")

	actual, err := ParseRequest(req)
	if err != nil {
		t.Error(err)
		t.FailNow()
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

	if actual.Result[0].Content.Location != nil {
		t.Error("actual.Result[0].Content.Location != nil")
	}

	if actual.Result[0].Content.Id != "325708" {
		t.Error("actual.Result[0].Content.Id != 325708")
	}

	if actual.Result[0].Content.ContentType != 1 {
		t.Errorf("actual.Result[0].Content.ContentType != %v", 1)
	}
	if actual.Result[0].Content.From != "uff2aec188e58752ee1fb0f9507c6529a" {
		t.Errorf("actual.Result[0].Content.From != %v", "uff2aec188e58752ee1fb0f9507c6529a")
	}
	if actual.Result[0].Content.Text != "Hello, BOT API Server!" {
		t.Errorf("actual.Result[0].Content.Text != %v", "Hello, BOT API Server!")
	}

	// fmt.Printf("%s\n", out)

}

func TestParseRequest2(t *testing.T) {
	payload := `
{"result":[
  {
    "from":"u2ddf2eb3c959e561f6c9fa2ea732e7eb8",
    "fromChannel":1341301815,
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

	// payload = `null`
	req, _ := http.NewRequest("POST", "", strings.NewReader(string(payload)))
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("User-Agent", "ChannelEventDispatcher/1.0")

	actual, err := ParseRequest(req)
	if err != nil {
		t.Error(err)
		t.FailNow()
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

	if actual.Result[0].Content.Location != nil {
		t.Error("actual.Result[0].Content.Location != nil")
	}

	if actual.Result[0].Content.Id != "325708" {
		t.Error("actual.Result[0].Content.Id != 325708")
	}

	if actual.Result[0].Content.ContentType != 1 {
		t.Errorf("actual.Result[0].Content.ContentType != %v", 1)
	}
	if actual.Result[0].Content.From != "uff2aec188e58752ee1fb0f9507c6529a" {
		t.Errorf("actual.Result[0].Content.From != %v", "uff2aec188e58752ee1fb0f9507c6529a")
	}
	if actual.Result[0].Content.Text != "Hello, BOT API Server!" {
		t.Errorf("actual.Result[0].Content.Text != %v", "Hello, BOT API Server!")
	}

	// fmt.Printf("%s\n", out)

}

func TestParseOperationRequest(t *testing.T) {
	payload := `
{"result":[
  {
    "from":"u2ddf2eb3c959e561f6c9fa2ea732e7eb8",
    "fromChannel":"1341301815",
    "to":["u0cc15697597f61dd8b01cea8b027050e"],
    "toChannel":1441301333,
    "eventType":"138311609100106403",
    "id":"ABCDEF-12345678901",
    "content": {
      "params":[
        "u0f3bfc598b061eba02183bfc5280886a",
        null,
        null
      ],
      "revision":2469,
      "opType":4
    }
  }
]}`

	// payload = `null`
	req, _ := http.NewRequest("POST", "", strings.NewReader(string(payload)))
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("User-Agent", "ChannelEventDispatcher/1.0")

	actual, err := ParseRequest(req)
	if err != nil {
		t.Error(err)
		t.FailNow()
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

	if actual.Result[0].Content.Location != nil {
		t.Error("actual.Result[0].Content.Location != nil")
	}

	if actual.Result[0].Content.OpType != 4 {
		t.Errorf("actual.Result[0].Content.opType != %v", 4)
	}

	if *actual.Result[0].Content.Params[0] != "u0f3bfc598b061eba02183bfc5280886a" {
		t.Errorf("*actual.Result[0].Content.Params[0] != %v", "u0f3bfc598b061eba02183bfc5280886a")
	}

	if actual.Result[0].Content.Params[1] != nil {
		t.Errorf("actual.Result[0].Content.Params[1] != %v", "u0f3bfc598b061eba02183bfc5280886a")
	}

	if actual.Result[0].Content.Params[2] != nil {
		t.Errorf("actual.Result[0].Content.Params[2] != %v", nil)
	}

	// fmt.Printf("%s\n", out)

}

func TestMarshalRequest(t *testing.T) {
	var payload Request
	payload.SetDefaults()
	payload.SetText("おはいよ")
	payload.AddTargetUser("mid")

	out, err := json.Marshal(payload)
	if err != nil {
		t.Error(err)
	}
	var target = `{"to":["mid"],"toChannel":1383378250,"eventType":"138311608800106203","content":{"contentType":1,"toType":1,"text":"おはいよ"}}`

	if target != string(out) {
		t.Errorf("Result mismatch, expected %v, actual %v\n", target, string(out))
	}
}

func TestParseProfile(t *testing.T) {
	payload := `{
    "contacts": [
        {
            "displayName": "BOT API",
            "mid": "u0047556f2e40dba2456887320ba7c76d",
            "pictureUrl": "http://dl.profile.line.naver.jp/abcdefghijklmn",
            "statusMessage": "Hello, LINE!"
        }
    ],
    "count": 1,
    "display": 1,
    "pagingRequest": {
        "display": 1,
        "sortBy": "MID",
        "start": 1
    },
    "start": 1,
    "total": 1
}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, payload)
	}))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}

	actual, err := ParseProfileResponse(res)
	if err != nil {
		log.Fatal(err)
	}

	if actual.Count != 1 {
		t.Errorf("actual.Count!= %v", 1)
	}
	if actual.Display != 1 {
		t.Errorf("actual.Display!= %v", 1)
	}
	if actual.Start != 1 {
		t.Errorf("actual.Start!= %v", 1)
	}

	if actual.Total != 1 {
		t.Errorf("actual.Total!= %v", 1)
	}

	if len(actual.Contacts) != 1 {
		t.Errorf("len(actual.Contacts) != %v", 1)
	}

	if actual.Contacts[0].MID != "u0047556f2e40dba2456887320ba7c76d" {
		t.Errorf("actual.Contacts[0].MID != %v", "u0047556f2e40dba2456887320ba7c76d")
	}

}

func TestSetText(t *testing.T) {
	var r Request
	r.SetText("text")
	if r.Content.ToType != ToTypeUser {
		t.Errorf("r.Content.ToType != %v\n", ToTypeUser)
	}
	if r.Content.ContentType != TextMessage {
		t.Errorf("r.Content.ContentType != %v\n", TextMessage)
	}

	if r.Content.Text != "text" {
		t.Errorf("r.Content.Text != %v\n", "text")
	}

}

func TestSetImage(t *testing.T) {
	var r Request
	r.SetImage("originalContentUrl", "previewImageUrl")
	if r.Content.ToType != ToTypeUser {
		t.Errorf("r.Content.ToType != %v\n", ToTypeUser)
	}
	if r.Content.ContentType != ImageMessage {
		t.Errorf("r.Content.ContentType != %v\n", ImageMessage)
	}

	if r.Content.OriginalContentUrl != "originalContentUrl" {
		t.Errorf("r.Content.OriginalContentUrl != %v\n", "originalContentUrl")
	}

	if r.Content.PreviewImageUrl != "previewImageUrl" {
		t.Errorf("r.Content.PreviewImageUrl != %v\n", "previewImageUrl")
	}

}
