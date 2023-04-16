package repo

import (
	"context"
	"github.com/vyroai/VyroAI/internal/domain/models"
)

type AuthRepo interface {
	GetUserByID(ctx context.Context, userID int64) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	GetUserFromOAuthID(ctx context.Context, oauthID string) (*models.User, error)

	CreateUser(ctx context.Context, username, email, password string) (int64, error)
	CreateUserWithOauthID(ctx context.Context, username, email, provider, accountID string) (int64, error)
}
