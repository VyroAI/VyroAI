package bcryptRepository

import (
	"context"
	"github.com/vyroai/VyroAI/commons/otel"
	"github.com/vyroai/VyroAI/internal/domain/repo"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/crypto/bcrypt"
)

type Bcrypt struct {
	tracer trace.Tracer
}

func NewBcryptRepository() repo.BcryptRepo {
	tracer := otel.InitTracing("bcryptRepository", "0.1.0")

	return &Bcrypt{
		tracer: tracer.NewTracer(),
	}
}

func (b *Bcrypt) CompareHashAndPassword(ctx context.Context, hashedPassword, password string) error {
	ctx, span := b.tracer.Start(ctx, "compare-hash-and-password")
	defer span.End()

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func (b *Bcrypt) GenerateFromPassword(ctx context.Context, password string) (string, error) {
	ctx, span := b.tracer.Start(ctx, "generate-from-password")
	defer span.End()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
