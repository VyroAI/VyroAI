package userRepository

import (
	"github.com/jmoiron/sqlx"
	"github.com/vyroai/VyroAI/commons/otel"
	"github.com/vyroai/VyroAI/internal/domain/authentication/repo"
	"github.com/vyroai/VyroAI/internal/infra/database/sqlc"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type UserRepository struct {
	database *sqlc.Queries
	db       *sqlx.DB
	logger   *zap.Logger
	tracer   trace.Tracer
}

func NewUserRepository(database *sqlx.DB, logger *zap.Logger) repo.AuthRepo {
	tracer := otel.InitTracing("userRepository", "0.1.0")
	return &UserRepository{
		database: sqlc.New(database),
		db:       database,
		logger:   logger,
		tracer:   tracer.NewTracer(),
	}
}
