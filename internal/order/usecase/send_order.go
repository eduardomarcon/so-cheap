package usecase

import (
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"so-cheap/cmd/producer"
	"so-cheap/internal/config"
	"so-cheap/internal/order/database"
	"so-cheap/internal/order/entity"
	"time"
)

type OrderMessage struct {
	ID uint64
}

func TimerSendPayedOrders() error {
	ticker := time.NewTicker(5 * time.Second)
	for {
		<-ticker.C
		err := sendOrdersPayed()
		if err != nil {
			return err
		}
	}
}

func sendOrdersPayed() error {
	orderRepository, err := database.NewOrderRepository()
	if err != nil {
		return err
	}
	fmt.Println("finding orders")
	orders, err := orderRepository.FindAllByStatus(entity.Payed)
	if err != nil {
		return err
	}
	if len(orders) <= 0 {
		return nil
	}
	conn, err := amqp.Dial(config.GetAMQP().URL)
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	for _, order := range orders {
		ordersJson, err := json.Marshal(OrderMessage{
			ID: order.ID,
		})
		if err != nil {
			return err
		}
		err = producer.Publish(ch, "transport", string(ordersJson))
		if err != nil {
			return err
		}
	}
	return nil
}
