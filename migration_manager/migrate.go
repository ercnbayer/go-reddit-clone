package migration_manager

import (
	"emreddit/db"
	"emreddit/logger"
	"emreddit/migration"
)

type CommittedMigration struct {
	ID   string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name string `gorm:"unique;column:name;not null;"`
}

//var CommittedMigs []CommittedMigration no longer neccessary i hope

func (table *CommittedMigration) TableName() string {

	return "migrations"
}

func InsertMigration(Name string) error {

	// inserting migration
	if err := db.Db.Save(&CommittedMigration{Name: Name}).Error; err != nil {
		return err
	}

	logger.Info(Name, "<?> has been saved")
	return nil
}

func DeleteMigration(Name string) error {
	//deleting migration
	err := db.Db.Where("Name=?", Name).Delete(&CommittedMigration{}).Error

	if err != nil {
		logger.Info("delete err", Name)
		return err
	}

	logger.Info("deleted", Name)
	return nil
}

// getting migrations from db
/*func getMigsFromDB() error {
	if err := db.Db.Find(&CommittedMigs).Error; err != nil {
		//check if err
		return err
	}

	return nil
}*/

func SearchMigration(Name string) error {
	commitedMig := CommittedMigration{Name: Name}
	if err := db.Db.Where("Name =?", Name).First(&commitedMig).Error; err != nil {
		logger.Info(err)
		return err
	}

	return nil
}

// searching for specific migration with Name
/*func SearchMigration(Name string) error {
	for _, CommittedMig := range CommittedMigs {
		if CommittedMig.Name == Name {
			return nil
		}
	}
	return errors.New("mig doesnt exist " + Name)
}*/

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
func RunUpMigration(Name string) error {

	// if it exists doesnt run it again
	if err := SearchMigration(Name); err == nil {
		logger.Error(" run up mig has found <?>, run before", Name)

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
func RunDownMigration(Name string) error {

	if err := SearchMigration(Name); err != nil {
		logger.Error(" Migration has not found <?>, Run Up Function before", Name)

		return err

	}
	for _, migElement := range migration.Migrations_Arr {
		if migElement.Name == Name {
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
	/*if err := getMigsFromDB(); err != nil {
		panic("no migration exists")
	}*/
}
