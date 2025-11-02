package v1

import (
	"context"

	authV1 "github.com/rocker-crm/shared/pkg/proto/auth/v1"
	"github.com/rocket-crm/iam/internal/converter"
)

func (a *api) Register(ctx context.Context, data *authV1.RegisterRequest) (*authV1.RegisterResponse, error) {
	userUuid, err := a.authService.Register(ctx, converter.ConvertUserToModel(data))
	if err != nil {
		return nil, err
	}
	return &authV1.RegisterResponse{UserUuid: userUuid}, nil
}
