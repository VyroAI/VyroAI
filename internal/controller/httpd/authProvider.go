package httpd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vyroai/VyroAI/commons/api/response"
)

func (s *WebServiceHttpServer) discordLogin(c *fiber.Ctx) error {
	discordUrl, _ := s.authService.GenerateDiscordAuthUrl(c.Context())

	err := c.Redirect(discordUrl)
	if err != nil {
		response.ErrorJson(c, 500, "Failed to generate discord url")
		return nil
	}
	return nil
}
