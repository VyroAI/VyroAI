package models

import "time"

type Chats struct {
	Id    int64  `json:"chat_id"`
	Title string `json:"title"`
}

type ChatMessages struct {
	ChatMessage []*Message `json:"messages"`
}
type Message struct {
	MessageID int64     `json:"id,omitempty"`
	ChatBotID int64     `json:"chat_id,omitempty"`
	UserId    int64     `json:"user_id,omitempty"`
	IsBot     bool      `json:"bot"`
	Message   string    `json:"message,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
