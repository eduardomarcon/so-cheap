package usecase

import (
	"so-cheap/internal/user/database"
	"so-cheap/internal/user/entity"
)

func UpdateUser(user entity.User) error {
	userRepository, err := database.NewUserRepository()
	if err != nil {
		return err
	}
	if _, err := userRepository.FindOne(user.ID); err != nil {
		return err
	}
	err = userRepository.Update(user)
	if err != nil {
		return err
	}
	return nil
}
