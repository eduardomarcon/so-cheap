package usecase

import (
	"so-cheap/internal/user/database"
	"so-cheap/internal/user/entity"
)

func InsertUser(user entity.User) (int, error) {
	userRepository, err := database.NewUserRepository()
	if err != nil {
		return 0, err
	}
	id, err := userRepository.Insert(user)
	if err != nil {
		return 0, err
	}
	return id, nil
}
