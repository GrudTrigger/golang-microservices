package user

import (
	"context"

	"github.com/rocket-crm/iam/internal/model"
)

func (r *repository) GetUserByUuid(ctx context.Context, userUuid string) (model.User, error) {
	var user model.User

	row := r.db.QueryRow(ctx, "SELECT id, login, password, email, notification_methods FROM users WHERE id=$1", userUuid)
	err := row.Scan(&user)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
