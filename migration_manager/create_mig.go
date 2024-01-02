package migration_manager

import (
	"emreddit/logger"
	"fmt"
	"os"
	"text/template"
	"time"
)

type Migration struct {
	Name      string `gorm:"column:name;not null;"`
	Timestamp string
	TableName string `gorm:"column:name;not null;"`
}

func Init(FileName string) {

	timestamp := time.Now().Format("20060102150405") //setting time format
	templateFile := "migration-utils/migration.tmpl" // migration.tmpl's  relative file path
	fileName := fmt.Sprintf("migration/%s-%s.go", timestamp, FileName)

	migrationFile := []Migration{{Name: "User", Timestamp: timestamp, TableName: "users"}} //creating first migrations

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
