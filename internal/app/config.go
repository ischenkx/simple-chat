package app

import (
	"github.com/ischenkx/vk-test-task/internal/app/data"
	"github.com/ischenkx/vk-test-task/internal/app/event"
	"github.com/ischenkx/vk-test-task/internal/app/security"
)

type Config struct {
	Repo       data.Repository
	Authorizer security.Authorizer
	Bus        event.Bus
}
