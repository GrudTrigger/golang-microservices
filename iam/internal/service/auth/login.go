package auth

import (
	"context"

	"github.com/rocket-crm/iam/internal/model"
)

func (s *service) Login(ctx context.Context, data model.LoginUser) (string, error) {
	user, err := s.userDb.Login(ctx, data.Login)
	if err != nil {
		return "", err
	}
	if user.Password != data.Password {
		return "", model.AuthErr
	}
	sessionData := model.UserSessionData{
		UserUuid: user.Id,
		Login: user.Login,
		Email: user.Email,
	}
	session, err := s.session.Create(ctx, sessionData)
	if err != nil {
		return "", err
	}
	return session, nil
}