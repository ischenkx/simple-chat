package web

import (
	"github.com/ischenkx/vk-test-task/internal/app"
	"github.com/ischenkx/vk-test-task/internal/transport/web/controllers/chats"
	"github.com/ischenkx/vk-test-task/internal/transport/web/controllers/common/middlewares"
	"github.com/ischenkx/vk-test-task/internal/transport/web/controllers/users"
	"net/http"
)

func NewRouter(a *app.App) http.Handler {
	mux := http.NewServeMux()

	// routes
	mux.Handle("/users/", http.StripPrefix("/users", users.NewController(a)))
	mux.Handle("/chats/", http.StripPrefix("/chats", chats.NewController(a)))

	// middlewares
	handler := middlewares.Auth(a, mux)
	handler = middlewares.ContextInitializer(handler)

	return handler
}
