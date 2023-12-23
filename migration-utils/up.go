package migrationutils

import (
	"emreddit/logger"
	"emreddit/migration"
)

func RunUp() error {

	migration.SortMigArr()

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

func RunUpMigration(Name string) error {

	if err := SearchMigration(Name); err == nil {
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
			// String found in the struct array
		}

	}
	logger.Error("Mig not found")
	return nil

}
