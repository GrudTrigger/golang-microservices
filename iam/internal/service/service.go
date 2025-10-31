package service

import (
	"context"

	"github.com/rocket-crm/iam/internal/model"
)

type AuthService interface {
	Register(ctx context.Context, user model.User) (string, error)
	Login(ctx context.Context, data model.LoginUser) (string, error)
	Whoami(ctx context.Context, sessionUuid string) (model.UserSessionData, error)
	GetUser(ctx context.Context, userUuid string) (model.GetUserResponse, error)
}
