package bandwidth

import "encoding/json"

type message struct {
	From        string `json:"from"`
	To          string `json:"to"`
	Text        string `json:"text"`
	CallbackURL string `json:"callbackUrl"`
	Media 		[]string `json:"media`
}


func (b* Bandwidth) SendSMS(from string, to string, text string, callback string) error{
	var sms = &message{
		From:from,
		To:to,
		Text:text,
		CallbackURL:callback,
	}
	return b.post(sms, "messages")
}

func (b* Bandwidth) SendMMS (from string, to string, text string, mediaURL []string, callback string) error{
	var mms = &message{
		From:from,
		To:to,
		Text:text,
		CallbackURL:callback,
		Media:mediaURL,
	}
	return b.post(mms, "messages")
}


type MessageEvent struct {
	EventType           string `json:"eventType"`
	Direction           string `json:"direction"`
	From                string `json:"from"`
	To                  string `json:"to"`
	MessageID           string `json:"messageId"`
	MessageURI          string `json:"messageUri"`
	Text                string `json:"text"`
	ApplicationID       string `json:"applicationId"`
	Time                string `json:"time"`
	State               string `json:"state"`
	DeliveryState       string `json:"deliveryState"`
	DeliveryCode        string `json:"deliveryCode"`
	DeliveryDescription string `json:"deliveryDescription"`
}


func (b* Bandwidth) MessageEvent (body *[]byte) error {
	event := &MessageEvent{}
	err := json.Unmarshal(*body, event)

	if err != nil {
		return err
	}
	switch event.EventType {
	case "sms": b.SMSEvent(event)
	case "mms": b.MMSEvent(event)
	default: b.DefaultMessageEvent(event)
	}
	return nil
}