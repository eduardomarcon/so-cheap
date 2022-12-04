package usecase

import (
	"so-cheap/internal/item/database"
)

func DeleteItem(id int64) error {
	itemRepository, err := database.NewItemRepository()
	if err != nil {
		return err
	}
	err = itemRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
