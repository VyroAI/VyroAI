package httpd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vyroai/VyroAI/commons/api/response"
	"github.com/vyroai/VyroAI/commons/errors"
	"github.com/vyroai/VyroAI/commons/jwt"
)

func (s *WebServiceHttpServer) MountDashboard(app *fiber.App) {
	auth := app.Group("/v1")

	auth.Use(jwt.New(jwt.Config{Permission: jwt.UserPermission}))

	auth.Get("/self", func(c *fiber.Ctx) error {
		profile, err := s.dashboardService.GetProfile(c.Context())
		if err != nil {
			switch errors.GetType(err) {
			case errors.ErrNotFound:
				response.ErrorJson(c, 401, "Unauthorized")
				return nil
			default:
				response.ServerError(c)
				return nil
			}
		}

		response.SuccessDataJson(c, 200, profile)
		return nil
	})

}
