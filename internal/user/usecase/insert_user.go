package usecase

import (
	"so-cheap/internal/user/database"
	"so-cheap/internal/user/entity"
	"so-cheap/pkg/util"
)

func InsertUser(user entity.User) (int, error) {
	userRepository, err := database.NewUserRepository()
	if err != nil {
		return 0, err
	}
	hashedPass, err := util.Hash(user.Password)
	if err != nil {
		return 0, err
	}
	user.Password = string(hashedPass)
	id, err := userRepository.Insert(user)
	if err != nil {
		return 0, err
	}
	return id, nil
}
