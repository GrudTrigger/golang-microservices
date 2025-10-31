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
