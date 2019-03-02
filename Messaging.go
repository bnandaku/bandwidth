package bandwidth


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