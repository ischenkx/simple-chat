package util

import (
	"context"
	"github.com/ischenkx/vk-test-task/internal/app"
)

func AppContext(ctx context.Context) (*app.Context, bool) {
	appCtx, ok := ctx.(*app.Context)
	return appCtx, ok
}
