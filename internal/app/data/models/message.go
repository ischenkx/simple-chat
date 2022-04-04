package models

import "time"

type Message struct {
	Payload    string
	TimeStamp  time.Time
	LastUpdate time.Time
	ChatID     string
	UserID     string
	ID         string
}
