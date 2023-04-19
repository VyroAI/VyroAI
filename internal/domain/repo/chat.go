package repo

import (
	"context"
	"github.com/vyroai/VyroAI/internal/domain/models"
)

type Chat interface {
	GetChatMessagesByChatID(ctx context.Context, chatID, userID int64, limit, offset int32) (*models.ChatMessages, error)
	CreateMessage(ctx context.Context, content string, chatID, userID int64) (*models.Message, error)
}
