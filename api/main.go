package api

import (
	"emreddit/logger"

	"github.com/gofiber/fiber/v2"
)

var App *fiber.App

func init() {

	App = fiber.New()

	UserInit()

	err := App.Listen(":3000")

	if err != nil {
		logger.Fatal("err", err)
	}
}
