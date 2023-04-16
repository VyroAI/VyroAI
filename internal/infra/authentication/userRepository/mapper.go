package userRepository

import (
	"github.com/vyroai/VyroAI/internal/domain/models"
	"github.com/vyroai/VyroAI/internal/infra/database/sqlc"
)

func userIdDbToModel(userDB sqlc.GetUserByIDRow) *models.User {
	return &models.User{
		Id:             userDB.ID,
		Username:       userDB.Username,
		Email:          userDB.Email,
		AvatarID:       userDB.AvatarID,
		Permission:     userDB.Permission,
		IsBanned:       userDB.IsBanned,
		EmailConfirmed: userDB.EmailConfirmed,
		CreatedAt:      userDB.CreatedAt,
	}
}

func userEmailDbToModel(userDB sqlc.GetUserByEmailRow) *models.User {
	return &models.User{
		Id:             userDB.ID,
		Username:       userDB.Username,
		Email:          userDB.Email,
		Password:       userDB.Password.String,
		AvatarID:       userDB.AvatarID,
		Permission:     userDB.Permission,
		IsBanned:       userDB.IsBanned,
		EmailConfirmed: userDB.EmailConfirmed,
		CreatedAt:      userDB.CreatedAt,
	}
}

func userNameToDbToModel(userDB sqlc.GetUserByUsernameRow) *models.User {
	return &models.User{
		Id:             userDB.ID,
		Username:       userDB.Username,
		Email:          userDB.Email,
		AvatarID:       userDB.AvatarID,
		Permission:     userDB.Permission,
		IsBanned:       userDB.IsBanned,
		EmailConfirmed: userDB.EmailConfirmed,
		CreatedAt:      userDB.CreatedAt,
	}
}

func userOauthToDbToModel(userDB sqlc.GetUserByOAuthIDRow) *models.User {
	return &models.User{
		Id:             userDB.ID,
		Username:       userDB.Username,
		Email:          userDB.Email,
		AvatarID:       userDB.AvatarID,
		Permission:     userDB.Permission,
		IsBanned:       userDB.IsBanned,
		EmailConfirmed: userDB.EmailConfirmed,
		CreatedAt:      userDB.CreatedAt,
	}
}
