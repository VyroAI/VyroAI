package models

type Profile struct {
	User  User     `json:"user"`
	Chats []*Chats `json:"chats"`
}
