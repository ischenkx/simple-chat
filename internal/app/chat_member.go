package app

import (
	"github.com/ischenkx/vk-test-task/internal/app/data/models"
	"github.com/ischenkx/vk-test-task/internal/app/errors"
	"github.com/ischenkx/vk-test-task/internal/app/event"
	"github.com/ischenkx/vk-test-task/internal/app/forms"
	"log"
	"time"
)

type ChatMember interface {
	Chat(ctx *Context) (Chat, error)
	User(ctx *Context) (User, error)
	Status(ctx *Context) (int, error)

	Delete(ctx *Context) error

	SendMessage(ctx *Context, form forms.SendMessage) (Message, error)
}

type chatMember struct {
	app    *App
	chatID string
	userID string
}

func (member chatMember) exists(ctx *Context) bool {
	if _, err := member.app.repo.GetChatMember(ctx, member.userID, member.chatID); err != nil {
		return false
	}
	return true
}

func (member chatMember) isWritable(ctx *Context) bool {
	if ctx.User() == nil {
		return false
	}
	c, err := member.Chat(ctx)
	if err != nil {
		return false
	}
	if _, err := c.Member(ctx, ctx.User().ID()); err != nil {
		return false
	}

	return true
}

func (member chatMember) Chat(ctx *Context) (Chat, error) {
	return newChat(ctx, member.app, member.chatID)
}

func (member chatMember) User(ctx *Context) (User, error) {
	return newUser(ctx, member.app, member.userID)
}

func (member chatMember) Status(ctx *Context) (int, error) {
	if !member.isWritable(ctx) {
		return 0, errors.ResourceInaccessible
	}
	m, err := member.app.repo.GetChatMember(ctx, member.userID, member.chatID)

	if err != nil {
		return 0, errors.DoesNotExist
	}

	return m.Status, nil
}

func (member chatMember) SendMessage(ctx *Context, form forms.SendMessage) (Message, error) {
	if !member.isWritable(ctx) {
		return nil, errors.ResourceInaccessible
	}

	if err := form.Validate(); err != nil {
		return nil, err
	}

	mes, err := member.app.repo.CreateMessage(ctx, models.Message{
		Payload:   form.Payload,
		TimeStamp: time.Now(),
		ChatID:    member.chatID,
		UserID:    member.userID,
	})
	if err != nil {
		return nil, err
	}

	e := event.New(NewMessageEventName, NewMessageEvent{
		MessageID: mes.ID,
		ChatID:    mes.ChatID,
	}, event.WithTime(time.Now()))

	if err := member.app.Events().Send(ctx, e); err != nil {
		// currently not handled
		log.Println("failed to send event:", err)
	}

	return unsafeMessageFromModel(member.app, mes), nil
}

func (member chatMember) Delete(ctx *Context) error {
	if !member.isWritable(ctx) {
		return errors.ResourceInaccessible
	}

	if err := member.app.repo.DeleteChatMember(ctx, member.userID, member.chatID); err != nil {
		return err
	}

	e := event.New(ChatMemberDeletedEventName, ChatMemberDeletedEvent{
		ChatID: member.chatID,
		UserID: member.userID,
	}, event.WithTime(time.Now()))

	if err := member.app.Events().Send(ctx, e); err != nil {
		// currently not handled
		log.Println("failed to send event:", err)
	}

	return nil
}

func unsafeChatMemberFromModel(app *App, model models.ChatMember) ChatMember {
	return chatMember{
		app:    app,
		userID: model.UserID,
		chatID: model.ChatID,
	}
}

func newChatMember(ctx *Context, app *App, userId, chatId string) (ChatMember, error) {
	m := chatMember{
		app:    app,
		userID: userId,
		chatID: chatId,
	}
	if !m.exists(ctx) {
		return chatMember{}, errors.DoesNotExist
	}
	return m, nil
}
