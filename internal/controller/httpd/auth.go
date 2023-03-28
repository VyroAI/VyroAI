package httpd

import (
	"github.com/gofiber/fiber/v2"
)

func (s *WebServiceHttpServer) MountAuth(app *fiber.App) {
	auth := app.Group("/v1/auth")

	auth.Post("/login", func(c *fiber.Ctx) error {
		s.authService.Login(c.Context(), "ab@gmail.com", "dsadsds")
		return c.SendString("received!")
	})

}
