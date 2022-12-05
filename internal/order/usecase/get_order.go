package usecase

import (
	"so-cheap/internal/order/database"
	"so-cheap/internal/order/entity"
)

func GetAllOrders() ([]entity.Order, error) {
	orderRepository, err := database.NewOrderRepository()
	if err != nil {
		return nil, err
	}
	orders, err := orderRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func GetOneOrder(id uint64) (entity.Order, error) {
	orderRepository, err := database.NewOrderRepository()
	if err != nil {
		return entity.Order{}, err
	}
	order, err := orderRepository.FindOne(id)
	if err != nil {
		return order, err
	}
	return order, nil
}
