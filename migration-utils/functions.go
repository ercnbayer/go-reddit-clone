package migrationutils

import (
	"emreddit/db"
	"emreddit/logger"
)

type CommittedMigration struct {
	ID   string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name string `gorm:"name"`
}

var CommittedMigs []CommittedMigration

func (table *CommittedMigration) TableName() string {

	return "migrations"
}

func InsertMigration(Name string) error {

	if err := db.Db.Save(&CommittedMigration{Name: Name}).Error; err != nil {
		return err
	}

	logger.Info(Name, "has been saved")
	return nil
}

func DeleteMigration(Name string) error {

	err := db.Db.Where("Name=?", Name).Delete(&CommittedMigration{}).Error

	if err != nil {
		logger.Info("delete err", Name)
		return err
	}

	logger.Info("deleted", Name)
	return nil
}
func GetMigsFromDB() error {
	if err := db.Db.Find(&CommittedMigs).Error; err != nil {
		//check if err
		return err
	}

	logger.Info("commited migs get success")

	return nil
}

func SearchMigration(Name string) error {
	err := db.Db.Where("Name=?", Name).First(&CommittedMigration{}).Error

	if err != nil {
		logger.Info("cant found entry")
		return err
	}
	logger.Info("FOUND ENTRY")
	return nil
}

func init() {

	if !db.Db.Migrator().HasTable(&CommittedMigration{}) {
		if err := db.Db.Migrator().CreateTable(&CommittedMigration{}); err != nil {
			panic("failed to create table")
		}
	}
}
