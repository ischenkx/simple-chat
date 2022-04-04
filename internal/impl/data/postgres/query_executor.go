package postgres

import (
	"context"
	"github.com/ischenkx/vk-test-task/internal/app/data"
	"github.com/ischenkx/vk-test-task/internal/app/data/models"
)

type QueryExecutor struct {
	pg PostgresInterface
}

func (r QueryExecutor) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	row := r.pg.QueryRow(ctx, createUserSql, user.Username, user.PasswordHash)
	return parseUser(row)
}

func (r QueryExecutor) DeleteUser(ctx context.Context, id string) error {
	_, err := r.pg.Exec(ctx, deleteUserSql, id)
	return err
}

func (r QueryExecutor) GetUser(ctx context.Context, id string) (models.User, error) {
	res := r.pg.QueryRow(ctx, getUserSql, id)
	return parseUser(res)
}

func (r QueryExecutor) GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	res := r.pg.QueryRow(ctx, getUserByUsernameSql, username)
	return parseUser(res)
}

func (r QueryExecutor) UpdateUser(ctx context.Context, user models.User) error {
	_, err := r.pg.Exec(ctx, updateUserSql, user.ID, user.Username, user.PasswordHash)
	return err
}

func (r QueryExecutor) CreateFriendConnection(ctx context.Context, id1, id2 string) error {
	_, err := r.pg.Exec(ctx, createFriendConnectionSql, id1, id2)
	return err
}

func (r QueryExecutor) DeleteFriendConnection(ctx context.Context, id1, id2 string) error {
	_, err := r.pg.Exec(ctx, deleteFriendConnectionSql, id1, id2)
	return err
}

func (r QueryExecutor) GetUserFriends(ctx context.Context, id string, offset int, count int) ([]models.User, error) {
	query, err := r.pg.Query(ctx, getUserFriendsSql, id, offset, count)
	if err != nil {
		return nil, err
	}

	var res []models.User
	for query.Next() {
		if query.Err() != nil {
			return nil, query.Err()
		}
		model, err := parseUser(query)
		if err != nil {
			return nil, err
		}
		res = append(res, model)
	}

	return res, err
}

func (r QueryExecutor) CountFriends(ctx context.Context, id string) (int, error) {
	row := r.pg.QueryRow(ctx, countUserFriendsSql, id)
	return parseInt(row)
}

// TODO: add error

func (r QueryExecutor) FriendConnectionExists(ctx context.Context, id1, id2 string) bool {
	row := r.pg.QueryRow(ctx, friendConnectionCheckerSql, id1, id2)
	res, err := parseBool(row)
	if err != nil {
		return false
	}
	return res
}

func (r QueryExecutor) CreateFriendRequest(ctx context.Context, request models.FriendRequest) (models.FriendRequest, error) {
	row := r.pg.QueryRow(ctx, createFriendRequestSql, request.From, request.To, request.Time)
	return parseFriendRequest(row)
}

func (r QueryExecutor) DeleteFriendRequest(ctx context.Context, id string) error {
	_, err := r.pg.Exec(ctx, deleteFriendRequestSql, id)
	return err
}

func (r QueryExecutor) GetUserIncomingFriendRequests(ctx context.Context, id string, offset int, count int) ([]models.FriendRequest, error) {
	query, err := r.pg.Query(ctx, getUserIncomingFriendRequestsSql, id, offset, count)
	if err != nil {
		return nil, err
	}

	var res []models.FriendRequest
	for query.Next() {
		if query.Err() != nil {
			return nil, query.Err()
		}
		model, err := parseFriendRequest(query)
		if err != nil {
			return nil, err
		}
		res = append(res, model)
	}

	return res, err
}

func (r QueryExecutor) GetUserOutgoingFriendRequests(ctx context.Context, id string, offset int, count int) ([]models.FriendRequest, error) {
	query, err := r.pg.Query(ctx, getUserOutgoingFriendRequestsSql, id, offset, count)
	if err != nil {
		return nil, err
	}

	var res []models.FriendRequest
	for query.Next() {
		if query.Err() != nil {
			return nil, query.Err()
		}
		model, err := parseFriendRequest(query)
		if err != nil {
			return nil, err
		}
		res = append(res, model)
	}

	return res, err
}

func (r QueryExecutor) GetUserIncomingFriendRequest(ctx context.Context, id, from string) (models.FriendRequest, error) {
	res := r.pg.QueryRow(ctx, getUserIncomingFriendRequestSql, id, from)
	return parseFriendRequest(res)
}

func (r QueryExecutor) CountUserIncomingFriendRequests(ctx context.Context, id string) (int, error) {
	row := r.pg.QueryRow(ctx, countUserIncomingFriendRequestsSql, id)
	return parseInt(row)
}

func (r QueryExecutor) CountUserOutgoingFriendRequests(ctx context.Context, id string) (int, error) {
	row := r.pg.QueryRow(ctx, countUserOutgoingFriendRequestsSql, id)
	return parseInt(row)
}

func (r QueryExecutor) GetFriendRequest(ctx context.Context, from, to string) (models.FriendRequest, error) {
	res := r.pg.QueryRow(ctx, getFriendRequestSql, from, to)
	return parseFriendRequest(res)
}

func (r QueryExecutor) GetFriendRequestByID(ctx context.Context, id string) (models.FriendRequest, error) {
	res := r.pg.QueryRow(ctx, getFriendRequestByIDSql, id)
	return parseFriendRequest(res)
}

func (r QueryExecutor) CreateChat(ctx context.Context, chat models.Chat) (models.Chat, error) {
	row := r.pg.QueryRow(ctx, createChatSql, chat.Name, chat.Description, chat.OwnerID)
	return parseChat(row)
}

func (r QueryExecutor) DeleteChat(ctx context.Context, id string) error {
	_, err := r.pg.Exec(ctx, deleteChatSql, id)
	return err
}

func (r QueryExecutor) GetChat(ctx context.Context, id string) (models.Chat, error) {
	row := r.pg.QueryRow(ctx, getChatSql, id)
	return parseChat(row)
}

func (r QueryExecutor) UpdateChat(ctx context.Context, chat models.Chat) (models.Chat, error) {
	row := r.pg.QueryRow(ctx, updateChatSql, chat.ID, chat.Name, chat.Description)
	return parseChat(row)
}

func (r QueryExecutor) CreateChatMember(ctx context.Context, member models.ChatMember) (models.ChatMember, error) {
	row := r.pg.QueryRow(ctx, createChatMemberSql, member.UserID, member.ChatID, member.Status)
	return parseChatMember(row)
}

func (r QueryExecutor) DeleteChatMember(ctx context.Context, userId, chatId string) error {
	_, err := r.pg.Exec(ctx, deleteChatMemberSql, userId, chatId)
	return err
}

func (r QueryExecutor) UpdateChatMember(ctx context.Context, model models.ChatMember) (models.ChatMember, error) {
	row := r.pg.QueryRow(ctx, updateChatMemberSql, model.UserID, model.ChatID, model.Status)
	return parseChatMember(row)
}

func (r QueryExecutor) GetChatMember(ctx context.Context, userId, chatId string) (models.ChatMember, error) {
	row := r.pg.QueryRow(ctx, getChatMemberSql, userId, chatId)
	return parseChatMember(row)
}

func (r QueryExecutor) CreateMessage(ctx context.Context, model models.Message) (models.Message, error) {
	row := r.pg.QueryRow(ctx, createMessageSql, model.UserID, model.ChatID, model.Payload, model.TimeStamp, model.TimeStamp)
	return parseMessage(row)
}

func (r QueryExecutor) DeleteMessage(ctx context.Context, id string) error {
	_, err := r.pg.Exec(ctx, deleteMessageSql, id)
	return err
}

func (r QueryExecutor) UpdateMessage(ctx context.Context, model models.Message) error {
	row := r.pg.QueryRow(ctx, updateMessageSql, model.ID, model.Payload, model.LastUpdate)
	_, err := parseMessage(row)
	return err
}

func (r QueryExecutor) GetMessage(ctx context.Context, id string) (models.Message, error) {
	row := r.pg.QueryRow(ctx, getMessageSql, id)
	return parseMessage(row)
}

func (r QueryExecutor) GetUserChats(ctx context.Context, userId string, offset int, count int) ([]models.ChatMember, error) {
	query, err := r.pg.Query(ctx, getUserChatsSql, userId, offset, count)
	if err != nil {
		return nil, err
	}

	var res []models.ChatMember
	for query.Next() {
		if query.Err() != nil {
			return nil, query.Err()
		}
		model, err := parseChatMember(query)
		if err != nil {
			return nil, err
		}
		res = append(res, model)
	}

	return res, err
}

func (r QueryExecutor) GetChatMembers(ctx context.Context, chatId string, offset int, count int) ([]models.ChatMember, error) {
	query, err := r.pg.Query(ctx, getChatMembersSql, chatId, offset, count)
	if err != nil {
		return nil, err
	}

	var res []models.ChatMember
	for query.Next() {
		if query.Err() != nil {
			return nil, query.Err()
		}
		model, err := parseChatMember(query)
		if err != nil {
			return nil, err
		}
		res = append(res, model)
	}

	return res, err
}

func (r QueryExecutor) GetChatMessages(ctx context.Context, chatId string, offset int, count int) ([]models.Message, error) {
	query, err := r.pg.Query(ctx, getChatMessagesSql, chatId, offset, count)
	if err != nil {
		return nil, err
	}

	var res []models.Message
	for query.Next() {
		if query.Err() != nil {
			return nil, query.Err()
		}
		model, err := parseMessage(query)
		if err != nil {
			return nil, err
		}
		res = append(res, model)
	}

	return res, err
}

func (r QueryExecutor) CountChatMembers(ctx context.Context, chatId string) (int, error) {
	row := r.pg.QueryRow(ctx, countChatMembersSql, chatId)
	return parseInt(row)
}

func (r QueryExecutor) CountChatMessages(ctx context.Context, chatId string) (int, error) {
	row := r.pg.QueryRow(ctx, countChatMessagesSql, chatId)
	return parseInt(row)
}

func (r QueryExecutor) CountUserChats(ctx context.Context, id string) (int, error) {
	row := r.pg.QueryRow(ctx, countUserChatsSql, id)
	return parseInt(row)
}

func queryExecutor(pg PostgresInterface) data.Tx {
	return QueryExecutor{
		pg: pg,
	}
}
