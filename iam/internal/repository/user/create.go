package user

import (
	"context"

	"github.com/rocket-crm/iam/internal/model"
)

func (r *repository) Create(ctx context.Context, user model.User) (string, error) {
	var userId string
	row := r.db.QueryRow(ctx, "INSERT INTO users(login, password, email, notification_method) VALUE($1, $2, $3,$4) RETURNING id", user.Login, user.Password, user.Email, user.NotificationMethod)
	err := row.Scan(&userId)
	if err != nil {
		return "", err
	}
	return userId, nil
}
