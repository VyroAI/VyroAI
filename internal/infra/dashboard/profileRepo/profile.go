package profileRepo

import (
	"github.com/jmoiron/sqlx"
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

func NewProfileRepository() {}
