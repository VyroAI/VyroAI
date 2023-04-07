package userRepository

import (
	"context"
	"database/sql"
	"github.com/vyroai/VyroAI/commons/snowflake"
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
	ctx, span := ur.tracer.Start(ctx, "create-user-subscription")
	defer span.End()

	tx, err := ur.db.BeginTx(ctx, nil)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		ur.logger.Error(err.Error())
		return -1, err
	}
	defer tx.Rollback()

	qtx := ur.database.WithTx(tx)

	userID, err := qtx.CreateUser(ctx, sqlc.CreateUserParams{
		Username: username,
		Email:    email,
		Password: sql.NullString{String: password, Valid: true},
	})

	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		ur.logger.Error(err.Error())
		return -1, err
	}

	apikey := snowflake.GenerateSha1SnowflakeIDWithTime()

	err = qtx.CreateUserSubscription(ctx, sqlc.CreateUserSubscriptionParams{
		UserID: userID,
		ApiKey: apikey,
	})
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		ur.logger.Error(err.Error())
		return -1, err
	}

	err = tx.Commit()
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		ur.logger.Error(err.Error())
		return -1, err
	}

	return userID, nil
}

func (ur *UserRepository) GetUserFromOAuthID(ctx context.Context, oauthID string) (*entites.User, error) {
	ctx, span := ur.tracer.Start(ctx, "get-user-by-oauth-id")
	defer span.End()

	user, err := ur.database.GetUserByOAuthID(ctx, oauthID)
	if err != nil {
		return nil, err
	}

	return userOauthToDbToModel(user), nil
}

func (ur *UserRepository) CreateUserWithOauthID(ctx context.Context, username, email, provider, accountID string) (int64, error) {
	ctx, span := ur.tracer.Start(ctx, "create-user-with-oauth-id")
	defer span.End()

	tx, err := ur.db.BeginTx(ctx, nil)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		ur.logger.Error(err.Error())
		return -1, err
	}
	defer tx.Rollback()
	qtx := ur.database.WithTx(tx)

	userID, err := qtx.CreateUser(ctx, sqlc.CreateUserParams{
		Username: username,
		Email:    email,
		Password: sql.NullString{Valid: false},
	})

	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		ur.logger.Error(err.Error())
		return -1, err
	}

	apikey := snowflake.GenerateSha1SnowflakeIDWithTime()

	err = qtx.CreateUserSubscription(ctx, sqlc.CreateUserSubscriptionParams{
		UserID: userID,
		ApiKey: apikey,
	})
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		ur.logger.Error(err.Error())
		return -1, err
	}

	err = qtx.CreateOAuthAccount(ctx, sqlc.CreateOAuthAccountParams{
		UserID: userID,
		OauthProvider: sqlc.NullOauthAccountOauthProvider{
			OauthAccountOauthProvider: sqlc.OauthAccountOauthProvider(provider),
			Valid:                     true,
		},
		AccountID: accountID,
	})
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		ur.logger.Error(err.Error())
		return -1, err
	}

	err = tx.Commit()
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		ur.logger.Error(err.Error())
		return -1, err
	}

	return userID, nil
}
