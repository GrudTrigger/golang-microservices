package v1

import (
	"context"

	authV1 "github.com/rocker-crm/shared/pkg/proto/auth/v1"
	"github.com/rocket-crm/iam/internal/converter"
)

func (a *api) Login(ctx context.Context, req *authV1.LoginRequest) (*authV1.LoginResponse, error) {
	sessionUuid, err := a.authService.Login(ctx, converter.ConvertLoginToModel(req))
	if err != nil {
		return &authV1.LoginResponse{}, err
	}
	return &authV1.LoginResponse{SessionUuid: sessionUuid}, nil
}
