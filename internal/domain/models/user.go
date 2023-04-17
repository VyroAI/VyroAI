package models

import "time"

type User struct {
	Id             int64     `json:"id"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	AvatarID       string    `json:"avatar_id"`
	Permission     int32     `json:"permission"`
	IsBanned       bool      `json:"is_banned"`
	EmailConfirmed bool      `json:"is_email_confirmed"`
	CreatedAt      time.Time `json:"created_at"`
}

type Subscription struct {
}
