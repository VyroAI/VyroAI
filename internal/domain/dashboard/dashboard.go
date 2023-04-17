package dashboard

import (
	"github.com/vyroai/VyroAI/internal/domain/repo"
)

type Dashboard struct {
	userRepo    repo.UserRepo
	profileRepo repo.Profile
}

func NewDashboardService(userRepo repo.UserRepo) *Dashboard {
	return &Dashboard{
		userRepo: userRepo,
	}
}
