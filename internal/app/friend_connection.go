package app

import (
	"github.com/ischenkx/vk-test-task/internal/app/errors"
	"github.com/ischenkx/vk-test-task/internal/app/event"
	"log"
	"time"
)

type FriendConnection interface {
	User(ctx *Context) (User, error)
	Friend(ctx *Context) (User, error)
	Delete(ctx *Context) error
}

type friendConnection struct {
	app    *App
	user   string
	friend string
}

func (f friendConnection) isWritable(ctx *Context) bool {
	if ctx.User() == nil {
		return false
	}
	userID := ctx.User().ID()
	return userID == f.user || userID == f.friend
}

func (f friendConnection) exists(ctx *Context) bool {
	return f.app.repo.FriendConnectionExists(ctx, f.user, f.friend)
}

func (f friendConnection) User(ctx *Context) (User, error) {
	return newUser(ctx, f.app, f.user)
}

func (f friendConnection) Friend(ctx *Context) (User, error) {
	return newUser(ctx, f.app, f.friend)
}

func (f friendConnection) Delete(ctx *Context) error {
	if !f.isWritable(ctx) {
		return errors.ResourceInaccessible
	}
	if err := f.app.repo.DeleteFriendConnection(ctx, f.user, f.friend); err != nil {
		return err
	}

	e := event.New(FriendDeletedEventName, FriendDeletedEvent{
		FriendID: f.friend,
		UserID:   f.user,
	}, event.WithTime(time.Now()))

	if err := f.app.Events().Send(ctx, e); err != nil {
		// currently not handled
		log.Println("failed to send event:", err)
	}

	return nil
}

func unsafeFriendConnection(app *App, src, friend string) friendConnection {
	return friendConnection{
		app:    app,
		user:   src,
		friend: friend,
	}
}

func newFriendConnection(ctx *Context, app *App, src, friend string) (FriendConnection, error) {
	u := unsafeFriendConnection(app, src, friend)
	if !u.exists(ctx) {
		return nil, errors.DoesNotExist
	}
	return u, nil
}
