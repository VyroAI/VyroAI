package dashboard

import (
	"context"
	"github.com/vyroai/VyroAI/internal/domain/models"
)

//func (d *Dashboard) GetSelf(ctx context.Context, userID int64) (*models.User, error) {
//	user, err := d.userRepo.GetUserByID(ctx, userID)
//	if err != nil {
//		return nil, err
//	}
//	return user, err
//}

func (ds *Service) GetProfile(ctx context.Context, userID int64) (*models.Profile, error) {
	profile, err := ds.profileRepo.GetProfileWithChat(ctx, userID)

	if err != nil {
		return nil, err
	}
	return profile, nil
}
