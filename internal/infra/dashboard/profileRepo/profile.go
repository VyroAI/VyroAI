package profileRepo

import (
	"github.com/jmoiron/sqlx"
	"github.com/vyroai/VyroAI/commons/otel"
	"github.com/vyroai/VyroAI/internal/domain/repo"
	"github.com/vyroai/VyroAI/internal/infra/database/sqlc"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type ProfileRepository struct {
	database *sqlc.Queries
	db       *sqlx.DB
	logger   *zap.Logger
	tracer   trace.Tracer
}

func NewProfileRepository(database *sqlx.DB, logger *zap.Logger) repo.Profile {
	tracer := otel.InitTracing("profileRepository", "0.1.0")

	return &ProfileRepository{
		database: sqlc.New(database),
		db:       database,
		logger:   logger,
		tracer:   tracer.NewTracer(),
	}
}
