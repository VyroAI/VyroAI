package chatRepo

import (
	"context"
	"fmt"
	"github.com/vyroai/VyroAI/internal/domain/models"
	"github.com/vyroai/VyroAI/internal/infra/database/sqlc"
	"go.opentelemetry.io/otel/codes"
)

func (cr *ChatRepository) GetChatMessagesByChatID(ctx context.Context, chatID, userID int64, limit, offset int32) (*models.ChatMessages, error) {
	ctx, span := cr.tracer.Start(ctx, "get-chat-messages-by-chat-id")
	defer span.End()
	messages, err := cr.database.GetChatMessageByChatID(ctx, sqlc.GetChatMessageByChatIDParams{
		ChatbotID: chatID,
		UserID:    userID,
		Limit:     limit,
		Offset:    offset,
	})
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, fmt.Sprintf("unknown error,  %+v\n", err))
		cr.logger.Error(fmt.Sprintf("unknown error,  %+v\n", err))
		return nil, err
	}
	if len(messages) == 0 {

	}

	return &models.ChatMessages{
		ChatMessage: chatToModel(messages),
	}, nil
}

func (cr *ChatRepository) CreateMessage(ctx context.Context, content string, chatID, userID int64) (*models.Message, error) {
	ctx, span := cr.tracer.Start(ctx, "create-message")
	defer span.End()
	messageID, err := cr.database.CreateMessage(ctx, sqlc.CreateMessageParams{
		Content:   content,
		CreatedBy: userID,
		ChatbotID: chatID,
	})
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, fmt.Sprintf("unknown error,  %+v\n", err))
		cr.logger.Error(fmt.Sprintf("unknown error,  %+v\n", err))
	}

	return &models.Message{
		MessageID: messageID,
		ChatBotID: chatID,
		UserId:    userID,
	}, err
}
