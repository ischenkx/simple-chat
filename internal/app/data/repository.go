package data

import (
	"context"
	"github.com/ischenkx/vk-test-task/internal/app/data/models"
)

type Tx interface {
	CreateUser(ctx context.Context, user models.User) (models.User, error)
	CreateFriendRequest(ctx context.Context, request models.FriendRequest) (models.FriendRequest, error)
	CreateChat(ctx context.Context, chat models.Chat) (models.Chat, error)
	CreateChatMember(ctx context.Context, member models.ChatMember) (models.ChatMember, error)
	CreateMessage(ctx context.Context, model models.Message) (models.Message, error)
	CreateFriendConnection(ctx context.Context, id1, id2 string) error

	DeleteUser(ctx context.Context, id string) error
	DeleteFriendConnection(ctx context.Context, id1, id2 string) error
	DeleteFriendRequest(ctx context.Context, id string) error
	DeleteChat(ctx context.Context, id string) error
	DeleteChatMember(ctx context.Context, userId, chatId string) error
	DeleteMessage(ctx context.Context, id string) error

	UpdateUser(ctx context.Context, user models.User) error
	UpdateChat(ctx context.Context, user models.Chat) (models.Chat, error)
	UpdateChatMember(ctx context.Context, model models.ChatMember) (models.ChatMember, error)
	UpdateMessage(ctx context.Context, model models.Message) error

	GetUserChats(ctx context.Context, userId string, offset int, count int) ([]models.ChatMember, error)
	GetChatMembers(ctx context.Context, chatId string, offset int, count int) ([]models.ChatMember, error)
	GetChatMessages(ctx context.Context, chatId string, offset int, count int) ([]models.Message, error)
	GetUser(ctx context.Context, id string) (models.User, error)
	GetUserByUsername(ctx context.Context, id string) (models.User, error)
	GetUserFriends(ctx context.Context, id string, offset int, count int) ([]models.User, error)
	GetUserIncomingFriendRequests(ctx context.Context, id string, offset int, count int) ([]models.FriendRequest, error)
	GetUserOutgoingFriendRequests(ctx context.Context, id string, offset int, count int) ([]models.FriendRequest, error)
	GetUserIncomingFriendRequest(ctx context.Context, id, from string) (models.FriendRequest, error)
	GetFriendRequest(ctx context.Context, from, to string) (models.FriendRequest, error)
	GetFriendRequestByID(ctx context.Context, id string) (models.FriendRequest, error)
	GetChat(ctx context.Context, id string) (models.Chat, error)
	GetChatMember(ctx context.Context, userId, chatId string) (models.ChatMember, error)
	GetMessage(ctx context.Context, id string) (models.Message, error)
	FriendConnectionExists(ctx context.Context, id1, id2 string) bool

	CountFriends(ctx context.Context, id string) (int, error)
	CountUserIncomingFriendRequests(ctx context.Context, id string) (int, error)
	CountUserOutgoingFriendRequests(ctx context.Context, id string) (int, error)
	CountChatMembers(ctx context.Context, chatId string) (int, error)
	CountChatMessages(ctx context.Context, chatId string) (int, error)
	CountUserChats(ctx context.Context, id string) (int, error)
}

type Repository interface {
	Tx

	Transaction(context.Context, func(Tx) (interface{}, error)) (interface{}, error)
}
