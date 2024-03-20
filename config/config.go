package config

import (
	"emreddit/logger"
	"os"

	"github.com/joho/godotenv"
)

var DbName string
var UserName string
var DbPassword string
var DbPort string
var ListenPort string
var JWTKey string

func getDbName() bool {
	DbName = os.Getenv("POSTGRES_DB")
	if DbName == "" {
		logger.Fatal("Null Value")
		return false

	}
	return true

}
func getUserName() bool {
	UserName = os.Getenv("POSTGRES_USER")
	if UserName == "" {
		logger.Fatal("Null Value db UserName")
		return false
	}
	return true

}
func getDbPort() bool {
	DbPort = os.Getenv("PORT")
	if DbPort == "" {
		logger.Fatal("Null Value")
		return false

	}
	return true

}
func getDbPassword() bool {
	DbPassword = os.Getenv("POSTGRES_PASSWORD")
	if DbPassword == "" {
		logger.Fatal("Null Value")
		return false

	}
	return true

}
func getListenPort() bool {
	ListenPort = os.Getenv("LISTEN_PORT")
	if ListenPort == "" {
		logger.Fatal("Null Value")
		return false

	}
	return true
}
func getJWTKey() bool {
	JWTKey = os.Getenv("JWTKEY")
	if ListenPort == "" {
		logger.Fatal("Null Value")
		return false

	}
	return true
}

func init() {

	logger.LogLevel = logger.AllLogs

	err := godotenv.Load() //load env
	if err != nil {
		logger.Error("Error loading .env file")
		return
	}

	DbName = "emreddit"   //getDbName()// getting env vars
	UserName = "postgres" //getUserName()
	DbPassword = "root"   //getDbPassword()
	DbPort = "6000"       //getDbPort()
	getJWTKey()

}
