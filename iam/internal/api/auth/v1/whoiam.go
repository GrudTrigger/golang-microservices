package v1

import (
	"context"

	authV1 "github.com/rocker-crm/shared/pkg/proto/auth/v1"
)

func (a *api) Whoami(ctx context.Context, req *authV1.WhoamiRequest) (*authV1.WhoamiResponse, error) {
	userSessionData, err := a.authService.Whoami(ctx, req.SessionUuid)
	if err != nil {
		return &authV1.WhoamiResponse{}, err
	}
	return &authV1.WhoamiResponse{Login: userSessionData.Login, Email: userSessionData.Email, UserUuid: userSessionData.UserUuid}, nil
}
