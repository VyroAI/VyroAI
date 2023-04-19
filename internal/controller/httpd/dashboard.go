package httpd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vyroai/VyroAI/commons/jwt"
)

func (s *WebServiceHttpServer) MountDashboard(app *fiber.App) {
	auth := app.Group("/v1")

	auth.Use(jwt.New(jwt.Config{Permission: jwt.UserPermission}))

	auth.Get("/self", s.profile)

}
