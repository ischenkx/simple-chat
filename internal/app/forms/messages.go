package forms

import "errors"

type MessageUpdate struct {
	Payload string
}

type SendMessage struct {
	Payload string
}

func (form *MessageUpdate) Validate() error {
	if len(form.Payload) == 0 {
		return errors.New("empty messages are not valid")
	}
	return nil
}

func (form *SendMessage) Validate() error {
	if len(form.Payload) == 0 {
		return errors.New("empty messages are not valid")
	}
	return nil
}
