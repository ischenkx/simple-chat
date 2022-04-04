package evbus

import (
	"context"
	"errors"
	"github.com/ischenkx/vk-test-task/internal/app/event"
	"sync"
)

type handle struct {
	channel <-chan event.Event
	bus     *Bus
	closed  bool
	id      int64
}

func (h *handle) Chan(ctx context.Context) (<-chan event.Event, error) {
	if h.closed {
		return nil, errors.New("already closed")
	}
	return h.channel, nil
}

func (h *handle) Close(ctx context.Context) error {
	if h.closed {
		return errors.New("already closed")
	}
	h.closed = true
	h.bus.deleteReader(h.id)
	return nil
}

type Bus struct {
	readers map[int64]chan event.Event
	seq     int64
	mu      sync.RWMutex
}

func (b *Bus) deleteReader(id int64) {
	b.mu.Lock()
	defer b.mu.Unlock()
	delete(b.readers, id)
}

func (b *Bus) Channel(ctx context.Context) (event.ChannelHandle, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	ch := make(chan event.Event, 256)
	b.readers[b.seq] = ch
	h := &handle{
		channel: ch,
		bus:     b,
		closed:  false,
		id:      b.seq,
	}

	b.seq += 1
	return h, nil
}

func (b *Bus) Send(ctx context.Context, event event.Event) error {
	b.mu.RLock()
	defer b.mu.RUnlock()

	for _, reader := range b.readers {
		select {
		case reader <- event:
		default:
			// no deadlock, but events are lost...
		}
	}

	return nil
}

func NewBus() event.Bus {
	return &Bus{
		readers: map[int64]chan event.Event{},
		seq:     0,
		mu:      sync.RWMutex{},
	}
}
