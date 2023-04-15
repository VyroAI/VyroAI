package response

import (
	"github.com/gofiber/fiber/v2"
)

type Payload struct {
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
	Message interface{} `json:"message"`
}

func SuccessJson(c *fiber.Ctx, statusCode int, message string, data interface{}) {
	messageResponse := &Payload{
		Data:    &data,
		Message: message,
	}
	c.Set(`Content-Type`, `application/json; charset=utf-8`)
	c.Status(statusCode)
	_ = c.JSON(messageResponse)
}

func ErrorJson(c *fiber.Ctx, statusCode int, error string) {
	errorResponse := &Payload{
		Error: error,
	}
	c.Set(`Content-Type`, `application/json; charset=utf-8`)
	c.Status(statusCode)
	_ = c.JSON(errorResponse)
}

func SuccessDataJson(c *fiber.Ctx, statusCode int, data interface{}) {
	messageResponse := &Payload{
		Data: &data,
	}
	c.Set(`Content-Type`, `application/json; charset=utf-8`)
	c.Status(statusCode)
	_ = c.JSON(messageResponse)
}

func SuccessMessage(c *fiber.Ctx, statusCode int, message string) {
	messageResponse := &Payload{
		Message: message,
	}
	c.Set(`Content-Type`, `application/json; charset=utf-8`)
	c.Status(statusCode)
	_ = c.JSON(messageResponse)
}
