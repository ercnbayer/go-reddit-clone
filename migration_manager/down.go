package migration_manager

import (
	"emreddit/logger"
	"emreddit/migration"
)

func RunDown() error {

	if err := migration.Migrations_Arr[len(migration.Migrations_Arr)-1].DownFn(); err != nil { //executing down func and  checking for err
		logger.Error(" DOWN FUNC ERR ", migration.Migrations_Arr[len(migration.Migrations_Arr)-1].Name)
		return err
	}
	if err := DeleteMigration(CommittedMigs[len(migration.Migrations_Arr)-1].Name); err != nil { //deleting migration from db

		logger.Error("DELETE ERR ", CommittedMigs[len(migration.Migrations_Arr)-1].Name)
		return err
	}

	return nil

}

func RunDownMigration(Name string) error { //rundown migration with given name

	for _, migElement := range migration.Migrations_Arr {
		if migElement.Name == Name {
			if err := migElement.DownFn(); err != nil { // if err happens when executing return err
				logger.Error(migElement.Name, " up func err")
				return err
			}

			if err := DeleteMigration(migElement.Name); err != nil { //deleting migration from db
				logger.Error("delete err ", migElement.Name)
				return err
			}

			logger.Info("DELETED MIGRATION ", migElement.Name)

			return nil
			// String found in the struct array
		}

	}

	return nil

}
