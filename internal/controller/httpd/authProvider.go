package httpd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vyroai/VyroAI/commons/api/response"
	"github.com/vyroai/VyroAI/commons/jwt"
	"os"
	"path"
	"strings"
	"time"
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

	var userID int64
	var permission int32
	var err error

	switch strings.ToLower(provider) {
	case "discord":
		userID, permission, err = s.authService.DiscordProviderLogin(c.Context(), code, state)
		if err != nil {
			response.ErrorJson(c, 400, err.Error())
			authCookie := &fiber.Cookie{
				Name:    "oauth_error",
				Value:   "a03f472b2cc9b6e15e500e2d8bb286a658c4037f",
				Expires: time.Now().Add(10 * time.Second),
			}
			c.Cookie(authCookie)
			err := c.Redirect(os.Getenv("WEBAPP_BASE_URL") + "login")
			if err != nil {
				return err
			}
			return nil
		}
	case "google":
	case "instagram":
	case "apple":

	}
	jwtToken := jwt.GenerateJwt(userID, permission)

	authCookie := &fiber.Cookie{
		Name:    "authorization",
		Value:   "bearer " + jwtToken,
		Expires: time.Now().Add(240 * time.Hour),
	}
	c.Cookie(authCookie)
	err = c.Redirect(os.Getenv("WEBAPP_BASE_URL") + "dashboard")
	if err != nil {
		return err
	}
	return nil

}

func (s *WebServiceHttpServer) callbackRegister(c *fiber.Ctx) error {
	provider := c.Params("provider")
	code := c.Query("code")
	state := c.Query("state")

	var userID int64
	var permission int32
	var err error

	switch strings.ToLower(provider) {
	case "discord":
		userID, permission, err = s.authService.DiscordProviderRegister(c.Context(), code, state)
		if err != nil {
			response.ErrorJson(c, 400, err.Error())
			return nil
		}

	case "google":
	case "instagram":
	case "apple":
	default:
		response.ErrorJson(c, 400, "provider not found")
	}

	jwtToken := jwt.GenerateJwt(userID, permission)

	authCookie := &fiber.Cookie{
		Name:    "authorization",
		Value:   "bearer " + jwtToken,
		Expires: time.Now().Add(240 * time.Hour),
	}
	c.Cookie(authCookie)
	err = c.Redirect(os.Getenv("WEBAPP_BASE_URL") + "dashboard")
	if err != nil {
		return err
	}
	return nil

}
