package model

type User struct {
	Id                 string
	Login              string
	Password           string
	Email              string
	NotificationMethod []Notification
}

type Notification struct {
	ProviderName string `json:"provider_name"`
	Target       string `json:"target"`
}

type UserSessionData struct {
	UserUuid string `redis:"user_uuid"`
	Login    string `redis:"login"`
	Email    string `redis:"email"`
}

type LoginUser struct {
	Login    string
	Password string
}

type GetUserResponse struct {
	UserUuid           string
	Login              string
	Email              string
	NotificationMethod []Notification
}

type RegisterUserRequest struct {
	Login              string
	Password           string
	Email              string
	NotificationMethod []Notification
}
