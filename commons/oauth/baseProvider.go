package oauth

type baseProvider struct {
	scopes     []string
	authUrl    string
	tokenUrl   string
	userApiUrl string
}