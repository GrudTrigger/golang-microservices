package converter

import (
	authV1 "github.com/rocker-crm/shared/pkg/proto/auth/v1"
	"github.com/rocket-crm/iam/internal/model"
)

func ConvertUserToModel(data *authV1.RegisterRequest) model.RegisterUserRequest {
	return model.RegisterUserRequest{
		Login:              data.Login,
		Email:              data.Email,
		Password:           data.Password,
		NotificationMethod: ConvertNotificationToModel(data.NotificationMethods),
	}
}

func ConvertNotificationToModel(n []*authV1.NotificationMethod) []model.Notification {
	nots := make([]model.Notification, 0, len(n))
	for _, v := range n {
		not := model.Notification{ProviderName: v.ProviderName, Target: v.Target}
		nots = append(nots, not)
	}
	return nots
}

func ConvertLoginToModel(data *authV1.LoginRequest) model.LoginUser {
	return model.LoginUser{
		Login:    data.Login,
		Password: data.Password,
	}
}

func ConverterNotificationToProto(not []model.Notification) []*authV1.NotificationMethod {
	result := make([]*authV1.NotificationMethod, 0, len(not))

	for _, v := range not {
		newNot := &authV1.NotificationMethod{ProviderName: v.ProviderName, Target: v.Target}
		result = append(result, newNot)
	}
	return result
}
