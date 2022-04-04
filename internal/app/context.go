package app

import (
	"context"
)

type Context struct {
	context.Context
	user User
}

func (ctx *Context) User() User {
	return ctx.user
}

func (ctx *Context) SetUser(user User) {
	ctx.user = user
}

func NewContext(ctx context.Context) *Context {
	return &Context{
		Context: ctx,
		user:    nil,
	}
}
