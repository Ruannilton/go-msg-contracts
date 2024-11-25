package messagecontracts

import (
	"encoding/json"
	"fmt"
)

type SMSContent struct {
	Number string `json:"number"`
	Body   string `json:"body"`
}

func NewSMSContentFromMessage(msg Message) (SMSContent, error) {
	if msg.Type != MessageTypeEmail {
		return SMSContent{}, fmt.Errorf("message is not sms type")
	}

	var smsContent SMSContent
	err := json.Unmarshal([]byte(msg.Content), &smsContent)

	if err != nil {
		return SMSContent{}, err
	}

	return smsContent, nil
}
