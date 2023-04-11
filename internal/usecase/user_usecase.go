package usecase

import (
	"github.com/TotrazOrstO/Imapro/internal/repository"
	"github.com/TotrazOrstO/Imapro/pkg/database"
)

type UserUsecase interface {
	CreateUser(user *database.User) error
	GetUserByID(id int) (*database.User, error)
}

type UserUsecaseImpl struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) *UserUsecaseImpl {
	return &UserUsecaseImpl{userRepository: userRepository}
}

func (u *UserUsecaseImpl) CreateUser(user *database.User) error {
	// Реализация бизнес-логики создания пользователя
	return u.userRepository.CreateUser(user)
}

func (u *UserUsecaseImpl) GetUserByID(id int) (*database.User, error) {
	// Реализация бизнес-логики получения пользователя по ID
	return u.userRepository.GetUserByID(id)
}
