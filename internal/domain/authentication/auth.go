package authentication

import (
	"context"
	"errors"
	"fmt"
	"github.com/vyroai/VyroAI/commons/otel"
	"github.com/vyroai/VyroAI/internal/domain/authentication/repo"
	"go.opentelemetry.io/otel/trace"
)

type Authentication interface {
	Login(ctx context.Context, email, password string) (int64, error)
	Register(ctx context.Context, username, email, password string) (int64, error)
}

type AuthService struct {
	tracer     trace.Tracer
	userRepo   repo.AuthRepo
	bcryptRepo repo.BcryptRepo
}

func NewAuthService(userRepo repo.AuthRepo, bcryptRepo repo.BcryptRepo) Authentication {
	tracer := otel.InitTracing("authenticationService", "0.1.0")

	return &AuthService{
		tracer:     tracer.NewTracer(),
		userRepo:   userRepo,
		bcryptRepo: bcryptRepo,
	}
}

func (as *AuthService) Login(ctx context.Context, email, password string) (int64, error) {
	ctx, span := as.tracer.Start(ctx, "login")
	defer span.End()

	userResult, err := as.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}
	fmt.Println(err)
	err = as.bcryptRepo.CompareHashAndPassword(ctx, userResult.Password, password)
	if err != nil {
		return -1, errors.New(`invalid email or password`)
	}

	return userResult.Id, nil
}

func (as *AuthService) Register(ctx context.Context, username, email, password string) (int64, error) {
	ctx, span := as.tracer.Start(ctx, "register")
	defer span.End()

	_, err := as.userRepo.GetUserByEmail(ctx, email)
	if err == nil {
		return -1, errors.New("email already exist")
	}

	_, err = as.userRepo.GetUserByUsername(ctx, username)
	if err == nil {
		return -1, errors.New("username already exist")
	}

	hashedPassword, err := as.bcryptRepo.GenerateFromPassword(ctx, password)
	if err != nil {
		return -1, errors.New(`server error`)
	}

	userID, err := as.userRepo.CreateUser(ctx, username, email, hashedPassword)
	if err != nil {
		return -1, errors.New(`server error`)
	}

	return userID, nil

}
