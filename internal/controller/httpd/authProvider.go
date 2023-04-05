package httpd

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/vyroai/VyroAI/commons/api/response"
	"path"
	"strings"
)

func (s *WebServiceHttpServer) generateUrl(c *fiber.Ctx) error {
	provider := c.Params("provider")
	action := path.Base(c.Path())

	c.Route()
	switch strings.ToLower(provider) {
	case "discord":
		discordUrl, _ := s.authService.GenerateDiscordAuthUrl(c.Context(), action)
		err := c.Redirect(discordUrl)
		if err != nil {
			response.ErrorJson(c, 500, "failed to generate discord url")
			return nil
		}
	case "google":
		discordUrl, _ := s.authService.GenerateDiscordAuthUrl(c.Context(), action)
		err := c.Redirect(discordUrl)
		if err != nil {
			response.ErrorJson(c, 500, "failed to generate discord url")
			return nil
		}
	case "instagram":
	case "apple":

	default:
		response.ErrorJson(c, 400, "provider not found")
	}

	return nil
}

func (s *WebServiceHttpServer) callbackLogin(c *fiber.Ctx) error {
	provider := c.Params("provider")
	code := c.Query("code")
	state := c.Query("state")
	fmt.Println(code)
	fmt.Println(state)

	switch strings.ToLower(provider) {
	case "discord":
	case "google":
	case "instagram":
	case "apple":

	}

	s.authService.DiscordProviderLogin(c.Context(), code, state)

	return nil
}

func (s *WebServiceHttpServer) callbackRegister(c *fiber.Ctx) error {
	provider := c.Params("provider")
	code := c.Query("code")
	state := c.Query("state")

	switch strings.ToLower(provider) {
	case "discord":
		register, err := s.authService.DiscordProviderRegister(c.Context(), code, state)
		if err != nil {
			response.ErrorJson(c, 400, err.Error())
			return nil
		}
		response.SuccessJson(c, 200, "created account", register)
	case "google":
	case "instagram":
	case "apple":
	default:
		response.ErrorJson(c, 400, "provider not found")
	}
	return nil

}
