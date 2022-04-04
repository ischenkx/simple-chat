package dto

import (
	"github.com/ischenkx/vk-test-task/internal/app"
)

type Chat struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	OwnerID     string `json:"owner_id"`
	ID          string `json:"id"`
}

func (dto *Chat) Load(ctx *app.Context, chat app.Chat) error {
	model, err := chat.Model(ctx)

	if err != nil {
		return err
	}

	dto.Name = model.Name
	dto.ID = model.ID
	dto.Description = model.Description
	dto.OwnerID = model.OwnerID

	return nil
}
