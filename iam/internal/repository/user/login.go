package user

import (
	"context"

	"github.com/rocket-crm/iam/internal/model"
)

func (r *repository) Login(ctx context.Context, login string) (model.User, error) {
	var user model.User

	row := r.db.QueryRow(ctx, "SELECT id, login, password, email, notification_methods FROM users WHERE login=$1", login)
	err := row.Scan(&user)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
