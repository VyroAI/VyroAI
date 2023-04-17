package dashboard

import (
	"context"
	"github.com/vyroai/VyroAI/internal/domain/models"
	"github.com/vyroai/VyroAI/internal/domain/repo"
)

type Dashboard interface {
	GetProfile(ctx context.Context) (*models.Profile, error)
}

type Service struct {
	userRepo    repo.UserRepo
	profileRepo repo.Profile
}

func NewDashboardService(userRepo repo.UserRepo, profileRepo repo.Profile) Dashboard {
	return &Service{
		userRepo:    userRepo,
		profileRepo: profileRepo,
	}
}
