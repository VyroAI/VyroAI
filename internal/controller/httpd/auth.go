package httpd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vyroai/VyroAI/commons/api/response"
)

type LoginRequestPayload struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Recaptcha string `json:"reCaptcha"`
}

type RegisterRequestPayload struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Recaptcha string `json:"reCaptcha"`
}

func (s *WebServiceHttpServer) MountAuth(app *fiber.App) {
	auth := app.Group("/v1/auth")

	auth.Post("/login", func(c *fiber.Ctx) error {
		var login LoginRequestPayload

		if err := c.BodyParser(&login); err != nil {
			response.ErrorJson(c, 401, err.Error())
			return nil
		}

		user, err := s.authService.Login(c.Context(), login.Email, login.Password)
		if err != nil {
			response.ErrorJson(c, 401, "invalid email or password")
			return nil
		}

		response.SuccessJson(c, 200, "", user)
		return nil

	})

	auth.Post("/register", func(c *fiber.Ctx) error {
		var register RegisterRequestPayload

		if err := c.BodyParser(&register); err != nil {
			response.ErrorJson(c, 401, err.Error())
			return nil
		}

		user, err := s.authService.Register(c.Context(), register.Username, register.Email, register.Password)
		if err != nil {
			response.ErrorJson(c, 401, err.Error())
			return nil
		}

		response.SuccessJson(c, 200, "", user)
		return nil

	})
}
