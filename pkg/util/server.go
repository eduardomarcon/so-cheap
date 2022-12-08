package util

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
	"so-cheap/internal/config"
	"so-cheap/internal/order/usecase"
)

func StartTimers() {
	go func() {
		err := usecase.TimerSendPayedOrders()
		if err != nil {
			log.Fatal(err)
		}
	}()
}

func StartServerWithGracefulShutdown(a *fiber.App) {
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := a.Shutdown(); err != nil {
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	if err := a.Listen(config.GetServer().URL); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}
