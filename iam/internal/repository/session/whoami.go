package session

import (
	"context"
	"errors"

	redigo "github.com/gomodule/redigo/redis"
	"github.com/rocket-crm/iam/internal/model"
)

func (r *repository) Whoami(ctx context.Context, sessionUuid string) (model.UserSessionData, error) {
	sessionKey := r.GetSessionKey(sessionUuid)

	values, err := r.cache.HGetAll(ctx, sessionKey)
	if err != nil {
		if errors.Is(err, redigo.ErrNil) {
			return model.UserSessionData{}, model.SessionNotFound
		}
		return model.UserSessionData{}, err
	}
	if len(values) == 0 {
		return model.UserSessionData{}, model.SessionNotFound
	}

	var sessionData model.UserSessionData
	err = redigo.ScanStruct(values, &sessionData)
	if err != nil {
		return model.UserSessionData{}, err
	}
	return sessionData, nil
}
