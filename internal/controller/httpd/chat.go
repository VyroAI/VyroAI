package httpd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vyroai/VyroAI/commons/jwt"
)

func (s *WebServiceHttpServer) MountChat(app *fiber.App) {
	auth := app.Group("/v1/chat")

	auth.Use(jwt.New(jwt.Config{Permission: jwt.UserPermission}))

	auth.Get("/:chatbotID/messages", s.allMessages)
	auth.Post("/:chatbotID/message", s.createMessage)

}
