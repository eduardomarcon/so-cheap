package usecase

import (
	"so-cheap/internal/user/database"
	"so-cheap/internal/user/entity"
	"so-cheap/pkg/util"
)

func AuthenticateUser(email string, password string) (entity.User, error) {
	userRepository, err := database.NewUserRepository()
	if err != nil {
		return entity.User{}, err
	}
	user, err := userRepository.FindOneByEmail(email)
	if err != nil {
		return user, err
	}
	err = util.Check(user.Password, password)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}
