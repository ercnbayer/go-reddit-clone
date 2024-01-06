package migration_manager

import (
	"emreddit/db"
	"emreddit/logger"
	"emreddit/migration"
	"errors"
)

type CommittedMigration struct {
	ID   string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name string `gorm:"unique;column:name;not null;"`
}

var CommittedMigs []CommittedMigration

func (table *CommittedMigration) TableName() string {

	return "migrations"
}

func InsertMigration(Name string) error {

	if err := db.Db.Save(&CommittedMigration{Name: Name}).Error; err != nil { // inserting migration
		return err
	}

	logger.Info(Name, "has been saved")
	return nil
}

func DeleteMigration(Name string) error {

	err := db.Db.Where("Name=?", Name).Delete(&CommittedMigration{}).Error //deleting migration

	if err != nil {
		logger.Info("delete err", Name)
		return err
	}

	logger.Info("deleted", Name)
	return nil
}
func getMigsFromDB() error { //getting migrations from db
	if err := db.Db.Find(&CommittedMigs).Error; err != nil {
		//check if err
		return err
	}

	return nil
}

func SearchMigration(Name string) error { //searching for specific migration with Name
	for _, CommittedMig := range CommittedMigs {
		if CommittedMig.Name == Name {
			return nil
		}
	}
	return errors.New("mig doesnt exist " + Name)
}

func RunUp() error {

	for _, migElement := range migration.Migrations_Arr {

		var err error = nil
		if err = SearchMigration(migElement.Name); err == nil {
			logger.Info("mig has found, inserted before")

			return err

		}

		if err := migElement.UpFn(); err != nil {
			logger.Error(migElement.Name, " up func err")
			return err
		}

		if err := InsertMigration(migElement.Name); err != nil {
			logger.Error("insert err ", migElement.Name)
			return err
		}

		logger.Info("INSERTED NEW MIGRATION", migElement.Name) // insert

	}

	return nil

}

func RunUpMigration(Name string) error { //running up for specific migration

	if err := SearchMigration(Name); err == nil { // if it exists doesnt run it again
		logger.Error(" run up mig has found, run before")

		return err

	}
	for _, migElement := range migration.Migrations_Arr {
		if migElement.Name == Name {
			if err := migElement.UpFn(); err != nil {
				logger.Error(migElement.Name, " up func err")
				return err
			}

			if err := InsertMigration(migElement.Name); err != nil {
				logger.Error("insert err ", migElement.Name)
				return err
			}

			logger.Info("INSERTED NEW MIGRATION", migElement.Name)
			return nil //

		}

	}
	logger.Error("Mig not found") // mig Not Found
	return nil

}

func init() {

	if !db.Db.Migrator().HasTable(&CommittedMigration{}) { //check if migration table exists
		if err := db.Db.Migrator().CreateTable(&CommittedMigration{}); err != nil {
			panic("failed to create table")
		}
	}
	if err := getMigsFromDB(); err != nil {
		panic("no migration exists")
	}
}
