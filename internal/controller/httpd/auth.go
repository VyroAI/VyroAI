package httpd

import "github.com/gofiber/fiber/v2"

func (s *WebServiceHttpServer) MountAuth(app *fiber.App) {
	auth := app.Group("/v1/auth")

	auth.Post("/login", s.login)
	auth.Post("/register", s.register)

	auth.Get("/:provider/login", s.generateUrl)
	auth.Get("/:provider/register", s.generateUrl)

	auth.Get("/:provider/login/callback", s.callbackLogin)
	auth.Get("/:provider/register/callback", s.callbackRegister)

}
