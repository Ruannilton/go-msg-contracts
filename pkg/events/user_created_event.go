package events

import (
	"time"

	"github.com/Ruannilton/go-msg-contracts/pkg/queues"
)

type CostumerCreatedEvent struct {
	Id        int
	Name      string
	Cpf       string
	Birthdate time.Time
	Phone     string
	Email     string
}

func (e CostumerCreatedEvent) GetEventName() string {
	return "UserCreatedEvent"
}

func (e CostumerCreatedEvent) GetEventQueue() string {
	return queues.CostumerCreatedEventQueue
}
