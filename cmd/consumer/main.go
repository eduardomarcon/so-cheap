package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"so-cheap/internal/config"
)

type Order struct {
	ID uint64 `json:"id"`
}

func main() {
	config.LoadEnvs()
	conn, err := amqp.Dial(config.GetAMQP().URL)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	db, err := sql.Open("postgres", config.GetDB().URL)
	if err != nil {
		failOnError(err, "failed to open db connection")
	}
	err = db.Ping()
	if err != nil {
		failOnError(err, "failed to ping db connection")
	}

	q, err := ch.QueueDeclare(
		"transport",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			order := Order{}
			err = json.Unmarshal(d.Body, &order)
			failOnError(err, "Failed to convert body to json")

			log.Println(order)
			if _, err = db.Exec("update orders set status = $2 where id = $1", order.ID, 3); err != nil {
				failOnError(err, "Failed to update the status order")
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
