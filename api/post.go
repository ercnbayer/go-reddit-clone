package api

import (
	"emreddit/app"
	"emreddit/db"
	"emreddit/logger"
	"emreddit/validator"

	"github.com/gofiber/fiber/v2"
)

type PostPayload struct {
	Description string `validate:"required"`
}

func mapPostPayloadToDbPost(Payload *PostPayload, dbPost *db.Post, id string) {
	dbPost.Description = Payload.Description
	dbPost.OwnerID = id
}

func patchUpdatePost(payload *PostPayload, dbPost *db.Post, id string) {

	dbPost.Description = payload.Description
	dbPost.ID = id
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
		logger.Info("ERR:", err)
		return c.Status(400).JSON(err.Error())
	}

	err = validator.Validate.Struct(&PostPayload)

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
	//dbPost.OwnerID = "b7c9b56a-1b80-47d5-b471-bded8b6dc8a5"

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

func updatePost(c *fiber.Ctx) error {

	id := c.Params("id") //getting id from params

	err := validator.ValidateUUID(id)

	if err != nil {

		logger.Error("invalid req", err)
		return c.Status(400).JSON(err.Error())

	}

	var tokenString = c.Get("X-Auth-Token")

	SessionTokens, err := app.DecryptToken(tokenString)

	if err != nil {

		return c.Status(400).JSON(err.Error())
	}

	userID, err := app.ParseJWT(SessionTokens.AccessToken)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var post PostPayload // creating instance

	if err := c.BodyParser(&post); err != nil { // check if err from body

		logger.Error(" Body Parse error = ", err, post)

		return c.Status(404).JSON(err.Error())
	}

	if err := validator.Validate.Struct(&post); err != nil { //validating updated values

		logger.Error("validator err= ", err)

		return c.Status(404).JSON(err.Error())
	}

	var dbPost db.Post

	// maping user to dbUser

	patchUpdatePost(&post, &dbPost, id)

	if err := app.UpdatePost(&dbPost, userID); err != nil {

		logger.Error("Update ERR:", err)

		return err
	}

	return c.Status(200).JSON(dbPost)
}

func init() {

	PostApi.Post("/", createPost)
	PostApi.Get(":id", readPost)
	PostApi.Delete(":id", deletePost)
	PostApi.Patch(":id", updatePost)
	logger.Info("SUCCESS")
}
