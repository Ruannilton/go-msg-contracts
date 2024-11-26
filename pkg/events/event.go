package events

type IEvent interface {
	GetEventName() string
	GetEventQueue() string
}
