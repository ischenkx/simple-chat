package dto

import (
	"github.com/ischenkx/vk-test-task/internal/app"
	"time"
)

type Message struct {
	Payload    string    `json:"payload"`
	ChatID     string    `json:"chat_id"`
	UserID     string    `json:"user_id"`
	LastUpdate time.Time `json:"last_update"`
	TimeStamp  time.Time `json:"time_stamp"`
	ID         string    `json:"id"`
}

func (dto *Message) Load(ctx *app.Context, req app.Message) error {
	model, err := req.Model(ctx)

	if err != nil {
		return err
	}
	dto.ChatID = model.ChatID
	dto.UserID = model.UserID
	dto.ID = model.ID
	dto.Payload = string(model.Payload)
	dto.TimeStamp = model.TimeStamp
	dto.LastUpdate = model.LastUpdate
	return nil
}
