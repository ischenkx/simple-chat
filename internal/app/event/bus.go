package event

import "context"

type ChannelHandle interface {
	Chan(ctx context.Context) (<-chan Event, error)
	Close(ctx context.Context) error
}

type Bus interface {
	Channel(ctx context.Context) (ChannelHandle, error)
	Send(ctx context.Context, event Event) error
}
