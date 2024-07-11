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
	TableName string
}

func Init(FileName string) {

	timestamp := time.Now().Format("20060102150405")
	templateFile := "migration_manager/migration.tmpl"
	fileName := fmt.Sprintf("migration/%s-%s.go", timestamp, FileName)

	migrationFile := []Migration{{Name: FileName, Timestamp: timestamp, TableName: FileName + "s"}}

	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		//panic(err)
		logger.Fatal(err.Error())
	}

	File, err := os.Create(fileName)
	if err != nil {
		logger.Fatal(err)
	}

	err = tmpl.Execute(File, migrationFile)

	if err != nil {

		File.Close()
		logger.Error()

	}

	File.Close()

}
