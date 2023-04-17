package httpd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vyroai/VyroAI/commons/api/response"
	"github.com/vyroai/VyroAI/commons/errors"
	"github.com/vyroai/VyroAI/commons/helper/formSantanizer"
	"github.com/vyroai/VyroAI/commons/jwt"
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

func (s *WebServiceHttpServer) login(c *fiber.Ctx) error {
	var login LoginRequestPayload

	if err := c.BodyParser(&login); err != nil {
		response.ErrorJson(c, 401, err.Error())
		return nil
	}

	if !formSantanizer.IsValidEmail(login.Email) {
		response.ErrorJson(c, 400, "invalid payload")
		return nil
	}

	if !formSantanizer.IsPasswordValid(login.Password) {
		response.ErrorJson(c, 400, "invalid payload")
		return nil
	}

	userID, permission, err := s.authService.Login(c.Context(), login.Email, login.Password)

	if err != nil {
		switch errors.GetType(err) {
		case errors.ErrInvalid:
			response.ErrorJson(c, 401, "Invalid Email or Password")
			return nil
		case errors.ErrNotFound:
			response.ErrorJson(c, 401, "Invalid Email or Password")
			return nil
		default:
			response.ServerError(c)
			return nil
		}
	}

	response.SuccessJson(c, 200, "Successfully created login", jwt.GenerateJwt(userID, permission))
	return nil
}

func (s *WebServiceHttpServer) register(c *fiber.Ctx) error {
	var register RegisterRequestPayload

	if err := c.BodyParser(&register); err != nil {
		response.ErrorJson(c, 401, err.Error())
		return nil
	}

	if !formSantanizer.IsValidEmail(register.Email) {
		response.ErrorJson(c, 400, "invalid payload")
		return nil
	}
	if formSantanizer.IsValidUsername(register.Username) {
		response.ErrorJson(c, 400, "invalid payload")
		return nil
	}
	if !formSantanizer.IsPasswordValid(register.Password) {
		response.ErrorJson(c, 400, "invalid payload")
		return nil
	}

	userID, permission, err := s.authService.Register(c.Context(), register.Username, register.Email, register.Password)
	if err != nil {
		switch errors.GetType(err) {
		case errors.ErrExist:
			response.ErrorJson(c, 409, err.Error())
			return nil
		default:
			response.ServerError(c)
			return nil
		}
	}

	response.SuccessJson(c, 201, "Successfully created an account", jwt.GenerateJwt(userID, permission))
	return nil

}
