package usecase

import (
	"so-cheap/internal/item/database"
	"so-cheap/internal/item/entity"
)

func GetAllItens() ([]entity.Item, error) {
	itemRepository, err := database.NewItemRepository()
	if err != nil {
		return nil, err
	}
	itens, err := itemRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return itens, nil
}

func GetOneItem(id int64) (entity.Item, error) {
	itemRepository, err := database.NewItemRepository()
	if err != nil {
		return entity.Item{}, err
	}
	item, err := itemRepository.FindOne(id)
	if err != nil {
		return item, err
	}
	return item, nil
}
