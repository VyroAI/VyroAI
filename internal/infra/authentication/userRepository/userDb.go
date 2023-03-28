package userRepository

import (
	"context"
	"github.com/vyroai/VyroAI/internal/domain/authentication/entites"
	"github.com/vyroai/VyroAI/internal/infra/database/sqlc"
	"go.opentelemetry.io/otel/codes"
)

func (ur *UserRepository) GetUserByID(ctx context.Context, userID int64) (*entites.User, error) {
	ctx, span := ur.tracer.Start(ctx, "get-user-by-id")
	defer span.End()

	user, err := ur.database.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return userIdDbToModel(user), nil
}

func (ur *UserRepository) GetUserByEmail(ctx context.Context, email string) (*entites.User, error) {
	ctx, span := ur.tracer.Start(ctx, "get-user-by-email")
	defer span.End()

	user, err := ur.database.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return userEmailDbToModel(user), nil
}

func (ur *UserRepository) GetUserByUsername(ctx context.Context, username string) (*entites.User, error) {
	ctx, span := ur.tracer.Start(ctx, "get-user-by-username")
	defer span.End()

	user, err := ur.database.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	return userNameToDbToModel(user), nil
}

func (ur *UserRepository) CreateUser(ctx context.Context, username, email, password string) (int64, error) {
	ctx, span := ur.tracer.Start(ctx, "create-user")
	defer span.End()

	userID, err := ur.database.CreateUser(ctx, sqlc.CreateUserParams{
		Username: username,
		Email:    email,
		Password: password,
	})
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		ur.logger.Error(err.Error())
		return -1, err
	}

	return userID, nil
}
