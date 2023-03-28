package authentication

import (
	"context"
	"errors"
	"github.com/vyroai/VyroAI/commons/otel"
	"github.com/vyroai/VyroAI/internal/domain/authentication/entites"
	"github.com/vyroai/VyroAI/internal/domain/authentication/repo"
	"go.opentelemetry.io/otel/trace"
)

type Authentication interface {
	Login(ctx context.Context, email, password string) (*entites.User, error)
}

type AuthService struct {
	tracer     trace.Tracer
	userRepo   repo.AuthRepo
	bcryptRepo repo.BcryptRepo
}

func NewAuthService(userRepo repo.AuthRepo, bcryptRepo repo.BcryptRepo) Authentication {
	tracer := otel.InitTracing("authenticationService", "0.1.0")

	return &AuthService{
		tracer:   tracer.NewTracer(),
		userRepo: userRepo,
	}
}

func (as *AuthService) Login(ctx context.Context, email, password string) (*entites.User, error) {
	ctx, span := as.tracer.Start(ctx, "login")
	defer span.End()

	userResult, err := as.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	err = as.bcryptRepo.CompareHashAndPassword(userResult.Password, password)
	if err != nil {
		return nil, errors.New(`invalid email or password`)
	}

	return userResult, nil

}
