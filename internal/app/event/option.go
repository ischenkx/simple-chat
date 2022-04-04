package event

import "time"

type Option func(*Event)

func WithTime(time time.Time) Option {
	return func(event *Event) {
		event.TimeStamp = time.UnixNano()
	}
}
