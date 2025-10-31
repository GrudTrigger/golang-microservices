package auth

import "github.com/rocket-crm/iam/internal/repository"

type service struct {
	userDb  repository.UserRepository
	session repository.SessionRepository
}

func NewService(userDb repository.UserRepository, session repository.SessionRepository) *service {
	return &service{userDb: userDb, session: session}
}
