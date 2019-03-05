package bandwidth

import (
	"bytes"
	"github.com/json-iterator/go"
	"net/http"
)

import b64 "encoding/base64"

type Bandwidth struct {
	UserID         string
	Authorization  string
	Token          string
	Secret         string
	AnswerEvent    PhoneCallBack
	HangupEvent    PhoneCallBack
	IncomingEvent  PhoneCallBack
	RecordingEvent PhoneCallBack
	DefaultEvent   PhoneCallBack
}

type PhoneCallBack func (event *CallEvent) error

func (t* Bandwidth) New (UserID string, Token string, Secret string) error{
	t.Secret = Secret
	t.Token = Token
	t.UserID = UserID
	data := Token + ":" + Secret
	t.Authorization = b64.StdEncoding.EncodeToString([]byte(data))
	return t.getRequest("account")
}


func(t* Bandwidth) getRequest (endpoint string) error{
	url := "https://api.catapult.inetwork.com/v1/users/" + t.UserID + "/"+ endpoint
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Basic " + t.Authorization)
	_, err := http.DefaultClient.Do(req)
	return err
}


func (t* Bandwidth) post(body interface{}, endpoint string) error{
	url := "https://api.catapult.inetwork.com/v1/users/" + t.UserID + "/"+ endpoint
	json, err := jsoniter.Marshal(body)
	if err != nil {
		return err
	}
	req, _ := http.NewRequest("POST", url, bytes.NewReader(json))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic " + t.Authorization)
	_, err = http.DefaultClient.Do(req)
	return err
}