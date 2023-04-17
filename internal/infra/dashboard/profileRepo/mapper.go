package profileRepo

import (
	"github.com/vyroai/VyroAI/internal/domain/models"
	"github.com/vyroai/VyroAI/internal/infra/database/sqlc"
)

func profileToModel(profile []sqlc.GetProfileAndChatsRow) *models.Profile {
	var chats []*models.Chats
	for i := 0; i < len(profile); i++ {
		if profile[i].ChatbotID.Valid {
			chats = append(chats, &models.Chats{
				Id:    profile[i].ChatbotID.Int64,
				Title: profile[i].Title.String,
			})
		}

	}

	if chats == nil {
		chats = []*models.Chats{}
	}

	return &models.Profile{
		User: models.User{
			Username:       profile[0].Username,
			Email:          profile[0].Email,
			AvatarID:       profile[0].AvatarID,
			Permission:     profile[0].Permission,
			IsBanned:       profile[0].IsBanned,
			EmailConfirmed: profile[0].EmailConfirmed,
			CreatedAt:      profile[0].CreatedAt,
		},
		Chats: chats,
	}
}
