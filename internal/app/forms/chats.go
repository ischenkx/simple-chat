package forms

import "errors"

type ChatCreationForm struct {
	Name        string
	Description string
}

func (form *ChatCreationForm) Validate() error {
	if len(form.Name) < 5 || len(form.Name) > 230 {
		return errors.New("invalid name length")
	}

	if len(form.Description) > 400 {
		return errors.New("invalid description length")
	}

	return nil
}
