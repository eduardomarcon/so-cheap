package usecase

import (
	"so-cheap/internal/user/database"
	"so-cheap/internal/user/entity"
)

func GetAllUsers() ([]entity.User, error) {
	userRepository, err := database.NewUserRepository()
	if err != nil {
		return nil, err
	}
	itens, err := userRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return itens, nil
}

func GetOneUser(id int64) (entity.User, error) {
	userRepository, err := database.NewUserRepository()
	if err != nil {
		return entity.User{}, err
	}
	user, err := userRepository.FindOne(id)
	if err != nil {
		return user, err
	}
	return user, nil
}
