package api

import (
	db "emreddit/db"
	"emreddit/logger"
	"emreddit/validator"

	"github.com/gofiber/fiber/v2"
)

func getSingleUser(c *fiber.Ctx) error {

	//getting single user

	// Get the ID from the URL parameter
	id := c.Params("id") // getting id from params

	err := validator.ValidateID(id) //validating id

	if err != nil {

		logger.Error("invalid req", err)

		return c.Status(400).JSON(err.Error())
	}

	var user db.User
	err = db.ReadUser(id, &user) //readingUser

	if err != nil { //check if err is not null

		logger.Error(err.Error(), err)

		return c.Status(404).JSON(err.Error())
	}

	return c.JSON(user) // return it as json

}
func deleteUser(c *fiber.Ctx) error {

	id := c.Params("id") // getting id from params

	err := validator.ValidateID(id) //validating id

	if err != nil {

		logger.Error("invalid req", err)
		return c.Status(400).JSON(err.Error())

	}

	_, err = db.DeleteUser(id) // for delete api

	if err != nil {
		logger.Error(err.Error())
		return c.Status(400).JSON(err.Error())
	}

	return c.JSON(id)

}

func updateUser(c *fiber.Ctx) error {

	id := c.Params("id") //getting id from params

	err := validator.ValidateID(id) //validating id

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

	if err := validator.ValidateUpdatedStruct(&user); err != nil { //validating updated values

		logger.Error("validator err= ", err)

		return c.Status(404).JSON(err.Error())
	}

	var dbUser db.User

	mapUserUpdatePayloadToDbUser(&user, &dbUser) // maping user to dbUser

	if err := db.PatchUpdateUser(&dbUser); err != nil { //sending it to db

		logger.Error("Update ERR:", err)

		return err
	}

	return c.Status(200).JSON(user)
}

func registerUser(c *fiber.Ctx) error { // for registering user

	var user UserPayload

	if err := c.BodyParser(&user); err != nil { //parsing body

		logger.Error("BodyParsing err:", user)
		return c.Status(400).JSON(err.Error())
	}

	if err := validator.ValidateStruct(&user); err != nil { //validating struct
		return c.Status(400).JSON(err.Error())
	}

	var dbUser db.User

	mapUserPayloadToDbUserCreate(&user, &dbUser) //maping to db obj

	if err := db.InsertUser(&dbUser); err != nil { //Inserting user

		logger.Error("User Insert Err", err)

		return c.Status(400).JSON(err.Error())

	}

	return c.Status(200).JSON(dbUser)

}

func userLogin(c *fiber.Ctx) error {

	var user UserLoginPayload

	if err := c.BodyParser(&user); err != nil { //parsing body

		logger.Error("BodyParsing err:", user)
		return c.Status(400).JSON(err.Error())
	}

	if err := validator.ValidateStruct(&user); err != nil { //validating struct

		return c.Status(400).JSON(err.Error())
	}

	var dbUser db.User

	mapUserLoginPayloadToDbUser(&user, &dbUser) //maping user to db obj

	if err := db.LoginUser(&dbUser); err != nil { // sending it to db

		logger.Error("login err", err)

		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON(dbUser.ID)

}

func listUsers(c *fiber.Ctx) error {

	people, err := db.GetUsers() //getting all userslist

	if err != nil {

		logger.Error(" no user found", err.Error())
		return c.Status(404).JSON(err.Error())
	}

	return c.JSON(people)
}

func UserInit() {
	userApi := App.Group("/user") // grouping rotues

	userApi.Post("/", registerUser) // creating user

	userApi.Get(":id", getSingleUser) // get single user

	userApi.Get("/", listUsers) //list all users

	userApi.Patch(":id", updateUser) //update user

	userApi.Delete(":id", deleteUser) //delete user
	userApi.Post("/login", userLogin) //USER LOGIN
}
