package postgres

import (
	"context"
	"github.com/ischenkx/vk-test-task/internal/app/data/models"
	"github.com/jackc/pgx/v4"
)

type Tx struct {
	pg pgx.Tx
}

func (t Tx) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	return queryExecutor(t.pg).CreateUser(ctx, user)
}

func (t Tx) DeleteUser(ctx context.Context, id string) error {
	return queryExecutor(t.pg).DeleteUser(ctx, id)
}

func (t Tx) GetUser(ctx context.Context, id string) (models.User, error) {
	return queryExecutor(t.pg).GetUser(ctx, id)
}

func (t Tx) GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	return queryExecutor(t.pg).GetUserByUsername(ctx, username)
}

func (t Tx) UpdateUser(ctx context.Context, user models.User) error {
	return queryExecutor(t.pg).UpdateUser(ctx, user)
}

func (t Tx) CreateFriendConnection(ctx context.Context, id1, id2 string) error {
	return queryExecutor(t.pg).CreateFriendConnection(ctx, id1, id2)
}

func (t Tx) DeleteFriendConnection(ctx context.Context, id1, id2 string) error {
	return queryExecutor(t.pg).DeleteFriendConnection(ctx, id1, id2)
}

func (t Tx) GetUserFriends(ctx context.Context, id string, offset int, count int) ([]models.User, error) {
	return queryExecutor(t.pg).GetUserFriends(ctx, id, offset, count)

}

func (t Tx) CountFriends(ctx context.Context, id string) (int, error) {
	return queryExecutor(t.pg).CountFriends(ctx, id)

}

func (t Tx) FriendConnectionExists(ctx context.Context, id1, id2 string) bool {
	return queryExecutor(t.pg).FriendConnectionExists(ctx, id1, id2)

}

func (t Tx) CreateFriendRequest(ctx context.Context, request models.FriendRequest) (models.FriendRequest, error) {
	return queryExecutor(t.pg).CreateFriendRequest(ctx, request)

}

func (t Tx) DeleteFriendRequest(ctx context.Context, id string) error {
	return queryExecutor(t.pg).DeleteFriendRequest(ctx, id)

}

func (t Tx) GetUserIncomingFriendRequests(ctx context.Context, id string, offset int, count int) ([]models.FriendRequest, error) {
	return queryExecutor(t.pg).GetUserIncomingFriendRequests(ctx, id, offset, count)

}

func (t Tx) GetUserOutgoingFriendRequests(ctx context.Context, id string, offset int, count int) ([]models.FriendRequest, error) {
	return queryExecutor(t.pg).GetUserOutgoingFriendRequests(ctx, id, offset, count)

}

func (t Tx) GetUserIncomingFriendRequest(ctx context.Context, id, from string) (models.FriendRequest, error) {
	return queryExecutor(t.pg).GetUserIncomingFriendRequest(ctx, id, from)

}

func (t Tx) CountUserIncomingFriendRequests(ctx context.Context, id string) (int, error) {
	return queryExecutor(t.pg).CountUserIncomingFriendRequests(ctx, id)

}

func (t Tx) CountUserOutgoingFriendRequests(ctx context.Context, id string) (int, error) {
	return queryExecutor(t.pg).CountUserOutgoingFriendRequests(ctx, id)

}

func (t Tx) GetFriendRequest(ctx context.Context, from, to string) (models.FriendRequest, error) {
	return queryExecutor(t.pg).GetFriendRequest(ctx, from, to)

}

func (t Tx) GetFriendRequestByID(ctx context.Context, id string) (models.FriendRequest, error) {
	return queryExecutor(t.pg).GetFriendRequestByID(ctx, id)

}

func (t Tx) CreateChat(ctx context.Context, chat models.Chat) (models.Chat, error) {
	return queryExecutor(t.pg).CreateChat(ctx, chat)

}

func (t Tx) DeleteChat(ctx context.Context, id string) error {
	return queryExecutor(t.pg).DeleteChat(ctx, id)

}

func (t Tx) GetChat(ctx context.Context, id string) (models.Chat, error) {
	return queryExecutor(t.pg).GetChat(ctx, id)
}

func (t Tx) UpdateChat(ctx context.Context, chat models.Chat) (models.Chat, error) {
	return queryExecutor(t.pg).UpdateChat(ctx, chat)

}

func (t Tx) CreateChatMember(ctx context.Context, member models.ChatMember) (models.ChatMember, error) {
	return queryExecutor(t.pg).CreateChatMember(ctx, member)

}

func (t Tx) DeleteChatMember(ctx context.Context, userId, chatId string) error {
	return queryExecutor(t.pg).DeleteChatMember(ctx, userId, chatId)

}

func (t Tx) UpdateChatMember(ctx context.Context, model models.ChatMember) (models.ChatMember, error) {
	return queryExecutor(t.pg).UpdateChatMember(ctx, model)

}

func (t Tx) GetChatMember(ctx context.Context, userId, chatId string) (models.ChatMember, error) {
	return queryExecutor(t.pg).GetChatMember(ctx, userId, chatId)

}

func (t Tx) GetMessage(ctx context.Context, id string) (models.Message, error) {
	return queryExecutor(t.pg).GetMessage(ctx, id)

}

func (t Tx) CreateMessage(ctx context.Context, model models.Message) (models.Message, error) {
	return queryExecutor(t.pg).CreateMessage(ctx, model)

}

func (t Tx) DeleteMessage(ctx context.Context, id string) error {
	return queryExecutor(t.pg).DeleteMessage(ctx, id)

}

func (t Tx) UpdateMessage(ctx context.Context, model models.Message) error {
	return queryExecutor(t.pg).UpdateMessage(ctx, model)

}

func (t Tx) GetUserChats(ctx context.Context, userId string, offset int, count int) ([]models.ChatMember, error) {
	return queryExecutor(t.pg).GetUserChats(ctx, userId, offset, count)

}

func (t Tx) GetChatMembers(ctx context.Context, chatId string, offset int, count int) ([]models.ChatMember, error) {
	return queryExecutor(t.pg).GetChatMembers(ctx, chatId, offset, count)

}

func (t Tx) GetChatMessages(ctx context.Context, chatId string, offset int, count int) ([]models.Message, error) {
	return queryExecutor(t.pg).GetChatMessages(ctx, chatId, offset, count)

}

func (t Tx) CountChatMembers(ctx context.Context, chatId string) (int, error) {
	return queryExecutor(t.pg).CountChatMembers(ctx, chatId)

}

func (t Tx) CountChatMessages(ctx context.Context, chatId string) (int, error) {
	return queryExecutor(t.pg).CountChatMembers(ctx, chatId)

}

func (t Tx) CountUserChats(ctx context.Context, id string) (int, error) {
	return queryExecutor(t.pg).CountUserChats(ctx, id)

}
