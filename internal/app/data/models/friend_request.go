package models

import "time"

type FriendRequest struct {
	ID   string
	From string
	To   string
	Time time.Time
}
