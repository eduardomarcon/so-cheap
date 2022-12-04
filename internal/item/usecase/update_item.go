package usecase

import (
	"so-cheap/internal/item/database"
	"so-cheap/internal/item/entity"
)

func UpdateItem(item entity.Item) error {
	itemRepository, err := database.NewItemRepository()
	if err != nil {
		return err
	}
	if _, err := itemRepository.FindOne(item.ID); err != nil {
		return err
	}
	err = itemRepository.Update(item)
	if err != nil {
		return err
	}
	return nil
}
