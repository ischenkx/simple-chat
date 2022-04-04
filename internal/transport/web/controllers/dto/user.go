package dto

import "github.com/ischenkx/vk-test-task/internal/app"

type User struct {
	Username string `json:"username"`
	ID       string `json:"id"`
}

func (dto *User) Load(ctx *app.Context, req app.User) error {
	model, err := req.Model(ctx)

	if err != nil {
		return err
	}
	dto.Username = model.Username
	dto.ID = model.ID
	return nil
}
