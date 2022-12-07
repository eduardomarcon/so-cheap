package usecase

import (
	"so-cheap/internal/user/database"
	"so-cheap/internal/user/entity"
)

func AuthenticateUser(email string, password string) (entity.User, error) {
	userRepository, err := database.NewUserRepository()
	if err != nil {
		return entity.User{}, err
	}
	user, err := userRepository.FindOneByEmailAndPassword(email, password)
	if err != nil {
		return user, err
	}
	return user, nil
}
