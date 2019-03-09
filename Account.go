package bandwidth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	SMSEvent 	MessageCallBack
	MMSEvent 	MessageCallBack
	DefaultMessageEvent	MessageCallBack

}

type PhoneCallBack func (event *CallEvent) error
type MessageCallBack func(event *MessageEvent) error

func (t* Bandwidth) New (UserID string, Token string, Secret string) error{
	t.Secret = Secret
	t.Token = Token
	t.UserID = UserID
	data := Token + ":" + Secret
	t.Authorization = b64.StdEncoding.EncodeToString([]byte(data))
	err, _ := t.getRequest("account")
	return err
}


func(t* Bandwidth) getRequest (endpoint string) (error, []byte){
	url := "https://api.catapult.inetwork.com/v1/users/" + t.UserID + "/"+ endpoint
	fmt.Println(url)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Basic " + t.Authorization)
	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	respBody, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(respBody))
	return err, respBody
}


func (t* Bandwidth) post(body interface{}, endpoint string) (error, []byte){
	url := "https://api.catapult.inetwork.com/v1/users/" + t.UserID + "/"+ endpoint
	json, err := json.Marshal(body)
	if err != nil {
		return err, nil
	}
	req, _ := http.NewRequest("POST", url, bytes.NewReader(json))
	fmt.Println(url)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic " + t.Authorization)
	res, err := http.DefaultClient.Do(req)
	fmt.Println(res.StatusCode)

	defer res.Body.Close()
	respBody, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(respBody))
	return err, respBody
}