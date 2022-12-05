package usecase

import (
	"so-cheap/internal/order/database"
)

func DeleteOrder(id uint64) error {
	orderRepository, err := database.NewOrderRepository()
	if err != nil {
		return err
	}
	err = orderRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
