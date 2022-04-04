package postgres

import (
	"context"
	"github.com/ischenkx/vk-test-task/internal/app/data"
	"github.com/ischenkx/vk-test-task/internal/app/data/models"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repo struct {
	pg *pgxpool.Pool
}

func NewRepo(conn *pgxpool.Pool) *Repo {
	return &Repo{conn}
}

func (r *Repo) Transaction(ctx context.Context, f func(data.Tx) (interface{}, error)) (interface{}, error) {
	var output interface{}
	err := r.pg.BeginFunc(ctx, func(tx pgx.Tx) error {
		res, err := f(queryExecutor(tx))
		output = res
		return err
	})
	return output, err
}

func (r *Repo) InitializeTables(ctx context.Context) error {
	_, err := r.pg.Exec(ctx, initializeTablesSql)
	return err
}

func (r *Repo) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	return queryExecutor(r.pg).CreateUser(ctx, user)
}

func (r *Repo) DeleteUser(ctx context.Context, id string) error {
	return queryExecutor(r.pg).DeleteUser(ctx, id)
}

func (r *Repo) GetUser(ctx context.Context, id string) (models.User, error) {
	return queryExecutor(r.pg).GetUser(ctx, id)
}

func (r *Repo) GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	return queryExecutor(r.pg).GetUserByUsername(ctx, username)
}

func (r *Repo) UpdateUser(ctx context.Context, user models.User) error {
	return queryExecutor(r.pg).UpdateUser(ctx, user)
}

func (r *Repo) CreateFriendConnection(ctx context.Context, id1, id2 string) error {
	return queryExecutor(r.pg).CreateFriendConnection(ctx, id1, id2)
}

func (r *Repo) DeleteFriendConnection(ctx context.Context, id1, id2 string) error {
	return queryExecutor(r.pg).DeleteFriendConnection(ctx, id1, id2)
}

func (r *Repo) GetUserFriends(ctx context.Context, id string, offset int, count int) ([]models.User, error) {
	return queryExecutor(r.pg).GetUserFriends(ctx, id, offset, count)

}

func (r *Repo) CountFriends(ctx context.Context, id string) (int, error) {
	return queryExecutor(r.pg).CountFriends(ctx, id)

}

func (r *Repo) FriendConnectionExists(ctx context.Context, id1, id2 string) bool {
	return queryExecutor(r.pg).FriendConnectionExists(ctx, id1, id2)

}

func (r *Repo) CreateFriendRequest(ctx context.Context, request models.FriendRequest) (models.FriendRequest, error) {
	return queryExecutor(r.pg).CreateFriendRequest(ctx, request)

}

func (r *Repo) DeleteFriendRequest(ctx context.Context, id string) error {
	return queryExecutor(r.pg).DeleteFriendRequest(ctx, id)

}

func (r *Repo) GetUserIncomingFriendRequests(ctx context.Context, id string, offset int, count int) ([]models.FriendRequest, error) {
	return queryExecutor(r.pg).GetUserIncomingFriendRequests(ctx, id, offset, count)

}

func (r *Repo) GetUserOutgoingFriendRequests(ctx context.Context, id string, offset int, count int) ([]models.FriendRequest, error) {
	return queryExecutor(r.pg).GetUserOutgoingFriendRequests(ctx, id, offset, count)

}

func (r *Repo) GetUserIncomingFriendRequest(ctx context.Context, id, from string) (models.FriendRequest, error) {
	return queryExecutor(r.pg).GetUserIncomingFriendRequest(ctx, id, from)

}

func (r *Repo) CountUserIncomingFriendRequests(ctx context.Context, id string) (int, error) {
	return queryExecutor(r.pg).CountUserIncomingFriendRequests(ctx, id)

}

func (r *Repo) CountUserOutgoingFriendRequests(ctx context.Context, id string) (int, error) {
	return queryExecutor(r.pg).CountUserOutgoingFriendRequests(ctx, id)

}

func (r *Repo) GetFriendRequest(ctx context.Context, from, to string) (models.FriendRequest, error) {
	return queryExecutor(r.pg).GetFriendRequest(ctx, from, to)

}

func (r *Repo) GetFriendRequestByID(ctx context.Context, id string) (models.FriendRequest, error) {
	return queryExecutor(r.pg).GetFriendRequestByID(ctx, id)

}

func (r *Repo) CreateChat(ctx context.Context, chat models.Chat) (models.Chat, error) {
	return queryExecutor(r.pg).CreateChat(ctx, chat)

}

func (r *Repo) DeleteChat(ctx context.Context, id string) error {
	return queryExecutor(r.pg).DeleteChat(ctx, id)

}

func (r *Repo) GetChat(ctx context.Context, id string) (models.Chat, error) {
	return queryExecutor(r.pg).GetChat(ctx, id)
}

func (r *Repo) UpdateChat(ctx context.Context, chat models.Chat) (models.Chat, error) {
	return queryExecutor(r.pg).UpdateChat(ctx, chat)

}

func (r *Repo) CreateChatMember(ctx context.Context, member models.ChatMember) (models.ChatMember, error) {
	return queryExecutor(r.pg).CreateChatMember(ctx, member)

}

func (r *Repo) DeleteChatMember(ctx context.Context, userId, chatId string) error {
	return queryExecutor(r.pg).DeleteChatMember(ctx, userId, chatId)

}

func (r *Repo) UpdateChatMember(ctx context.Context, model models.ChatMember) (models.ChatMember, error) {
	return queryExecutor(r.pg).UpdateChatMember(ctx, model)

}

func (r *Repo) GetChatMember(ctx context.Context, userId, chatId string) (models.ChatMember, error) {
	return queryExecutor(r.pg).GetChatMember(ctx, userId, chatId)

}

func (r *Repo) GetMessage(ctx context.Context, id string) (models.Message, error) {
	return queryExecutor(r.pg).GetMessage(ctx, id)

}

func (r *Repo) CreateMessage(ctx context.Context, model models.Message) (models.Message, error) {
	return queryExecutor(r.pg).CreateMessage(ctx, model)

}

func (r *Repo) DeleteMessage(ctx context.Context, id string) error {
	return queryExecutor(r.pg).DeleteMessage(ctx, id)

}

func (r *Repo) UpdateMessage(ctx context.Context, model models.Message) error {
	return queryExecutor(r.pg).UpdateMessage(ctx, model)

}

func (r *Repo) GetUserChats(ctx context.Context, userId string, offset int, count int) ([]models.ChatMember, error) {
	return queryExecutor(r.pg).GetUserChats(ctx, userId, offset, count)

}

func (r *Repo) GetChatMembers(ctx context.Context, chatId string, offset int, count int) ([]models.ChatMember, error) {
	return queryExecutor(r.pg).GetChatMembers(ctx, chatId, offset, count)

}

func (r *Repo) GetChatMessages(ctx context.Context, chatId string, offset int, count int) ([]models.Message, error) {
	return queryExecutor(r.pg).GetChatMessages(ctx, chatId, offset, count)

}

func (r *Repo) CountChatMembers(ctx context.Context, chatId string) (int, error) {
	return queryExecutor(r.pg).CountChatMembers(ctx, chatId)

}

func (r *Repo) CountChatMessages(ctx context.Context, chatId string) (int, error) {
	return queryExecutor(r.pg).CountChatMembers(ctx, chatId)

}

func (r *Repo) CountUserChats(ctx context.Context, id string) (int, error) {
	return queryExecutor(r.pg).CountUserChats(ctx, id)
}
