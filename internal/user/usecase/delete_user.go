package usecase

import (
	"so-cheap/internal/user/database"
)

func DeleteUser(id int64) error {
	itemRepository, err := database.NewUserRepository()
	if err != nil {
		return err
	}
	err = itemRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
