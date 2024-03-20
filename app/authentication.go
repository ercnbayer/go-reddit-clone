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
	//logine de ek olarak tek kullanılımlık refresh token uret 1-2 saat veya 1 gün
	//bunlar bir keyle şifreleniyor +jwt tokenle olacak bicimde // tek bir string dönüyorsun bunu da base64lemen dönüyor
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour).Unix(), //15 dk üret
		"iat":     time.Now().Unix(),
		"subject": id})

	ss, err := token.SignedString([]byte(config.JWTKey))

	if err != nil {
		logger.Error("JWT Generate Error:<?>", err)
		return "", err
	}
	logger.Info(token.Header)
	//base64le encode
	return ss, nil
} //jwt token ise 401 // 401
func EncryptToken(id string) {
	//kullanıcı token // jwt token
}

// bunlara ek olarak refresh token olustur bunu db de yaz
func ParseJWT(tokenString string) (string, error) {
	//base64 decode ekle
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
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

//yeni bir endpoint apiye eklenecek
