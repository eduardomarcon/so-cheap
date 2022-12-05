package usecase

import (
	"so-cheap/internal/order/database"
	"so-cheap/internal/order/entity"
)

func UpdateOrder(order entity.Order) error {
	orderRepository, err := database.NewOrderRepository()
	if err != nil {
		return err
	}
	if _, err := orderRepository.FindOne(order.ID); err != nil {
		return err
	}
	err = orderRepository.Update(order)
	if err != nil {
		return err
	}
	return nil
}
