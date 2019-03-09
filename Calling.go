package bandwidth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type call struct {
	From string `json:"from"`
	To string `json:"to"`
	Callback string `json:"callbackUrl"`
}

func (b* Bandwidth) CreateCall(from string, to string, callback string) error{
	newCall := &call{
		From:from,
		To:to,
		Callback:callback,
	}
	err, _ := b.post(newCall, "calls")
	return err
}


type CallEvent struct {
	EventType string `json:"eventType"`
	From      string `json:"from"`
	To        string `json:"to"`
	CallID    string `json:"callId"`
	CallURI   string `json:"callUri"`
	CallState string `json:"callState"`
	Time      string `json:"time"`
}

func (b* Bandwidth) CallEvent(body []byte) error {

	event := &CallEvent{}
	err := json.Unmarshal(body, event)

	if err != nil {
		return err
	}
	fmt.Println(event.EventType)
	switch event.EventType {
	case "incomingcall": b.IncomingEvent(event)
	case "answer": b.AnswerEvent(event)
	case "recording": b.RecordingEvent(event)
	case "hangup": b.HangupEvent(event)
	default: b.DefaultEvent(event)
	}
	return nil

}

func (b *Bandwidth) StartRecording(callID string, format string) error  {
	var body = json.RawMessage(`{"recordingEnabled":"true", "recordingFileFormat": "mp3"}`)
	if format == "wav"{
		body = json.RawMessage(`{"recordingEnabled":"true", "recordingFileFormat": "wav"}`)
	}
	fmt.Println("Starting to record")
	endpoint := "calls/" + callID
	err, _ := b.post(body,endpoint )
	return  err
}

func (b *Bandwidth) StopRecording(callID string) error{
	var body = json.RawMessage(`{"recordingEnabled":"false"}`)
	endpoint := "calls/" + callID
	err, _ := b.post(body,endpoint)
	return  err
}


func (b *Bandwidth) Hangup(callID string) error{
	var body = json.RawMessage(`{"state":"completed"}`)
	endpoint := "calls/" + callID
	err, _ := b.post(body,endpoint)
	return  err

}

type RecordingArray []struct {
	EndTime   time.Time `json:"endTime"`
	ID        string    `json:"id"`
	Media     string    `json:"media"`
	Call      string    `json:"call"`
	StartTime time.Time `json:"startTime"`
	State     string    `json:"state"`
}

func (b *Bandwidth) GetRecording(callID string) (error, RecordingArray ) {
	endpoint := "calls/" + callID + "/recordings"
	err, body := b.getRequest(endpoint)
	if err != nil {
		return err, RecordingArray{}
	}
	var recordingArr RecordingArray
	json.Unmarshal(body, &recordingArr)
	return err, recordingArr
}


func (b *Bandwidth) DownloadMedia(mediaURL string)(error, *http.Response) {

	req, _ := http.NewRequest("GET", mediaURL, nil)
	req.Header.Add("Authorization", "Basic " + b.Authorization)
	res, err := http.DefaultClient.Do(req)

	return err, res
}