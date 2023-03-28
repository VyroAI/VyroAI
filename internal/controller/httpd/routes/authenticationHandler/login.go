package authenticationHandler

import (
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	fiber.Router
}

func (ah *AuthHandler) Login(c *fiber.Ctx) {

}
