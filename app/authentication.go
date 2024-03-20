package app

import (
	"emreddit/config"
	"emreddit/db"
	"emreddit/logger"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func RegisterUser(user *db.UserEntity) error {

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

func CreateJWT(id string) (string, error) {

	var expire_date = time.Minute * 15
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(expire_date).Unix(), //15 dk üret
		"iat":     time.Now().Unix(),
		"subject": id})

	ss, err := token.SignedString([]byte(config.JWTKey))

	if err != nil {
		logger.Error("JWT Generate Error:<?>", err)
		return "", err
	}
	logger.Info(token.Header)

	return ss, nil
}

func ParseJWT(tokenString string) (string, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(config.JWTKey), nil
	})
	if err != nil {
		logger.Error(err)
		return "", err
	}
	var id string
	if _, ok := token.Claims.(jwt.MapClaims); ok {
		id, err = token.Claims.GetSubject()
		if err != nil {
			logger.Error("Claim Error <?>", err)
			return "", err
		}

	} else {
		logger.Error("Error:?", err)
	}

	return id, nil
}
