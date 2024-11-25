package messagecontracts

import (
	"encoding/json"
	"fmt"
)

type EmailContent struct {
	Receiver string `json:"receiver"`
	Subject  string `json:"subject"`
	Body     string `json:"body"`
}

func NewEmailContentFromMessage(msg Message) (EmailContent, error) {
	if msg.Type != MessageTypeEmail {
		return EmailContent{}, fmt.Errorf("message is not email type")
	}

	var emailContent EmailContent
	err := json.Unmarshal([]byte(msg.Content), &emailContent)

	if err != nil {
		return EmailContent{}, err
	}

	return emailContent, nil
}
