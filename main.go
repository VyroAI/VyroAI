package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/vyroai/VyroAI/commons/database/sql"
	"github.com/vyroai/VyroAI/commons/logger"
	"github.com/vyroai/VyroAI/internal/controller/httpd"
	"github.com/vyroai/VyroAI/internal/domain/authentication"
	"github.com/vyroai/VyroAI/internal/infra/authentication/bcryptRepository"
	"github.com/vyroai/VyroAI/internal/infra/authentication/userRepository"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func init() {
	godotenv.Load(".env")
}

func main() {

	fx.New(
		fx.Provide(logger.InitLogger),

		fx.Provide(sql.NewSqlConn),

		fx.Provide(userRepository.NewUserRepository),
		fx.Provide(bcryptRepository.NewBcryptRepository),
		fx.Provide(authentication.NewAuthService),

		fx.Provide(httpd.NewWebServiceHttpServer),

		fx.Invoke(StartHttpRouter),
	).Run()
}

func StartHttpRouter(lc fx.Lifecycle, httpd *httpd.WebServiceHttpServer, logger *zap.Logger) {
	var app *fiber.App
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			app = httpd.Router()
			go func() {
				if err := app.Listen(":3000"); err != nil {
					logger.Fatal("Failed to start internal http server", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			err := app.Shutdown()
			if err != nil {
				return err
			}
			return nil
		},
	})
}
