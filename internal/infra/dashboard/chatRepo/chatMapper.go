package chatRepo

import (
	"github.com/vyroai/VyroAI/internal/domain/models"
	"github.com/vyroai/VyroAI/internal/infra/database/sqlc"
)

func chatToModel(chats []sqlc.GetChatMessageByChatIDRow) []*models.Message {
	var chatList []*models.Message

	for i := 0; i < len(chats); i++ {
		chatList = append(chatList, &models.Message{
			MessageID: chats[i].ID,
			ChatBotID: chats[i].ChatbotID.Int64,
			UserId:    chats[i].CreatedBy,
			IsBot:     chats[i].Bot,
			Message:   chats[i].Content,
			CreatedAt: chats[i].CreatedAt.Time,
		})
	}

	if chatList == nil {
		chatList = []*models.Message{}
	}

	return chatList
}
