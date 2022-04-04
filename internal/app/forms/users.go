package forms

import (
	"errors"
)

type UserUpdate struct {
	Username string
}

type UserLogin struct {
	Username string
	Password string
}

type UserRegistration struct {
	Username string
	Password string
}

func (form UserUpdate) Validate() error {
	if len(form.Username) < 5 || len(form.Username) > 20 {
		return errors.New("invalid username length")
	}
	return nil
}

func (form UserLogin) Validate() error {
	if len(form.Username) < 5 || len(form.Username) > 20 {
		return errors.New("invalid username length")
	}

	if len(form.Password) < 5 || len(form.Password) > 20 {
		return errors.New("invalid password length")
	}

	return nil
}

func (form UserRegistration) Validate() error {
	if len(form.Username) < 5 || len(form.Username) > 20 {
		return errors.New("invalid username length")
	}

	if len(form.Password) < 5 || len(form.Password) > 20 {
		return errors.New("invalid password length")
	}

	return nil
}
