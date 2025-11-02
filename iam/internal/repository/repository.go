package repository

import (
	"context"

	"github.com/rocket-crm/iam/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, user model.RegisterUserRequest) (string, error)
	GetUserByUuid(ctx context.Context, userUuid string) (model.User, error)
	Login(ctx context.Context, login string) (model.User, error)
}

type SessionRepository interface {
	Create(ctx context.Context, userData model.UserSessionData) (string, error)
	GetSessionKey(newUuid string) string
	Whoami(ctx context.Context, sessionUuid string) (model.UserSessionData, error)
}
