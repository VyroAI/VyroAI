package httpd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/vyroai/VyroAI/internal/domain/authentication"
)

type WebServiceHttpServer struct {
	authService authentication.Authentication
}

func NewWebServiceHttpServer(authService authentication.Authentication) *WebServiceHttpServer {
	return &WebServiceHttpServer{
		authService: authService,
	}
}

func (s *WebServiceHttpServer) Router() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())

	s.MountAuth(app)

	return app
}
