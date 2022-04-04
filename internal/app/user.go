package app

import (
	goerrors "errors"
	"github.com/ischenkx/vk-test-task/internal/app/data"
	"github.com/ischenkx/vk-test-task/internal/app/data/models"
	"github.com/ischenkx/vk-test-task/internal/app/errors"
	"github.com/ischenkx/vk-test-task/internal/app/event"
	"github.com/ischenkx/vk-test-task/internal/app/forms"
	"log"
	"time"
)

type User interface {
	ID() string
	Username(ctx *Context) (string, error)
	Update(ctx *Context, update forms.UserUpdate) error
	Chats(ctx *Context, offset int, count int) ([]ChatMember, error)
	Friends(ctx *Context, offset int, count int) ([]FriendConnection, error)
	Friend(ctx *Context, id string) (FriendConnection, error)

	IncomingFriendRequests(ctx *Context, offset int, count int) ([]FriendRequest, error)
	OutgoingFriendRequests(ctx *Context, offset int, count int) ([]FriendRequest, error)
	IncomingFriendRequest(ctx *Context, from string) (FriendRequest, error)
	OutgoingFriendRequest(ctx *Context, to string) (FriendRequest, error)
	SendFriendRequest(ctx *Context, to string) (FriendRequest, error)

	CountIncomingFriendRequests(ctx *Context) (int, error)
	CountOutgoingFriendRequests(ctx *Context) (int, error)
	CountChats(ctx *Context) (int, error)
	CountFriends(ctx *Context) (int, error)

	Delete(ctx *Context) error

	Model(ctx *Context) (models.User, error)
}

type user struct {
	app    *App
	userID string
}

func (u user) isWritable(ctx *Context) bool {
	if ctx.User() == nil {
		return false
	}
	return ctx.User().ID() == u.userID
}

func (u user) Model(ctx *Context) (models.User, error) {
	return u.app.repo.GetUser(ctx, u.userID)
}

func (u user) exists(ctx *Context) bool {
	if _, err := u.Model(ctx); err != nil {
		return false
	}
	return true
}

func (u user) Username(ctx *Context) (string, error) {
	if m, err := u.Model(ctx); err != nil {
		return "", err
	} else {
		return m.Username, nil
	}
}

func (u user) ID() string {
	return u.userID
}

func (u user) Update(ctx *Context, update forms.UserUpdate) error {
	if !u.isWritable(ctx) {
		return errors.ResourceInaccessible
	}

	if err := update.Validate(); err != nil {
		return err
	}

	m, err := u.Model(ctx)
	if err != nil {
		return err
	}

	m.Username = update.Username
	return u.app.repo.UpdateUser(ctx, m)
}

func (u user) Chats(ctx *Context, offset int, count int) ([]ChatMember, error) {
	repoChats, err := u.app.repo.GetUserChats(ctx, u.userID, offset, count)
	if err != nil {
		return nil, err
	}
	chats := make([]ChatMember, 0, len(repoChats))

	for i := 0; i < len(repoChats); i++ {
		cm := repoChats[i]
		chats = append(chats, unsafeChatMemberFromModel(u.app, cm))
	}

	return chats, err
}

func (u user) CountChats(ctx *Context) (int, error) {
	return u.app.repo.CountUserChats(ctx, u.userID)
}

func (u user) Delete(ctx *Context) error {
	if !u.isWritable(ctx) {
		return errors.ResourceInaccessible
	}
	return u.app.repo.DeleteUser(ctx, u.userID)
}

func (u user) SendFriendRequest(ctx *Context, to string) (FriendRequest, error) {
	if ctx.User() == nil {
		return nil, errors.NotAuthorized
	}

	if !u.isWritable(ctx) {
		return nil, errors.RightsViolation
	}

	if to == ctx.User().ID() {
		return nil, goerrors.New("sending friend requests to yourself is wierd")
	}

	req, err := u.app.repo.Transaction(ctx, func(repo data.Tx) (interface{}, error) {
		_, err := repo.GetFriendRequest(ctx, ctx.User().ID(), to)
		if err == nil {
			return nil, goerrors.New("already exists")
		}

		_, err = repo.GetFriendRequest(ctx, to, ctx.User().ID())
		if err == nil {
			return nil, goerrors.New("inverse exists")
		}

		if repo.FriendConnectionExists(ctx, ctx.User().ID(), to) {
			return nil, goerrors.New("already friends")
		}

		return repo.CreateFriendRequest(ctx, models.FriendRequest{
			From: ctx.User().ID(),
			To:   to,
			Time: time.Now(),
		})
	})

	if err != nil {
		return nil, err
	}

	model := req.(models.FriendRequest)

	e := event.New(NewFriendRequestEventName, NewFriendRequestEvent{
		FromID: u.userID,
		ToID:   to,
		ID:     model.ID,
	}, event.WithTime(time.Now()))

	if err := u.app.Events().Send(ctx, e); err != nil {
		// currently not handled
		log.Println("failed to send event:", err)
	}

	return unsafeFriendRequestFromModel(u.app, model), nil
}

func (u user) IncomingFriendRequests(ctx *Context, offset int, count int) ([]FriendRequest, error) {
	rawRequests, err := u.app.repo.GetUserIncomingFriendRequests(ctx, u.userID, offset, count)
	if err != nil {
		return nil, err
	}

	requests := make([]FriendRequest, 0, len(rawRequests))
	for _, model := range rawRequests {
		requests = append(requests, unsafeFriendRequestFromModel(u.app, model))
	}
	return requests, nil
}

func (u user) OutgoingFriendRequests(ctx *Context, offset int, count int) ([]FriendRequest, error) {
	rawRequests, err := u.app.repo.GetUserOutgoingFriendRequests(ctx, u.userID, offset, count)
	if err != nil {
		return nil, err
	}

	requests := make([]FriendRequest, 0, len(rawRequests))
	for _, model := range rawRequests {
		requests = append(requests, unsafeFriendRequestFromModel(u.app, model))
	}
	return requests, nil
}

func (u user) IncomingFriendRequest(ctx *Context, from string) (FriendRequest, error) {
	rawRequest, err := u.app.repo.GetUserIncomingFriendRequest(ctx, u.userID, from)
	if err != nil {
		return nil, err
	}
	return unsafeFriendRequestFromModel(u.app, rawRequest), nil
}

func (u user) OutgoingFriendRequest(ctx *Context, to string) (FriendRequest, error) {
	rawRequest, err := u.app.repo.GetUserIncomingFriendRequest(ctx, to, u.userID)
	if err != nil {
		return nil, err
	}
	return unsafeFriendRequestFromModel(u.app, rawRequest), nil
}

func (u user) CountIncomingFriendRequests(ctx *Context) (int, error) {
	return u.app.repo.CountUserIncomingFriendRequests(ctx, u.userID)
}

func (u user) CountOutgoingFriendRequests(ctx *Context) (int, error) {
	return u.app.repo.CountUserOutgoingFriendRequests(ctx, u.userID)
}

func (u user) Friends(ctx *Context, offset int, count int) ([]FriendConnection, error) {
	repoFriends, err := u.app.repo.GetUserFriends(ctx, u.userID, offset, count)
	if err != nil {
		return nil, err
	}
	chats := make([]FriendConnection, 0, len(repoFriends))

	for i := 0; i < len(repoFriends); i++ {
		friend := repoFriends[i]
		chats = append(chats, unsafeFriendConnection(u.app, u.userID, friend.ID))
	}

	return chats, err
}

func (u user) Friend(ctx *Context, id string) (FriendConnection, error) {
	return newFriendConnection(ctx, u.app, u.userID, id)
}

func (u user) CountFriends(ctx *Context) (int, error) {
	return u.app.repo.CountFriends(ctx, u.userID)
}

func unsafeUserFromModel(app *App, u models.User) User {
	return user{
		app:    app,
		userID: u.ID,
	}
}

func newUser(ctx *Context, app *App, id string) (User, error) {
	u := user{
		app:    app,
		userID: id,
	}
	if !u.exists(ctx) {
		return nil, errors.DoesNotExist
	}
	return u, nil
}
