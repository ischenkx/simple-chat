package event

type Event struct {
	Name      string
	Data      interface{}
	TimeStamp int64
}

func New(name string, data interface{}, options ...Option) Event {
	event := Event{
		Name: name,
		Data: data,
	}
	for _, option := range options {
		option(&event)
	}
	return event
}
