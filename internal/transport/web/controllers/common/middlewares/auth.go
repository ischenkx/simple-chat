package middlewares

import (
	"github.com/ischenkx/vk-test-task/internal/app"
	"github.com/ischenkx/vk-test-task/internal/transport/web/controllers/common"
	"github.com/ischenkx/vk-test-task/internal/transport/web/controllers/common/auth"
	"github.com/ischenkx/vk-test-task/internal/transport/web/controllers/common/result"
	"github.com/ischenkx/vk-test-task/internal/transport/web/util"
	"net/http"
)

func Auth(application *app.App, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		appCtx, ok := util.AppContext(req.Context())

		if !ok {
			result.WriteSilent(w, result.New(nil, common.InternalServerErr))
			return
		}

		if token, err := auth.LoadVerificationToken(w, req); err == nil {
			userID, err := application.Auth().Verify(req.Context(), token)
			if err == nil {
				if user, err := application.Users().Get(appCtx, userID); err == nil {
					appCtx.SetUser(user)
				}
			}
		}

		next.ServeHTTP(w, req)
	})
}
