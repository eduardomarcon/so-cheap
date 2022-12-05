package usecase

import (
	"so-cheap/internal/order/database"
	"so-cheap/internal/order/entity"
)

func InsertOrder(order entity.Order) (int, error) {
	orderRepository, err := database.NewOrderRepository()
	if err != nil {
		return 0, err
	}
	id, err := orderRepository.Insert(order)
	if err != nil {
		return 0, err
	}
	return id, nil
}
