package usecase

import (
	"so-cheap/internal/item/database"
	"so-cheap/internal/item/entity"
)

func InsertItem(item entity.Item) (int, error) {
	itemRepository, err := database.NewItemRepository()
	if err != nil {
		return 0, err
	}
	id, err := itemRepository.Insert(item)
	if err != nil {
		return 0, err
	}
	return id, nil
}
