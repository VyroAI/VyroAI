package userRepository

import (
	"github.com/vyroai/VyroAI/internal/domain/authentication/entites"
	"github.com/vyroai/VyroAI/internal/infra/database/sqlc"
)

func userIdDbToModel(userDB sqlc.GetUserByIDRow) *entites.User {
	return &entites.User{
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

func userEmailDbToModel(userDB sqlc.GetUserByEmailRow) *entites.User {
	return &entites.User{
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

func userNameToDbToModel(userDB sqlc.GetUserByUsernameRow) *entites.User {
	return &entites.User{
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
