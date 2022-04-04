package postgres

import (
	"github.com/ischenkx/vk-test-task/internal/app/data/models"
	"github.com/jackc/pgx/v4"
)

func parseUser(row pgx.Row) (models.User, error) {
	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.PasswordHash)
	return user, err
}

func parseInt(row pgx.Row) (int, error) {
	var num int
	err := row.Scan(&num)
	return num, err
}

func parseBool(row pgx.Row) (bool, error) {
	var res bool
	err := row.Scan(&res)
	return res, err
}

func parseFriendRequest(row pgx.Row) (models.FriendRequest, error) {
	var res models.FriendRequest
	err := row.Scan(&res.ID, &res.From, &res.To, &res.Time)
	return res, err
}

func parseChat(row pgx.Row) (models.Chat, error) {
	var res models.Chat
	err := row.Scan(&res.ID, &res.Name, &res.Description, &res.OwnerID)
	return res, err
}

func parseChatMember(row pgx.Row) (models.ChatMember, error) {
	var res models.ChatMember
	err := row.Scan(&res.UserID, &res.ChatID, &res.Status)
	return res, err
}

func parseMessage(row pgx.Row) (models.Message, error) {
	var res models.Message
	err := row.Scan(&res.ID, &res.UserID, &res.ChatID, &res.Payload, &res.TimeStamp, &res.LastUpdate)
	return res, err
}
