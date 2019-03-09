package bandwidth

import "encoding/json"

type call struct {
	From string `json:"from"`
	To string `json:"to"`
	Callback string `json:"callbackurl"`
}

func (b* Bandwidth) CreateCall(from string, to string, callback string) error{
	newCall := &call{
		From:from,
		To:to,
		Callback:callback,
	}
	return b.post(newCall, "call")
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
	switch event.EventType {
	case "incomingcall": b.IncomingEvent(event)
	case "answer": b.AnswerEvent(event)
	case "recording": b.RecordingEvent(event)
	case "hangup": b.HangupEvent(event)
	default: b.DefaultEvent(event)
	}
	return nil

}