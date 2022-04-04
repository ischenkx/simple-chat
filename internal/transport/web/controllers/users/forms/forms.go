package forms

type SendFriendRequest struct {
	To string `json:"to"`
}

type DeclineFriendRequest struct {
	ID string `json:"id"`
}

type AcceptFriendRequest struct {
	ID string `json:"id"`
}

type GetOutgoingFriendRequests struct {
	Offset int `json:"offset"`
	Count  int `json:"count"`
}

type GetIncomingFriendRequests struct {
	Offset int `json:"offset"`
	Count  int `json:"count"`
}

type GetFriends struct {
	Offset int `json:"offset"`
	Count  int `json:"count"`
}

type GetChats struct {
	Offset int `json:"offset"`
	Count  int `json:"count"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Register struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
