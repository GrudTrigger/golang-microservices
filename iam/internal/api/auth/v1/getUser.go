package v1

import (
	"context"

	authV1 "github.com/rocker-crm/shared/pkg/proto/auth/v1"
	"github.com/rocket-crm/iam/internal/converter"
)

func (a *api) GetUser(ctx context.Context, req *authV1.GetUserRequest) (*authV1.GetUserResponse, error) {
	user, err := a.authService.GetUser(ctx, req.UserUuid)
	if err != nil {
		return &authV1.GetUserResponse{}, err
	}
	return &authV1.GetUserResponse{UserUuid: user.UserUuid, Email: user.Email, Login: user.Login, NotificationMethods: converter.ConverterNotificationToProto(user.NotificationMethod)}, nil
}
