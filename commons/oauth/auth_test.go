package oauth

import (
	"context"
	"fmt"
	"testing"
)

func TestDiscordProvider(t *testing.T) {
	b := BaseProvider{
		AuthUrl:      "https://discord.com/api/oauth2/authorize",
		TokenUrl:     "tokenUrl_test",
		RedirectUrl:  "https://api.vyroai.com/v1/auth/callback/discord",
		ClientId:     "1092026012570751067",
		ClientSecret: "pa4ccVHlcfqFTN9dCr_7vO5nnsf6f1CP",
		Scopes:       []string{"identify", "email"},
	}
	fmt.Println(b.BuildAuthUrl("Hello"))
}

func TestNewDiscordProvider(t *testing.T) {
	discord, _ := NewProviderByName("discord",
		&BaseProvider{
			Ctx:          context.Background(),
			Scopes:       []string{"identify", "email"},
			RedirectUrl:  "https://api.vyroai.com/v1/auth/callback/discord",
			ClientId:     "1092026012570751067",
			ClientSecret: "pa4ccVHlcfqFTN9dCr_7vO5nnsf6f1CP",
			AuthUrl:      "https://discord.com/api/oauth2/authorize",
			TokenUrl:     "https://discord.com/api/oauth2/token",
			UserApiUrl:   "https://discord.com/api/users/@me",
		})

	fmt.Println(discord.BuildAuthUrl("hello"))
}
