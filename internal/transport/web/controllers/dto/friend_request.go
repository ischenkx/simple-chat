package dto

import (
	"github.com/ischenkx/vk-test-task/internal/app"
	"time"
)

type FriendRequest struct {
	FromID    string    `json:"from_id"`
	ToID      string    `json:"to_id"`
	ID        string    `json:"id"`
	TimeStamp time.Time `json:"time_stamp"`
}

func (dto *FriendRequest) Load(ctx *app.Context, req app.FriendRequest) error {
	model, err := req.Model(ctx)

	if err != nil {
		return err
	}
	dto.FromID = model.From
	dto.ToID = model.To
	dto.ID = model.ID
	dto.TimeStamp = model.Time
	return nil
}
