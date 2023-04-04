package oauth

import (
	"errors"
	"golang.org/x/oauth2"
	"net/http"
)

type AuthUser struct {
	Id           string         `json:"id"`
	Name         string         `json:"name"`
	Username     string         `json:"username"`
	Email        string         `json:"email"`
	AvatarUrl    string         `json:"avatarUrl"`
	RawUser      map[string]any `json:"rawUser"`
	AccessToken  string         `json:"accessToken"`
	RefreshToken string         `json:"refreshToken"`
}

type Provider interface {
	Client(token *oauth2.Token) *http.Client
	sendRawUserDataRequest(req *http.Request, token *oauth2.Token) ([]byte, error)
	FetchRawUserData(token *oauth2.Token) ([]byte, error)
	BuildAuthUrl(state string, opts ...oauth2.AuthCodeOption) string
}

func NewProviderByName(name string, base *BaseProvider) (Provider, error) {
	switch name {
	case "discord":
		return discordProvider(base), nil

	case "google":
		return googleProvider(base), nil

	}
	return nil, errors.New("provider not found")
}
