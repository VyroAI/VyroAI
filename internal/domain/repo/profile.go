package repo

import (
	"context"
	"github.com/vyroai/VyroAI/internal/domain/models"
)

type Profile interface {
	GetProfileWithChat(ctx context.Context, userID int64) (*models.Profile, error)
}
