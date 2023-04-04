package authentication

import (
	"context"
)

func (as *AuthService) GenerateDiscordAuthUrl(ctx context.Context) (string, error) {
	ctx, span := as.tracer.Start(ctx, "generate-discord-oauth-url")
	defer span.End()

	url := as.authProvider.DiscordProvider.BuildAuthUrl("hello")
	return url, nil
}
