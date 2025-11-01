package auth

import (
	"context"

	"github.com/rocket-crm/iam/internal/model"
)

func (s *service) Register(ctx context.Context, user model.User) (string, error) {
	existedUser, err := s.userDb.Login(ctx, user.Login)
	if err != nil {
		return "", err
	}

	if existedUser.Login == user.Login {
		return "", model.AlreadyRegistered
	}

	uuid,err := s.userDb.Create(ctx, user)
	if err != nil {
		return "", err
	}
	return uuid, nil
}