package migrationutils

import (
	"emreddit/logger"
	"emreddit/migration"
)

func RunDown() error {

	if err := GetMigsFromDB(); err != nil { //getting MigsFrom DB
		logger.Error("getCommittedMigErr")
		return err
	}
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
	var remove_index int = 0
	if err := SearchMigration(Name); err != nil { //searching for migration if it doesnt exist return err
		logger.Error("mig has not found")

		return err

	}
	for index, migElement := range migration.Migrations_Arr {
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
			remove_index = index
			break //
			// String found in the struct array
		}

	}
	migration.Migrations_Arr = append(migration.Migrations_Arr[:remove_index], migration.Migrations_Arr[remove_index+1:]...) // reordering migrations_Arr after deleting
	return nil

}
