package app

import (
	goerrors "errors"
	"github.com/ischenkx/vk-test-task/internal/app/data/models"
	"github.com/ischenkx/vk-test-task/internal/app/errors"
	"github.com/ischenkx/vk-test-task/internal/app/forms"
	"golang.org/x/crypto/bcrypt"
)

type UserManager struct {
	app *App
}

func (manager UserManager) Register(ctx *Context, form forms.UserRegistration) (User, error) {
	if ctx.User() != nil {
		return nil, errors.AlreadyAuthorized
	}

	if err := form.Validate(); err != nil {
		return nil, err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(form.Password), 10)

	if err != nil {
		return nil, goerrors.New("failed to hash the password")
	}

	model := models.User{
		PasswordHash: passwordHash,
		Username:     form.Username,
	}

	res, err := manager.app.repo.CreateUser(ctx, model)

	if err != nil {
		return nil, err
	}

	return unsafeUserFromModel(manager.app, res), nil
}

func (manager UserManager) Login(ctx *Context, form forms.UserLogin) (User, error) {
	if ctx.User() != nil {
		return nil, errors.AlreadyAuthorized
	}

	if err := form.Validate(); err != nil {
		return nil, err
	}

	u, err := manager.app.repo.GetUserByUsername(ctx, form.Username)

	if err != nil {
		return nil, goerrors.New("failed to find such a user")
	}

	if err := bcrypt.CompareHashAndPassword(u.PasswordHash, []byte(form.Password)); err != nil {
		return nil, goerrors.New("invalid password")
	}

	return unsafeUserFromModel(manager.app, u), nil
}

func (manager UserManager) Get(ctx *Context, id string) (User, error) {
	return newUser(ctx, manager.app, id)
}
