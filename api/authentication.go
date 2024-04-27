package api

import (
	"emreddit/app"
	db "emreddit/db"
	"emreddit/logger"
	"emreddit/validator"

	"github.com/gofiber/fiber/v2"
)

type UserPayload struct { //payload for registerUser
	ID       string
	Name     string `validate:"required"`
	Password string `validate:"required"`
	Email    string `validate:"required,email"`
}

type UserLoginPayload struct { // payload for Login User
	Email    string `validate:"required,email"` //validate email
	Password string `validate:"required"`
}

func mapUserLoginPayloadToDbUser(user *UserLoginPayload, dbUser *db.UserEntity) {

	dbUser.Email = user.Email
	dbUser.Password = user.Password

}

func mapUserPayloadToDbUserCreate(user *UserPayload, dbUser *db.UserEntity) {

	dbUser.Name = user.Name
	dbUser.Email = user.Email
	dbUser.Password = user.Password

}

func registerUser(c *fiber.Ctx) error { // for registering user

	var user UserPayload

	if err := c.BodyParser(&user); err != nil { //parsing body

		logger.Error("BodyParsing err <?>", user)
		return c.Status(400).JSON(err.Error())
	}

	if err := validator.Validate.Struct(&user); err != nil { //validating struct
		return c.Status(400).JSON(err.Error())
	}

	var dbUser db.UserEntity

	mapUserPayloadToDbUserCreate(&user, &dbUser) //maping to db obj

	if err := app.RegisterUser(&dbUser); err != nil { //Inserting user

		logger.Error("Error <?>	", err)

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

	if err := validator.Validate.Struct(&user); err != nil { //validating struct

		return c.Status(400).JSON(err.Error())
	}

	var dbUser db.UserEntity

	mapUserLoginPayloadToDbUser(&user, &dbUser) //maping user to db obj

	if err := app.UserLogin(&dbUser); err != nil { // sending it to db

		logger.Error("login err <?>", err)

		return c.Status(404).JSON(err.Error())

	}
	refreshToken, err := app.CreateRefreshToken(dbUser.ID)
	if err != nil {
		return c.Status(401).JSON(err.Error())
	}

	accessToken, err := app.CreateJWT(dbUser.ID)
	if err != nil {
		return c.Status(401).JSON(err.Error())
	}

	userTokens := app.SessionToken{AccessToken: accessToken, RefreshToken: refreshToken}
	bytes, err := app.JSONToBytes(&userTokens)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	encoded_token, err := app.EncryptToken(bytes)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON(encoded_token)

}

func getAccessToken(c *fiber.Ctx) error {

	var tokenString = c.Get("X-Auth-Token", "")
	SessionTokens, err := app.DecryptToken(tokenString)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	userID, err := app.CheckIfTokenValid(SessionTokens.RefreshToken)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	refreshToken, err := app.CreateRefreshToken(userID)
	if err != nil {
		return c.Status(401).JSON(err.Error())
	}
	accessToken, err := app.CreateJWT(userID)

	if err != nil {
		logger.Error("JWT Token Error:<?>", err)
		return c.Status(401).JSON(err.Error())
	}

	userTokens := app.SessionToken{AccessToken: accessToken, RefreshToken: refreshToken}
	bytes, err := app.JSONToBytes(&userTokens)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	encoded_token, err := app.EncryptToken(bytes)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON(encoded_token)

}
func me(c *fiber.Ctx) error {
	var tokenString = c.Get("X-Auth-Token", "null")

	tokens, err := app.DecryptToken(tokenString)

	if err != nil {
		logger.Error("Decrypt Token Error:<?>", err)
		return c.Status(401).JSON(err.Error())
	}

	tokenString = tokens.AccessToken
	id, err := app.ParseJWT(tokenString)

	if err != nil {
		logger.Error("JWT Token Error:<?>", err)
		return c.Status(401).JSON(err.Error())
	}

	user, err := app.GetUser(id)

	if err != nil {
		logger.Error("Get User Error:<?>", err)
		return c.Status(400).JSON(err.Error())
	}
	return c.Status(200).JSON(&user)
}

func init() {

	UserApi.Post("/", registerUser)
	AuthApi.Post("/login", userLogin)
	AuthApi.Get("/me", me)
	AuthApi.Get("/getAccessToken", getAccessToken)

}
