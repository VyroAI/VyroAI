package httpd

import (
	"context"
	"fmt"
	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/vyroai/VyroAI/commons/otel"
	"github.com/vyroai/VyroAI/internal/domain/authentication"
	"github.com/vyroai/VyroAI/internal/domain/chat"
	"github.com/vyroai/VyroAI/internal/domain/dashboard"
	"log"
)

type WebServiceHttpServer struct {
	authService      authentication.Authentication
	dashboardService dashboard.Dashboard
	chatService      chat.Chat
}

func NewWebServiceHttpServer(authService authentication.Authentication, dashboardService dashboard.Dashboard, chatService chat.Chat) *WebServiceHttpServer {
	return &WebServiceHttpServer{
		authService:      authService,
		dashboardService: dashboardService,
		chatService:      chatService,
	}
}

func (s *WebServiceHttpServer) Router() *fiber.App {
	trace := otel.InitTracing("backend-api", "0.1.0")
	tp := trace.InitFiberTrace()

	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	app := fiber.New()
	app.Use(cors.New())

	app.Use(otelfiber.Middleware(otelfiber.WithSpanNameFormatter(func(ctx *fiber.Ctx) string {
		return fmt.Sprintf("%s - %s", ctx.Method(), ctx.Route().Path)
	})))

	s.MountAuth(app)

	s.MountDashboard(app)

	s.MountChat(app)

	app.Get("/health", func(ctx *fiber.Ctx) error {
		err := ctx.SendString("ok")
		if err != nil {
			return err
		}
		return nil
	})

	return app
}
