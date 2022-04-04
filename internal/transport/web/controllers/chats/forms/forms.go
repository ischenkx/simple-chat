package forms

type GetChat struct {
	ID string `json:"id"`
}

type DeleteChat struct {
	ID string `json:"id"`
}

type CreateChat struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateChatMember struct {
	ChatID string `json:"chat_id"`
	UserID string `json:"user_id"`
}

type DeleteChatMember struct {
	ChatID string `json:"chat_id"`
	UserID string `json:"user_id"`
}

type UpdateMessage struct {
	ID      string `json:"id"`
	Payload string `json:"payload"`
}

type SendMessage struct {
	ChatID  string `json:"chat_id"`
	Payload string `json:"payload"`
}

type DeleteMessage struct {
	ID string `json:"id"`
}

type GetMessages struct {
	ChatID string `json:"id"`
	Offset int    `json:"offset"`
	Count  int    `json:"count"`
}

type GetChatMembers struct {
	ChatID string `json:"id"`
	Offset int    `json:"offset"`
	Count  int    `json:"count"`
}
