package httpd

import (
	"github.com/gofiber/fiber/v2"
)

func (s *WebServiceHttpServer) MountDashboard(app *fiber.App) {
	auth := app.Group("/v1")

	auth.Get("/self", func(c *fiber.Ctx) error {

		return nil
	})

}
