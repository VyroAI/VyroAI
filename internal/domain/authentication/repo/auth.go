package repo

import (
	"context"
	"github.com/vyroai/VyroAI/internal/domain/authentication/entites"
)

type AuthRepo interface {
	GetUserByID(ctx context.Context, userID int64) (*entites.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entites.User, error)
}
