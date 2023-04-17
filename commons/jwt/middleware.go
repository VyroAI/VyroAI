package jwt

import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

const (
	UserPermission  = Permission(iota)
	AdminPermission = Permission(iota)
)

type Permission = uint

type Config struct {
	Permission   Permission
	Unauthorized fiber.Handler
}

func unauthorized(c *fiber.Ctx) error {
	c.Set(`Content-Type`, `application/json; charset=utf-8`)
	c.Status(fiber.StatusUnauthorized)
	return c.SendString(`{
    "data": null,
    "error": "Unauthorized",
    "message": null
}`)
}

func New(cfg Config) fiber.Handler {

	return func(c *fiber.Ctx) error {
		authorizationBearer := c.Get("authorization")

		claims, err := VerifyJwt(strings.Split(authorizationBearer, " ")[1], cfg.Permission)
		if err == nil {
			c.Locals("userID", claims.UserId)
			return c.Next()
		}

		return unauthorized(c)
	}
}
