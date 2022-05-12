package pkg

import "time"

type User struct {
	User_id int
	Token string
	Device string
	Settings string
	Created time.Time
	Updated time.Time
}

type UserId_Token struct {
	User_id int
	Token string
}