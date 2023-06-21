package managers

import (
	"context"
	"soft-weater/internal/database/drivers"
	"soft-weater/internal/models"
)

type User interface {
	Create(ctx context.Context, req models.User) error
}

type UserManager struct {
	userRepository drivers.UserRepository
}

func NewUser(ds drivers.Datastore) User {
	return &UserManager{
		userRepository: ds.UserRepository(),
	}
}

func (s *UserManager) Create(ctx context.Context, req models.User) error {
	return s.userRepository.Create(ctx, req)
}
