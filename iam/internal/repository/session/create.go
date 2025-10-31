package session

import (
	"context"

	"github.com/google/uuid"
	"github.com/rocket-crm/iam/internal/model"
)

func (r *repository) Create(ctx context.Context, userData model.UserSessionData) (string, error) {
	newUuid := uuid.NewString()
	sessionKey := r.GetSessionKey(newUuid)
	err := r.cache.HashSet(ctx, sessionKey, userData)
	if err != nil {
		return "", err
	}
	return newUuid, nil
}
