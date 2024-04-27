package app

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"emreddit/config"
	"emreddit/db"
	"emreddit/logger"
	"encoding/base64"
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type SessionToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

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
func tryCreate(userID string) (db.RefreshToken, error) {

	token := db.RefreshToken{UserID: userID}
	if err := db.CreateToken(&token); err != nil {
		return db.RefreshToken{}, err
	}
	return token, nil
}

func CreateRefreshToken(userID string) (string, error) {

	count := 10
	var err error
	for count > 1 {

		token, err := tryCreate(userID)
		if err == nil {
			return token.ID, nil
		}
		count--
	}
	return "", err
}

func JSONToBytes(userTokens *SessionToken) ([]byte, error) {
	return json.Marshal(userTokens)
}

func CheckIfTokenValid(refresh_token string) (string, error) {

	token, err := db.ReadToken(refresh_token)

	if err != nil {
		return "", err
	}

	if token.ExpireDate.Compare(time.Now()) < 0 || !token.IsUsed {
		return "", errors.New("invalid token")
	}
	token.IsUsed = true

	return token.UserID, nil

}

func pkcs5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])

	return src[:(length - unpadding)]
}

func EncryptToken(tokens []byte) (string, error) {

	var plainTextBlock []byte
	length := len(tokens)

	if length%16 != 0 {
		extendBlock := 16 - (length % 16)
		plainTextBlock = make([]byte, length+extendBlock)
		copy(plainTextBlock[length:], bytes.Repeat([]byte{uint8(extendBlock)}, extendBlock))
	} else {
		plainTextBlock = make([]byte, length)
	}

	copy(plainTextBlock, tokens)
	block, err := aes.NewCipher(config.AES_KEY)

	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, len(plainTextBlock))
	mode := cipher.NewCBCEncrypter(block, config.IV)
	mode.CryptBlocks(ciphertext, plainTextBlock)

	str := base64.URLEncoding.EncodeToString(ciphertext)

	return str, nil
}

func bytesToSessionToken(pb []byte) (*SessionToken, error) {

	SessionTokens := new(SessionToken)
	err := json.Unmarshal(pb, SessionTokens)
	if err != nil {
		logger.Info("JSON Unmarshal Error:<?>", err)
		return nil, err
	}
	return SessionTokens, nil
}

func DecryptToken(encoded_token string) (*SessionToken, error) {

	decoded_str, err := decodeFromb64(encoded_token)

	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(config.AES_KEY)

	if err != nil {
		return nil, err
	}

	if len(decoded_str)%aes.BlockSize != 0 {
		return nil, errors.New("BLOCK SIZE CANNOT BE ZERO")
	}

	mode := cipher.NewCBCDecrypter(block, config.IV)
	mode.CryptBlocks(decoded_str, decoded_str)
	decoded_str = pkcs5UnPadding(decoded_str)

	return bytesToSessionToken(decoded_str)
}

func decodeFromb64(str string) ([]byte, error) {

	byte_arr, err := b64.URLEncoding.DecodeString(str)
	if err != nil {
		return nil, errors.New("INVALID TOKEN")

	}

	return byte_arr, nil
}

func CreateJWT(id string) (string, error) {

	var expire_date = time.Minute * 15
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(expire_date).Unix(),
		"iat":     time.Now().Unix(),
		"subject": id,
	})

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
		logger.Error("Error <?>", err)
	}

	return id, nil
}
