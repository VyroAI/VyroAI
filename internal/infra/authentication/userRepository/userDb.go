package userRepository

import (
	"context"
	"github.com/vyroai/VyroAI/internal/domain/authentication/entites"
	"go.uber.org/zap"
)

func (ur *UserRepository) GetUserByID(ctx context.Context, userID int64) (*entites.User, error) {
	ctx, span := ur.tracer.Start(ctx, "get-user-by-id")
	defer span.End()

	user, err := ur.database.GetUserByID(ctx, userID)
	if err != nil {
		ur.logger.Error("failed to get user by id", zap.Error(err))
		return nil, err
	}

	return userIdDbToModel(user), nil
}

func (ur *UserRepository) GetUserByEmail(ctx context.Context, email string) (*entites.User, error) {
	ctx, span := ur.tracer.Start(ctx, "get-user-by-email")
	defer span.End()

	user, err := ur.database.GetUserByEmail(ctx, email)
	if err != nil {
		ur.logger.Error("failed to get user by email", zap.Error(err))
		return nil, err
	}

	return userEmailDbToModel(user), nil
}
