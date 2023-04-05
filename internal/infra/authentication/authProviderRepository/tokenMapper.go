package authProviderRepository

import (
	"github.com/vyroai/VyroAI/internal/domain/authentication/entites"
	"golang.org/x/oauth2"
)

func Oauth2TokenToModel(token *oauth2.Token) *entites.Token {
	return &entites.Token{
		AccessToken:  token.AccessToken,
		TokenType:    token.TokenType,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
	}
}
