package main

import (
	"github.com/gofiber/fiber/v2"
	"so-cheap/internal/config"
	configexternal "so-cheap/pkg/config"
	"so-cheap/pkg/route"
	"so-cheap/pkg/util"
)

func main() {
	config.LoadEnvs()

	fiberConfig := configexternal.FiberConfig()
	app := fiber.New(fiberConfig)

	route.Routes(app)
	util.StartTimers()
	util.StartServerWithGracefulShutdown(app)
}
