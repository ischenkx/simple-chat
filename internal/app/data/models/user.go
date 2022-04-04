package models

type User struct {
	PasswordHash []byte
	Username     string
	ID           string
}
