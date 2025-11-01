package auth

import (
	"context"

	"github.com/rocket-crm/iam/internal/model"
)

func (s *service) GetUser(ctx context.Context, userUuid string) (model.GetUserResponse, error) {
	user, err := s.userDb.GetUserByUuid(ctx, userUuid)
	if err != nil {
		return model.GetUserResponse{}, err
	}
	return model.GetUserResponse{UserUuid: user.Id, Login: user.Login, Email: user.Email, NotificationMethod: user.NotificationMethod}, nil
}