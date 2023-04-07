package authentication

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/vyroai/VyroAI/internal/domain/authentication/entites"
	"log"
	"strings"
)

func (as *AuthService) GenerateDiscordAuthUrl(ctx context.Context, action string) (string, error) {
	ctx, span := as.tracer.Start(ctx, "generate-discord-oauth-url")
	defer span.End()

	var url string

	switch strings.ToLower(action) {
	case "login":
		url = as.authProvider.DiscordProvider.BuildLoginUrl("hello")
	case "register":
		url = as.authProvider.DiscordProvider.BuildRegisterUrl("hello")

	}
	return url, nil
}

func (as *AuthService) DiscordProviderLogin(ctx context.Context, code, state string) (int64, int32, error) {
	ctx, span := as.tracer.Start(ctx, "discord-provider-login")
	defer span.End()

	token, err := as.authProvider.DiscordProvider.ExchangeCode(code)
	if err != nil {
		log.Println(err)
		return -1, -1, err
	}

	discordUser, err := as.authProvider.DiscordProvider.FetchRawUserData(token)
	if err != nil {
		log.Println(err)
		return -1, -1, err
	}
	var discord entites.DiscordUser

	err = json.Unmarshal(discordUser, &discord)
	if err != nil {
		return -1, -1, err
	}

	user, err := as.userRepo.GetUserFromOAuthID(ctx, discord.Id)
	if err != nil {
		fmt.Println(err)
		return -1, -1, err
	}

	return user.Id, user.Permission, err
}

func (as *AuthService) DiscordProviderRegister(ctx context.Context, code, state string) (int64, int32, error) {
	ctx, span := as.tracer.Start(ctx, "discord-provider-register")
	defer span.End()

	token, err := as.authProvider.DiscordProvider.ExchangeCode(code)
	if err != nil {
		return -1, -1, err
	}

	discordUser, err := as.authProvider.DiscordProvider.FetchRawUserData(token)
	if err != nil {
		return -1, -1, err
	}
	var discord entites.DiscordUser

	err = json.Unmarshal(discordUser, &discord)
	if err != nil {
		return -1, -1, err
	}

	_, err = as.userRepo.GetUserFromOAuthID(ctx, discord.Id)
	if err == nil {
		return -1, -1, errors.New("account already exist")
	}

	_, err = as.userRepo.GetUserByEmail(ctx, discord.Email)
	if err == nil {
		return -1, -1, errors.New("email already exist")
	}

	userID, err := as.userRepo.CreateUserWithOauthID(ctx, discord.Username, discord.Email, "discord", discord.Id)
	if err != nil {
		return -1, -1, errors.New(`server error`)
	}

	return userID, 1, nil

}
