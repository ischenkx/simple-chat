package app

import (
	"github.com/ischenkx/vk-test-task/internal/app/data"
	"github.com/ischenkx/vk-test-task/internal/app/data/models"
	"github.com/ischenkx/vk-test-task/internal/app/errors"
	"github.com/ischenkx/vk-test-task/internal/app/forms"
)

type ChatManager struct {
	app *App
}

func (manager ChatManager) Get(ctx *Context, id string) (Chat, error) {
	return newChat(ctx, manager.app, id)
}

func (manager ChatManager) Create(ctx *Context, form forms.ChatCreationForm) (Chat, error) {
	if ctx.User() == nil {
		return nil, errors.NotAuthorized
	}

	if err := form.Validate(); err != nil {
		return nil, err
	}

	c, err := manager.app.repo.Transaction(ctx, func(tx data.Tx) (interface{}, error) {
		c, err := tx.CreateChat(ctx, models.Chat{
			Name:        form.Name,
			Description: form.Description,
			OwnerID:     ctx.User().ID(),
		})

		if err != nil {
			return nil, err
		}

		_, err = tx.CreateChatMember(ctx, models.ChatMember{
			ChatID: c.ID,
			UserID: ctx.User().ID(),
			Status: 1,
		})

		return c, err
	})

	if err != nil {
		return nil, err
	}
	return unsafeChatFromModel(manager.app, c.(models.Chat)), nil
}

func (manager ChatManager) GetMessage(ctx *Context, id string) (Message, error) {
	if ctx.User() == nil {
		return nil, errors.NotAuthorized
	}

	mes, err := manager.app.repo.GetMessage(ctx, id)

	if err != nil {
		return nil, err
	}

	if mes.UserID != ctx.User().ID() {
		return nil, errors.ResourceInaccessible
	}

	return unsafeMessageFromModel(manager.app, mes), nil
}
