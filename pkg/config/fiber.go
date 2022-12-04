package configexternal

import (
	"github.com/gofiber/fiber/v2"
	"so-cheap/internal/config"
	"time"
)

func FiberConfig() fiber.Config {
	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(config.GetServer().TimeOut),
	}
}
