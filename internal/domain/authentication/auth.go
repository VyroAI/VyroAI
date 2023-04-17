package authentication

import (
	"context"
	"github.com/vyroai/VyroAI/commons/errors"
	"github.com/vyroai/VyroAI/commons/otel"
	repo2 "github.com/vyroai/VyroAI/internal/domain/repo"
	"github.com/vyroai/VyroAI/internal/infra/authentication/authProviderRepository"
	"go.opentelemetry.io/otel/trace"
)

type Authentication interface {
	Login(ctx context.Context, email, password string) (int64, int32, error)
	Register(ctx context.Context, username, email, password string) (int64, int32, error)

	GenerateDiscordAuthUrl(ctx context.Context, action string) (string, error)
	DiscordProviderLogin(ctx context.Context, code, state string) (int64, int32, error)
	DiscordProviderRegister(ctx context.Context, code, state string) (int64, int32, error)
}

type AuthService struct {
	tracer       trace.Tracer
	userRepo     repo2.UserRepo
	bcryptRepo   repo2.BcryptRepo
	authProvider *authProviderRepository.AuthProvider
}

func NewAuthService(userRepo repo2.UserRepo, bcryptRepo repo2.BcryptRepo, authProvider *authProviderRepository.AuthProvider) Authentication {
	tracer := otel.InitTracing("authenticationService", "0.1.0")

	return &AuthService{
		tracer:       tracer.NewTracer(),
		userRepo:     userRepo,
		bcryptRepo:   bcryptRepo,
		authProvider: authProvider,
	}
}

func (as *AuthService) Login(ctx context.Context, email, password string) (int64, int32, error) {
	ctx, span := as.tracer.Start(ctx, "login")
	defer span.End()

	userResult, err := as.userRepo.GetUserByEmail(ctx, email)
	if userResult == nil || err != nil {
		return -1, -1, err
	}

	err = as.bcryptRepo.CompareHashAndPassword(ctx, userResult.Password, password)
	if err != nil {
		return -1, -1, errors.ErrInvalid.Error()
	}

	return userResult.Id, userResult.Permission, nil
}

func (as *AuthService) Register(ctx context.Context, username, email, password string) (int64, int32, error) {
	ctx, span := as.tracer.Start(ctx, "register")
	defer span.End()

	userByEmail, err := as.userRepo.GetUserByEmail(ctx, email)
	if errors.GetType(err) != errors.ErrNotFound && err != nil {
		return -1, -1, err
	}
	if userByEmail != nil {
		return -1, -1, errors.ErrExist.New("Email already exist")
	}

	userByUsername, err := as.userRepo.GetUserByUsername(ctx, username)
	if errors.GetType(err) != errors.ErrNotFound && err != nil {
		return -1, -1, err
	}
	if userByUsername != nil {
		return -1, -1, errors.ErrExist.New("Username already exist")
	}

	hashedPassword, err := as.bcryptRepo.GenerateFromPassword(ctx, password)
	if err != nil {
		return -1, -1, err
	}

	userID, err := as.userRepo.CreateUser(ctx, username, email, hashedPassword)
	if err != nil {
		return -1, -1, err
	}

	return userID, 1, nil

}
