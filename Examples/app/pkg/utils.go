package pkg

import "time"

type User struct {
	User_id uint32
	Token string
	Device string
	Settings string
	Created time.Time
	Updated time.Time
}

type UserId_Token struct {
	User_id uint32
	Token string
}

type TokensWithMessage struct {
	Token []string
	Message string
}

func (u *User) Init (user_id uint32, token string)  {
	u.User_id = user_id
	u.Token = token
	u.Device = "android"
	u.Settings = "1,1,1,1,1,1,1"
}