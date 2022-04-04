package middlewares

import (
	"github.com/ischenkx/vk-test-task/internal/app"
	"net/http"
)

func ContextInitializer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		req = req.WithContext(app.NewContext(req.Context()))
		next.ServeHTTP(w, req)
	})
}
