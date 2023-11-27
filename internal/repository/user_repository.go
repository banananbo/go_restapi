package repository

import (
	"github.com/google/uuid"
	"myapi/internal/domain"
)

type UserRepository interface {
	AddUser(user domain.User) (domain.User, error)
	GetUsers() ([]domain.User, error)
}

type userRepo struct {
	users []domain.User
}

func NewUserRepository() UserRepository {
	return &userRepo{}
}

func (ur *userRepo) AddUser(user domain.User) (domain.User, error) {
	user.ID = uuid.New().String()
	ur.users = append(ur.users, user)
	return user, nil
}

func (ur *userRepo) GetUsers() ([]domain.User, error) {
	return ur.users, nil
}
