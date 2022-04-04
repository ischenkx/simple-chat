package app

import (
	"github.com/ischenkx/vk-test-task/internal/app/data/models"
	"github.com/ischenkx/vk-test-task/internal/app/errors"
	"github.com/ischenkx/vk-test-task/internal/app/event"
	"github.com/ischenkx/vk-test-task/internal/app/forms"
	"log"
	"time"
)

type Message interface {
	ID() string
	Sender(ctx *Context) (User, error)
	Chat(ctx *Context) (Chat, error)
	Payload(ctx *Context) (string, error)
	TimeStamp(ctx *Context) (time.Time, error)
	LastUpdate(ctx *Context) (time.Time, error)
	Update(ctx *Context, update forms.MessageUpdate) error
	Delete(ctx *Context) error
	Model(ctx *Context) (models.Message, error)
}

type message struct {
	app *App
	id  string
}

func (m message) isAccessible(ctx *Context) bool {
	if ctx.User() == nil {
		return false
	}
	if model, err := m.Model(ctx); err != nil {
		return false
	} else {
		return model.UserID == ctx.User().ID()
	}
}

func (m message) exists(ctx *Context) bool {
	if _, err := m.Model(ctx); err != nil {
		return false
	}
	return true
}

func (m message) Model(ctx *Context) (models.Message, error) {
	return m.app.repo.GetMessage(ctx, m.id)
}

func (m message) ID() string {
	return m.id
}

func (m message) Sender(ctx *Context) (User, error) {
	if !m.isAccessible(ctx) {
		return nil, errors.ResourceInaccessible
	}
	if model, err := m.Model(ctx); err != nil {
		return nil, err
	} else {
		return newUser(ctx, m.app, model.UserID)
	}
}

func (m message) Chat(ctx *Context) (Chat, error) {
	if !m.isAccessible(ctx) {
		return nil, errors.ResourceInaccessible
	}
	if model, err := m.Model(ctx); err != nil {
		return nil, err
	} else {
		return newChat(ctx, m.app, model.ChatID)
	}
}

func (m message) Payload(ctx *Context) (string, error) {
	if !m.isAccessible(ctx) {
		return "", errors.ResourceInaccessible
	}
	if model, err := m.Model(ctx); err != nil {
		return "", err
	} else {
		return model.Payload, nil
	}
}

func (m message) TimeStamp(ctx *Context) (time.Time, error) {
	if !m.isAccessible(ctx) {
		return time.Time{}, errors.ResourceInaccessible
	}
	if model, err := m.Model(ctx); err != nil {
		return time.Time{}, err
	} else {
		return model.TimeStamp, nil
	}
}

func (m message) LastUpdate(ctx *Context) (time.Time, error) {
	if !m.isAccessible(ctx) {
		return time.Time{}, errors.ResourceInaccessible
	}
	if model, err := m.Model(ctx); err != nil {
		return time.Time{}, err
	} else {
		return model.LastUpdate, nil
	}
}

func (m message) Update(ctx *Context, update forms.MessageUpdate) error {

	if !m.isAccessible(ctx) {
		return errors.ResourceInaccessible
	}

	if err := update.Validate(); err != nil {
		return err
	}

	model, err := m.Model(ctx)
	if err != nil {
		return err
	}
	model.Payload = update.Payload
	model.LastUpdate = time.Now()

	if err := m.app.repo.UpdateMessage(ctx, model); err != nil {
		return err
	}

	e := event.New(MessageUpdatedEventName, MessageUpdatedEvent{
		MessageID: model.ID,
		ChatID:    model.ChatID,
	}, event.WithTime(time.Now()))

	if err := m.app.Events().Send(ctx, e); err != nil {
		// currently not handled
		log.Println("failed to send event:", err)
	}

	return nil
}

func (m message) Delete(ctx *Context) error {
	if !m.isAccessible(ctx) {
		return errors.ResourceInaccessible
	}

	chat, err := m.Chat(ctx)
	if err != nil {
		return err
	}

	if err := m.app.repo.DeleteMessage(ctx, m.id); err != nil {
		return err
	}

	e := event.New(MessageDeletedEventName, MessageDeletedEvent{
		MessageID: m.id,
		ChatID:    chat.ID(),
	}, event.WithTime(time.Now()))

	if err := m.app.Events().Send(ctx, e); err != nil {
		// currently not handled
		log.Println("failed to send event:", err)
	}

	return nil
}

func unsafeMessageFromModel(app *App, m models.Message) Message {
	return message{
		app: app,
		id:  m.ID,
	}
}

func newMessage(ctx *Context, app *App, id string) (Message, error) {
	m := message{
		app: app,
		id:  id,
	}
	if !m.exists(ctx) {
		return nil, errors.DoesNotExist
	}
	return m, nil
}
