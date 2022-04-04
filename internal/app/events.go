package app

const NewMessageEventName = "new_message"
const MessageDeletedEventName = "message_deleted"
const MessageUpdatedEventName = "message_updated"
const ChatDeletedEventName = "chat_deleted"
const ChatMemberCreatedEventName = "chat_member_created"
const ChatMemberDeletedEventName = "chat_member_deleted"
const NewFriendRequestEventName = "friend_request"
const FriendRequestUpdateEventName = "friend_request_update"
const FriendAddedEventName = "friend_added"
const FriendDeletedEventName = "friend_deleted"

const FriendRequestUpdateAccepted = 1
const FriendRequestUpdateDeclined = 2
const FriendRequestUpdateDeleted = 3

type NewMessageEvent struct {
	MessageID string
	ChatID    string
}

type MessageDeletedEvent struct {
	MessageID string
	ChatID    string
}

type MessageUpdatedEvent struct {
	MessageID string
	ChatID    string
}

type NewFriendRequestEvent struct {
	FromID string
	ToID   string
	ID     string
}

type ChatDeletedEvent struct {
	ChatID string
}

type ChatMemberDeletedEvent struct {
	ChatID string
	UserID string
}

type ChatMemberCreatedEvent struct {
	ChatID string
	UserID string
}

type FriendRequestUpdateEvent struct {
	FriendRequestID string
	From            string
	To              string
	Code            int
}

type FriendAddedEvent struct {
	FriendID string
	UserID   string
}

type FriendDeletedEvent struct {
	FriendID string
	UserID   string
}
