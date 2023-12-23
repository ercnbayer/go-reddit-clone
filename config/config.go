package config

import (
	"emreddit/logger"
	"os"

	"github.com/joho/godotenv"
)

type PgDbStruct struct {
	DbName     string
	UserName   string
	DbPassword string
	DbPort     string
}

var PsqlDbConfig PgDbStruct

func CheckEnvArgs() bool {

	return PsqlDbConfig.DbName == "" || PsqlDbConfig.DbPassword == "" || PsqlDbConfig.UserName == "" || PsqlDbConfig.DbPort == ""
}

func init() {

	logger.LogLevel = 3

	err := godotenv.Load() //load env
	if err != nil {
		logger.Error("Error loading .env file")
		return
	}

	PsqlDbConfig.DbName = os.Getenv("POSTGRES_DB") // getting env vars
	PsqlDbConfig.UserName = os.Getenv("POSTGRES_USER")
	PsqlDbConfig.DbPassword = os.Getenv("POSTGRES_PASSWORD")
	PsqlDbConfig.DbPort = os.Getenv("PORT")

	if CheckEnvArgs() {

		logger.Fatal("ENV VALUE EMPTY ")
	}

}
