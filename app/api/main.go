package api

import (
	"emreddit/logger"

	"github.com/gofiber/fiber/v2"
)

var (
	App     *fiber.App   = fiber.New()
	UserApi fiber.Router = App.Group("/user")
)

func ListenPort() {

	err := App.Listen(":3000")

	if err != nil {
		logger.Fatal("err", err)
	}
}

func init() {

}
