package api

import (
	"emreddit/config"
	"emreddit/logger"

	"github.com/gofiber/fiber/v2"
)

var (
	App     *fiber.App   = fiber.New()
	UserApi fiber.Router = App.Group("/user")
	AuthApi fiber.Router = App.Group("/auth")
)

func ListenPort() {

	if err := App.Listen(config.ListenPort); err != nil {
		logger.Fatal("Error: <?>", err)
	}
}
