package api

import (
	"emreddit/app"
	"emreddit/db"
	"emreddit/logger"
	"emreddit/validator"

	"github.com/gofiber/fiber/v2"
)

type PostPayload struct {
	Description string
}

func mapPostPayloadToDbPost(Payload *PostPayload, dbPost *db.Post, id string) {
	dbPost.Description = Payload.Description
	dbPost.OwnerID = id
}

func readPost(c *fiber.Ctx) error {

	id := c.Params("id")

	err := validator.ValidateUUID(id) //validating id

	if err != nil {

		logger.Error("invalid req", err)

		return c.Status(400).JSON(err.Error())
	}

	post, err := app.ReadPost(id)
	if err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).JSON(post)
}

func createPost(c *fiber.Ctx) error {

	var tokenString = c.Get("X-Auth-Token")

	SessionTokens, err := app.DecryptToken(tokenString)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var PostPayload PostPayload
	err = c.BodyParser(&PostPayload)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	logger.Info(SessionTokens.AccessToken)

	userID, err := app.ParseJWT(SessionTokens.AccessToken)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var dbPost db.Post
	mapPostPayloadToDbPost(&PostPayload, &dbPost, userID)
	dbPost.OwnerID = userID

	err = app.CreatePost(&dbPost)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON(&dbPost)
}

func deletePost(c *fiber.Ctx) error {

	id := c.Params("id") // getting id from params

	err := validator.ValidateUUID(id) //validating id

	if err != nil {

		logger.Error("invalid req", err)
		return c.Status(400).JSON(err.Error())

	}

	_, err = app.DeletePost(id) // for delete api

	if err != nil {
		logger.Error(err.Error())
		return c.Status(400).JSON(err.Error())
	}

	return c.JSON(id)

}

func init() {

	PostApi.Post("/", createPost)
	PostApi.Get(":id", readPost)
	PostApi.Delete(":id", deletePost)
	logger.Info("SUCCESS")
}
