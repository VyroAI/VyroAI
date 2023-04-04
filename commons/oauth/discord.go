package oauth

type Discord struct {
	*BaseProvider
}

func discordProvider(base *BaseProvider) *Discord {
	return &Discord{base}
}
