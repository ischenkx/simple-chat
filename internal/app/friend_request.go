package app

import (
	"github.com/ischenkx/vk-test-task/internal/app/data"
	"github.com/ischenkx/vk-test-task/internal/app/data/models"
	"github.com/ischenkx/vk-test-task/internal/app/errors"
	"github.com/ischenkx/vk-test-task/internal/app/event"
	"log"
	"time"
)

type FriendRequest interface {
	ID() string
	From(ctx *Context) (User, error)
	To(ctx *Context) (User, error)
	Accept(ctx *Context) error
	Decline(ctx *Context) error
	Delete(ctx *Context) error
	Model(ctx *Context) (models.FriendRequest, error)
}

type friendRequest struct {
	id  string
	app *App
}

func (f friendRequest) Model(ctx *Context) (models.FriendRequest, error) {
	return f.app.repo.GetFriendRequestByID(ctx, f.id)
}

func (f friendRequest) exists(ctx *Context) bool {
	if _, err := f.Model(ctx); err != nil {
		return false
	}
	return true
}

func (f friendRequest) isAccessible(ctx *Context) bool {
	if ctx.User() == nil {
		return false
	}
	if model, err := f.Model(ctx); err == nil {
		return model.To == ctx.User().ID() || model.From == ctx.User().ID()
	}
	return false
}

func (f friendRequest) ID() string {
	return f.id
}

func (f friendRequest) From(ctx *Context) (User, error) {
	if !f.isAccessible(ctx) {
		return nil, errors.ResourceInaccessible
	}

	if model, err := f.Model(ctx); err != nil {
		return nil, err
	} else {
		return newUser(ctx, f.app, model.From)
	}
}

func (f friendRequest) To(ctx *Context) (User, error) {
	if !f.isAccessible(ctx) {
		return nil, errors.ResourceInaccessible
	}

	if model, err := f.Model(ctx); err != nil {
		return nil, err
	} else {
		return newUser(ctx, f.app, model.To)
	}
}

func (f friendRequest) Accept(ctx *Context) error {
	if ctx.User() == nil {
		return errors.NotAuthorized
	}

	model, err := f.Model(ctx)

	if err != nil {
		return err
	}

	if model.To != ctx.User().ID() {
		return errors.RightsViolation
	}

	_, err = f.app.repo.Transaction(ctx, func(repo data.Tx) (interface{}, error) {
		if err := f.app.repo.CreateFriendConnection(ctx, model.From, model.To); err != nil {
			return nil, err
		}

		if err := f.app.repo.DeleteFriendRequest(ctx, model.ID); err != nil {
			return nil, err
		}
		return nil, nil
	})

	if err != nil {
		return err
	}

	e := event.New(FriendAddedEventName, FriendAddedEvent{
		FriendID: model.From,
		UserID:   model.To,
	}, event.WithTime(time.Now()))

	if err := f.app.Events().Send(ctx, e); err != nil {
		// currently not handled
		log.Println("failed to send event:", err)
	}

	e = event.New(FriendRequestUpdateEventName, FriendRequestUpdateEvent{
		FriendRequestID: model.ID,
		From:            model.To,
		To:              model.From,
		Code:            FriendRequestUpdateAccepted,
	}, event.WithTime(time.Now()))

	if err := f.app.Events().Send(ctx, e); err != nil {
		// currently not handled
		log.Println("failed to send event:", err)
	}

	return nil
}

func (f friendRequest) Decline(ctx *Context) error {
	if ctx.User() == nil {
		return errors.NotAuthorized
	}
	model, err := f.Model(ctx)
	if err != nil {
		return err
	}
	if model.To != ctx.User().ID() {
		return errors.RightsViolation
	}
	if err := f.app.repo.DeleteFriendRequest(ctx, model.ID); err != nil {
		return err
	}

	e := event.New(FriendRequestUpdateEventName, FriendRequestUpdateEvent{
		FriendRequestID: model.ID,
		From:            model.To,
		To:              model.From,
		Code:            FriendRequestUpdateDeclined,
	}, event.WithTime(time.Now()))

	if err := f.app.Events().Send(ctx, e); err != nil {
		// currently not handled
		log.Println("failed to send event:", err)
	}

	return nil
}

func (f friendRequest) Delete(ctx *Context) error {
	if ctx.User() == nil {
		return errors.NotAuthorized
	}
	model, err := f.Model(ctx)
	if err != nil {
		return err
	}
	if model.From != ctx.User().ID() {
		return errors.RightsViolation
	}
	if err := f.app.repo.DeleteFriendRequest(ctx, model.ID); err != nil {
		return err
	}

	e := event.New(FriendRequestUpdateEventName, FriendRequestUpdateEvent{
		FriendRequestID: model.ID,
		From:            model.To,
		To:              model.From,
		Code:            FriendRequestUpdateDeleted,
	}, event.WithTime(time.Now()))

	if err := f.app.Events().Send(ctx, e); err != nil {
		// currently not handled
		log.Println("failed to send event:", err)
	}

	return nil
}

func unsafeFriendRequestFromModel(app *App, request models.FriendRequest) friendRequest {
	return friendRequest{
		id:  request.ID,
		app: app,
	}
}

func newFriendRequest(ctx *Context, app *App, id string) (FriendRequest, error) {
	req := friendRequest{
		id:  id,
		app: app,
	}

	if !req.exists(ctx) {
		return nil, errors.DoesNotExist
	}
	return req, nil
}
