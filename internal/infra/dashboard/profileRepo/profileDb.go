package profileRepo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/vyroai/VyroAI/commons/errors"
	"github.com/vyroai/VyroAI/internal/domain/models"
	"go.opentelemetry.io/otel/codes"
)

func (pr *ProfileRepository) GetProfileWithChat(ctx context.Context, userID int64) (*models.Profile, error) {
	ctx, span := pr.tracer.Start(ctx, "get-profile-and-chat")
	defer span.End()

	profile, err := pr.database.GetProfileAndChats(ctx, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrNotFound.Wrap(err, "get profile and chat")
		} else {
			span.RecordError(err)
			span.SetStatus(codes.Error, fmt.Sprintf("unknown error,  %+v\n", err))
			pr.logger.Error(fmt.Sprintf("unknown error,  %+v\n", err))
			return nil, err
		}
	}
	return profileToModel(profile), nil
}
