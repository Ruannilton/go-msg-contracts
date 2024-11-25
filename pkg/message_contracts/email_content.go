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

func NewEmailMessage(receiver string, subject string, body string) (Message, error) {
	content := EmailContent{
		Receiver: receiver,
		Subject:  subject,
		Body:     body,
	}

	data, err := json.Marshal(content)

	if err != nil {
		return Message{}, err
	}

	msg := Message{
		Type:    MessageTypeSMS,
		Content: string(data),
	}

	return msg, nil
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
