package oauth

type Discord struct {
	*baseProvider
}

func NewDiscordProvider() *Discord {
	return &Discord{&baseProvider{
		scopes:     []string{"identify", "email"},
		authUrl:    "https://discord.com/api/oauth2/authorize",
		tokenUrl:   "https://discord.com/api/oauth2/token",
		userApiUrl: "https://discord.com/api/users/@me",
	}}
}
