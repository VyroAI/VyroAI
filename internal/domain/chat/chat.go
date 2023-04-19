package chat

import (
	"context"
	"github.com/vyroai/VyroAI/internal/domain/models"
	"github.com/vyroai/VyroAI/internal/domain/repo"
)

type Chat interface {
	CreateMessage(ctx context.Context, content string, chatID int64) (*models.Message, error)
	GetChatMessagesByChatID(ctx context.Context, chatID int64, limit, offset int32) (*models.ChatMessages, error)
}

type Service struct {
	chatRepo repo.Chat
}

func NewChatService(chatRepo repo.Chat) Chat {
	return &Service{
		chatRepo: chatRepo,
	}
}

func (s *Service) GetChatMessagesByChatID(ctx context.Context, chatID int64, limit, offset int32) (*models.ChatMessages, error) {
	return s.chatRepo.GetChatMessagesByChatID(ctx, chatID, ctx.Value("userID").(int64), limit, offset)
}

func (s *Service) CreateMessage(ctx context.Context, content string, chatID int64) (*models.Message, error) {
	return s.chatRepo.CreateMessage(ctx, content, chatID, ctx.Value("userID").(int64))
}
