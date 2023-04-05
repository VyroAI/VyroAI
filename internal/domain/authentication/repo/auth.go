package repo

import (
	"context"
	"github.com/vyroai/VyroAI/internal/domain/authentication/entites"
)

type AuthRepo interface {
	GetUserByID(ctx context.Context, userID int64) (*entites.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entites.User, error)
	GetUserByUsername(ctx context.Context, username string) (*entites.User, error)
	GetUserFromOAuthID(ctx context.Context, oauthID string) (*entites.User, error)

	CreateUser(ctx context.Context, username, email, password string) (int64, error)
	CreateUserWithOauthID(ctx context.Context, username, email, provider, accountID string) (int64, error)
}
