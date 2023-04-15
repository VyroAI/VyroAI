package response

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func ServerError(c *fiber.Ctx) {
	messageResponse := &Payload{
		Message: "Internal Server Error",
	}
	c.Set(`Content-Type`, `application/json; charset=utf-8`)
	c.Status(500)
	_ = c.JSON(messageResponse)
}

func NotFoundError(c *fiber.Ctx) {
	messageResponse := &Payload{
		Message: "The requested resource wasn't found.",
	}
	c.Set(`Content-Type`, `application/json; charset=utf-8`)
	c.Status(http.StatusNotFound)
	_ = c.JSON(messageResponse)
}
