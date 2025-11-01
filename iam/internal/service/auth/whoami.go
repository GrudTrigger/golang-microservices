package auth

import (
	"context"

	"github.com/rocket-crm/iam/internal/model"
)

func (s *service) Whoami(ctx context.Context, sessionUuid string) (model.UserSessionData, error) {
	userData, err := s.session.Whoami(ctx, sessionUuid)
	if err != nil {
		return model.UserSessionData{}, err
	}
	return userData, nil
} 