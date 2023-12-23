package db

import (
	"fmt"

	"emreddit/config"
	_ "emreddit/config"
	"emreddit/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {

	var err error
	PostgresConfigString := fmt.Sprintf("host=localhost user=%s password=%s database=%s port=%s sslmode=disable TimeZone=Etc/UTC",
		config.PsqlDbConfig.UserName, config.PsqlDbConfig.DbPassword, config.PsqlDbConfig.DbName, config.PsqlDbConfig.DbPort)

	Db, err = gorm.Open(postgres.Open(PostgresConfigString), &gorm.Config{}) //connecting gorm

	if err != nil { //check if err
		logger.Fatal(err)
	}

}
