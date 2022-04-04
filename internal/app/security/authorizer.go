package security

import (
	"context"
)

type Authorizer interface {
	Verify(ctx context.Context, token string) (userId string, err error)
	GenerateToken(ctx context.Context, userId string) (token string, err error)
}
