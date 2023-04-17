package models

import "time"

type User struct {
	Id             int64     `json:"id,omitempty"`
	Username       string    `json:"username,omitempty"`
	Email          string    `json:"email,omitempty"`
	Password       string    `json:"password,omitempty"`
	AvatarID       string    `json:"avatar_id,omitempty"`
	Permission     int32     `json:"permission"`
	IsBanned       bool      `json:"is_banned"`
	EmailConfirmed bool      `json:"is_email_confirmed"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
}

type Subscription struct {
}
