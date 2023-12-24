package migrationutils

import (
	"emreddit/logger"
	"fmt"
	"os"
	"text/template"
	"time"
)

type Migration struct {
	Name           string
	Timestamp      string
	TableName      string
	SuccessInitLog string
}

func Init() {

	timestamp := time.Now().Format("20060102150405") //setting time format
	templateFile := "migration-utils/migration.tmpl" // migration.tmpl's  relative file path
	fileName := fmt.Sprintf("migration/%s.go", timestamp)

	migrationFile := []Migration{{Name: "User", Timestamp: timestamp, TableName: "users", SuccessInitLog: "Table Init"}} //creating first migrations

	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		//panic(err)
		logger.Error(err.Error())
	}

	File, err := os.Create(fileName) //creating file

	if err != nil {
		logger.Error(err)
	}

	err = tmpl.Execute(File, migrationFile) //executing tmpl

	if err != nil {

		File.Close()
		logger.Error()

	}

	File.Close()
	// end main

}
