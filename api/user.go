package api

import (
	"emreddit/app"
	db "emreddit/db"
	"emreddit/logger"
	"emreddit/validator"

	"github.com/gofiber/fiber/v2"
)

type UserUpdatePayload struct { //payload for updating user
	ID       string
	Name     string `validate:"omitempty,required"` /* if struct field not empty validate */
	Password string `validate:"omitempty,required"`
	Email    string `validate:"omitempty,required,email"`
}

type UserProfile struct {
	ID    string
	Name  string
	Email string
}

func mapUserUpdatePayloadToDbUser(user *UserUpdatePayload, dbUser *db.UserEntity) {

	dbUser.ID = user.ID
	dbUser.Name = user.Name
	dbUser.Email = user.Email
	dbUser.Password = user.Password
}

func mapUserEntityToUserProfile(dbUser *db.UserEntity) UserProfile {

	var userProfile UserProfile
	userProfile.ID = dbUser.ID
	userProfile.Name = dbUser.Name
	userProfile.Email = dbUser.Email
	return userProfile
}

func getUser(c *fiber.Ctx) error {

	//getting single user

	// Get the ID from the URL parameter
	id := c.Params("id")

	err := validator.ValidateUUID(id) //validating id

	if err != nil {

		logger.Error("invalid req", err)

		return c.Status(400).JSON(err.Error())
	}
	user, err := app.GetUser(id)

	if err != nil {
		return c.Status(404).JSON(err)
	}

	// calling app function
	return c.JSON(mapUserEntityToUserProfile(user)) // return it as json

}

func deleteUser(c *fiber.Ctx) error {

	id := c.Params("id") // getting id from params

	err := validator.ValidateUUID(id) //validating id

	if err != nil {

		logger.Error("invalid req", err)
		return c.Status(400).JSON(err.Error())

	}

	_, err = app.DeleteUser(id) // for delete api

	if err != nil {
		logger.Error(err.Error())
		return c.Status(400).JSON(err.Error())
	}

	return c.JSON(id)

}

func updateUser(c *fiber.Ctx) error {

	id := c.Params("id") //getting id from params

	err := validator.ValidateUUID(id)

	if err != nil {

		logger.Error("invalid req", err)
		return c.Status(400).JSON(err.Error())

	}

	var user UserUpdatePayload // creating instance

	if err := c.BodyParser(&user); err != nil { // check if err from body

		logger.Error(" Body Parse error = ", err, user)

		return c.Status(404).JSON(err.Error())
	}
	user.ID = id

	if err := validator.Validate.Struct(&user); err != nil { //validating updated values

		logger.Error("validator err= ", err)

		return c.Status(404).JSON(err.Error())
	}

	var dbUser db.UserEntity

	// maping user to dbUser

	mapUserUpdatePayloadToDbUser(&user, &dbUser)

	if err := app.UpdateUser(&dbUser); err != nil {

		logger.Error("Update ERR:", err)

		return err
	}

	return c.Status(200).JSON(user)
}

func listUsers(c *fiber.Ctx) error {

	people, err := app.GetUsers() //getting all userslist

	if err != nil {

		logger.Error(" no user found", err.Error())
		return c.Status(404).JSON(err.Error())
	}

	return c.JSON(people)
}

func init() { //creating routes
	// grouping rotues

	UserApi.Get(":id", getUser) // get single user

	UserApi.Get("/", listUsers) //list all users

	UserApi.Patch(":id", updateUser) //update user

	UserApi.Delete(":id", deleteUser) //delete user

}
