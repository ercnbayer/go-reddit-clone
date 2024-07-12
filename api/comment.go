package api

import (
	"emreddit/app"
	"emreddit/db"
	"emreddit/logger"
	"emreddit/validator"

	"github.com/gofiber/fiber/v2"
)

type CommentPayload struct {
	Description string `validate:"required"`
}

func mapCommentPayloadToDbComment(Payload *CommentPayload, dbPost *db.Comment, id string) {
	dbPost.Description = Payload.Description
	dbPost.OwnerID = id
}

func readComment(c *fiber.Ctx) error {

	id := c.Params("id")

	err := validator.ValidateUUID(id) //validating id

	if err != nil {

		logger.Error("invalid req", err)

		return c.Status(400).JSON(err.Error())
	}

	post, err := app.ReadComment(id)
	if err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).JSON(post)
}

func createComment(c *fiber.Ctx) error {

	var tokenString = c.Get("X-Auth-Token")

	SessionTokens, err := app.DecryptToken(tokenString)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var CommentPayload CommentPayload
	err = c.BodyParser(&CommentPayload)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	validator.Validate.Struct(&CommentPayload)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	//logger.Info(SessionTokens.AccessToken)

	userID, err := app.ParseJWT(SessionTokens.AccessToken)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var dbComment db.Comment
	mapCommentPayloadToDbComment(&CommentPayload, &dbComment, userID)

	err = app.CreateComment(&dbComment)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON(&dbComment)
}
func deleteComment(c *fiber.Ctx) error {

	id := c.Params("id") // getting id from params

	err := validator.ValidateUUID(id) //validating id

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

	_, err = app.DeleteComment(id, userID) // for delete api

	if err != nil {
		logger.Error(err.Error())
		return c.Status(400).JSON(err.Error())
	}

	return c.JSON(id)

}

func updateComment(c *fiber.Ctx) error {

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

	var comment CommentPayload // creating instance

	if err := c.BodyParser(&comment); err != nil { // check if err from body

		logger.Error(" Body Parse error = ", err, comment)

		return c.Status(404).JSON(err.Error())
	}

	if err := validator.Validate.Struct(&comment); err != nil { //validating updated values

		logger.Error("validator err= ", err)

		return c.Status(404).JSON(err.Error())
	}

	var dbComment db.Comment

	dbComment.ID = id

	// maping user to dbComment

	mapCommentPayloadToDbComment(&comment, &dbComment, id)

	if err := app.UpdateComment(&dbComment, userID); err != nil {

		logger.Error("Update ERR:", err)

		return err
	}

	return c.Status(200).JSON(comment)
}

func init() {

	CommentApi.Post("/", createComment)
	CommentApi.Get(":id", readComment)
	CommentApi.Delete(":id", deleteComment)
	CommentApi.Patch(":id", updateComment)

	logger.Info("SUCCESS")
}
