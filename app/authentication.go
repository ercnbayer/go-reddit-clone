package app

import (
	"emreddit/db"
	"emreddit/logger"
)

func CreateUser(user *db.UserEntity) error {

	err := db.CreateUser(user)

	if err != nil {
		logger.Error(" User cant be created :<?>", err)
		return err
	}
	return nil
}

func UserLogin(user *db.UserEntity) error {

	if err := db.GetUserByEmailAndPassword(user); err != nil { // sending it to db

		logger.Error("Login Error: <?>", err)
		return err
	}

	return nil
}
