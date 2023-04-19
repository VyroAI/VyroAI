package httpd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vyroai/VyroAI/commons/api/response"
	"github.com/vyroai/VyroAI/commons/errors"
)

func (s *WebServiceHttpServer) profile(c *fiber.Ctx) error {
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
}
