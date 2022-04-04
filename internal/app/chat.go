package app

import (
	"github.com/ischenkx/vk-test-task/internal/app/data/models"
	"github.com/ischenkx/vk-test-task/internal/app/errors"
	"github.com/ischenkx/vk-test-task/internal/app/event"
	"log"
	"time"
)

type Chat interface {
	ID() string
	Name(ctx *Context) (string, error)
	Description(ctx *Context) (string, error)
	Owner(ctx *Context) (User, error)
	Members(ctx *Context, offset int, amount int) ([]ChatMember, error)
	CountMembers(ctx *Context) (int, error)
	Member(ctx *Context, id string) (ChatMember, error)
	Add(ctx *Context, id string, status int) (ChatMember, error)

	Model(ctx *Context) (models.Chat, error)

	Messages(ctx *Context, offset int, amount int) ([]Message, error)
	CountMessages(ctx *Context) (int, error)

	Delete(ctx *Context) error
}

type chat struct {
	id  string
	app *App
}

func (c chat) exists(ctx *Context) bool {
	if _, err := c.app.repo.GetChat(ctx, c.id); err != nil {
		return false
	}
	return true
}

func (c chat) isAccessible(ctx *Context) bool {
	if ctx.User() == nil {
		return false
	}
	userMember, err := c.Member(ctx, ctx.User().ID())
	if userMember == nil || err != nil {
		return false
	}
	return true
}

func (c chat) Model(ctx *Context) (models.Chat, error) {
	if !c.isAccessible(ctx) {
		return models.Chat{}, errors.ResourceInaccessible
	}
	return c.app.repo.GetChat(ctx, c.id)
}

func (c chat) ID() string {
	return c.id
}

func (c chat) Name(ctx *Context) (string, error) {
	if m, err := c.Model(ctx); err != nil {
		return "", err
	} else {
		return m.Name, nil
	}
}

func (c chat) Description(ctx *Context) (string, error) {
	if m, err := c.Model(ctx); err != nil {
		return "", err
	} else {
		return m.Description, nil
	}
}

func (c chat) Owner(ctx *Context) (User, error) {
	if m, err := c.Model(ctx); err != nil {
		return nil, err
	} else {
		return newUser(ctx, c.app, m.OwnerID)
	}
}

func (c chat) Members(ctx *Context, offset int, count int) ([]ChatMember, error) {
	if !c.isAccessible(ctx) {
		return nil, errors.ResourceInaccessible
	}

	members, err := c.app.repo.GetChatMembers(ctx, c.id, offset, count)

	if err != nil {
		return nil, err
	}

	chatMembers := make([]ChatMember, 0, len(members))

	for i := 0; i < len(members); i++ {
		m := members[i]
		chatMembers = append(chatMembers, unsafeChatMemberFromModel(c.app, m))
	}

	return chatMembers, nil
}

func (c chat) CountMembers(ctx *Context) (int, error) {
	if !c.isAccessible(ctx) {
		return 0, errors.ResourceInaccessible
	}

	return c.app.repo.CountChatMembers(ctx, c.id)
}

func (c chat) Member(ctx *Context, id string) (ChatMember, error) {
	return newChatMember(ctx, c.app, id, c.id)
}

func (c chat) Add(ctx *Context, id string, status int) (ChatMember, error) {
	if !c.isAccessible(ctx) {
		return nil, errors.ResourceInaccessible
	}

	_, err := c.app.repo.CreateChatMember(ctx, models.ChatMember{
		ChatID: c.id,
		UserID: id,
		Status: status,
	})

	if err != nil {
		return nil, err
	}

	e := event.New(ChatMemberCreatedEventName, ChatMemberCreatedEvent{
		UserID: id,
		ChatID: c.id,
	}, event.WithTime(time.Now()))

	if err := c.app.Events().Send(ctx, e); err != nil {
		// currently not handled
		log.Println("failed to send event:", err)
	}

	return newChatMember(ctx, c.app, id, c.id)
}

func (c chat) Messages(ctx *Context, offset int, count int) ([]Message, error) {
	if !c.isAccessible(ctx) {
		return nil, errors.ResourceInaccessible
	}

	messages, err := c.app.repo.GetChatMessages(ctx, c.id, offset, count)

	if err != nil {
		return nil, err
	}

	chatMessages := make([]Message, 0, len(messages))

	for i := 0; i < len(messages); i++ {
		m := messages[i]
		chatMessages = append(chatMessages, unsafeMessageFromModel(c.app, m))
	}

	return chatMessages, nil
}

func (c chat) CountMessages(ctx *Context) (int, error) {
	if !c.isAccessible(ctx) {
		return 0, errors.ResourceInaccessible
	}

	return c.app.repo.CountChatMessages(ctx, c.id)
}

func (c chat) Delete(ctx *Context) error {
	if !c.isAccessible(ctx) {
		return errors.ResourceInaccessible
	}

	owner, err := c.Owner(ctx)

	if err != nil {
		return err
	}

	if owner.ID() != ctx.User().ID() {
		return errors.ResourceInaccessible
	}

	err = c.app.repo.DeleteChat(ctx, c.id)

	if err != nil {
		return err
	}

	e := event.New(ChatDeletedEventName, ChatDeletedEvent{
		ChatID: c.id,
	}, event.WithTime(time.Now()))

	if err := c.app.Events().Send(ctx, e); err != nil {
		// currently not handled
		log.Println("failed to send event:", err)
	}

	return nil
}

func unsafeChatFromModel(app *App, c models.Chat) chat {
	return chat{
		app: app,
		id:  c.ID,
	}
}

func newChat(ctx *Context, app *App, id string) (Chat, error) {
	c := chat{
		app: app,
		id:  id,
	}
	if !c.exists(ctx) {
		return nil, errors.DoesNotExist
	}
	return c, nil
}
