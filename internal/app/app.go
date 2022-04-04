package app

import (
	"github.com/ischenkx/vk-test-task/internal/app/data"
	"github.com/ischenkx/vk-test-task/internal/app/event"
	"github.com/ischenkx/vk-test-task/internal/app/security"
)

type App struct {
	repo       data.Repository
	authorizer security.Authorizer
	events     event.Bus
}

func (app *App) Events() event.Bus {
	return app.events
}

func (app *App) Auth() security.Authorizer {
	return app.authorizer
}

func (app *App) Users() UserManager {
	return UserManager{app}
}

func (app *App) Chats() ChatManager {
	return ChatManager{app}
}

func New(cfg Config) *App {
	return &App{
		repo:       cfg.Repo,
		authorizer: cfg.Authorizer,
		events:     cfg.Bus,
	}
}
