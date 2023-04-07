package authProviderRepository

import (
	"context"
	"github.com/vyroai/VyroAI/commons/oauth"
	"os"
)

type AuthProvider struct {
	DiscordProvider oauth.Provider
	GoogleProvider  oauth.Provider
}

func NewAuthProvider() *AuthProvider {
	discord, _ := oauth.NewProviderByName("discord",
		&oauth.BaseProvider{
			Ctx:                 context.Background(),
			Scopes:              []string{"identify", "email"},
			RedirectLoginUrl:    os.Getenv("BASE_URL") + "v1/auth/discord/login/callback",
			RedirectRegisterUrl: os.Getenv("BASE_URL") + "v1/auth/discord/register/callback",
			ClientId:            os.Getenv("DISCORD_CLIENT_ID"),
			ClientSecret:        os.Getenv("DISCORD_CLIENT_SECRET"),
			AuthUrl:             "https://discord.com/api/oauth2/authorize",
			TokenUrl:            "https://discord.com/api/oauth2/token",
			UserApiUrl:          "https://discord.com/api/users/@me",
		})

	google, _ := oauth.NewProviderByName("google",
		&oauth.BaseProvider{
			Ctx: context.Background(),
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.profile",
				"https://www.googleapis.com/auth/userinfo.email",
			},
			RedirectLoginUrl:    os.Getenv("BASE_URL") + "v1/auth/google/login/callback",
			RedirectRegisterUrl: os.Getenv("BASE_URL") + "v1/auth/google/register/callback",
			ClientId:            os.Getenv("GOOGLE_CLIENT_ID"),
			ClientSecret:        os.Getenv("GOOGLE_CLIENT_SECRET"),
			AuthUrl:             "https://accounts.google.com/o/oauth2/auth",
			TokenUrl:            "https://accounts.google.com/o/oauth2/token",
			UserApiUrl:          "https://www.googleapis.com/oauth2/v1/userinfo",
		})

	return &AuthProvider{
		DiscordProvider: discord,
		GoogleProvider:  google,
	}
}
