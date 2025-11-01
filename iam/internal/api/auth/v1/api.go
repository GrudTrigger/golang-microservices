package v1

import (
	authV1 "github.com/rocker-crm/shared/pkg/proto/auth/v1"
	"github.com/rocket-crm/iam/internal/service"
)

type api struct {
	authV1.UnimplementedAuthServiceServer
	authService service.AuthService
}

func NewApi(authService service.AuthService) *api {
	return &api{authService: authService}
}