package user

import (
	"context"
	"encoding/json"

	"github.com/rocket-crm/iam/internal/model"
)

func (r *repository) Create(ctx context.Context, user model.RegisterUserRequest) (string, error) {
	var userId string
	jsonNotificationMethod, _ := json.Marshal(user.NotificationMethod)
	row := r.db.QueryRow(ctx, "INSERT INTO users(login, password, email, notification_methods) VALUES ($1, $2, $3, $4) RETURNING id", user.Login, user.Password, user.Email, jsonNotificationMethod)
	err := row.Scan(&userId)
	if err != nil {
		return "", err
	}
	return userId, nil
}
