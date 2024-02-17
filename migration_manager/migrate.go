package migration_manager

import (
	"emreddit/db"
	"emreddit/logger"
	"emreddit/migration"
	"errors"
)

type CommittedMigration struct {
	Name string `gorm:"unique;column:name;not null;"`
}

var CommittedMigs []CommittedMigration

func (table *CommittedMigration) TableName() string {

	return "migrations"
}

func InsertMigration(name string) error {

	// inserting migration
	if err := db.Db.Create(&CommittedMigration{Name: name}).Error; err != nil {
		return err
	}

	logger.Info(name, " has been saved")
	return nil
}

func DeleteMigration(name string) error {
	//deleting migration
	err := db.Db.Where("Name=?", name).Delete(&CommittedMigration{}).Error

	if err != nil {
		logger.Info("delete err", name)
		return err
	}

	logger.Info("deleted", name)
	return nil
}

// getting migrations from db
func getMigsFromDB() error {
	if err := db.Db.Find(&CommittedMigs).Error; err != nil {
		//check if err
		return err
	}

	return nil
}

// searching for specific migration with Name
func SearchMigration(name string) error {
	for _, CommittedMig := range CommittedMigs {
		if CommittedMig.Name == name {
			return nil
		}
	}
	return errors.New("mig doesnt exist " + name)
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

// running up for specific migration
func RunUpMigration(name string) error {

	// if it exists doesnt run it again
	if err := SearchMigration(name); err == nil {
		logger.Error(" run up mig has found <?> run before", name)

		return err

	}
	for _, migElement := range migration.Migrations_Arr {
		if migElement.Name == name {
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
	logger.Error("Error:Mig not found in Run Time Array")
	return nil

}

func RunDown() error {

	for _, migElement := range migration.Migrations_Arr {

		var err error = nil
		if err = SearchMigration(migElement.Name); err != nil {
			logger.Info("Migration has not found, Never Inserted before")

			return err

		}

		if err := migElement.DownFn(); err != nil {
			logger.Error(migElement.Name, " Migration Down Function Error: <?>", err)
			return err
		}

		if err := DeleteMigration(migElement.Name); err != nil {
			logger.Error("Delete err <?> ", migElement.Name)
			return err
		}

		logger.Info("Deleted Migration", migElement.Name) // insert

	}

	return nil

}

// Run Down Migration With Given Name
func RunDownMigration(name string) error {

	if err := SearchMigration(name); err != nil {
		logger.Error(" Migration has not found <?>, Run Up Function before", name)

		return err

	}
	for _, migElement := range migration.Migrations_Arr {
		if migElement.Name == name {
			if err := migElement.DownFn(); err != nil {
				logger.Error(migElement.Name, "Migration Down Function Error due to <?>", err)
				return err
			}

			if err := DeleteMigration(migElement.Name); err != nil {
				logger.Error("Migration Cannot Be Deleted Due To <?> ", err, migElement.Name)
				return err
			}

			logger.Info("Migration deleted  <?>", migElement.Name)
			return nil //

		}

	}
	logger.Error("Migration Not Found In Run Time Arr")
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
