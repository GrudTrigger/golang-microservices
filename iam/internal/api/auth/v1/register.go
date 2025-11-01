package v1

import (
	"context"

	authV1 "github.com/rocker-crm/shared/pkg/proto/auth/v1"
)

func (a *api) Register(context.Context, *authV1.RegisterRequest) (*authV1.RegisterResponse, error) {
	a.authService.Register()
}