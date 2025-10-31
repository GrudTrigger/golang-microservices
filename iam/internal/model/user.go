package model

type User struct {
	Login              string
	Password           string
	Email              string
	NotificationMethod []Notification
}

type Notification struct {
	ProviderName string
	Target       string
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
