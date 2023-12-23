package migration

import (
	"emreddit/db"
	"emreddit/logger"
)

type User20231222185637 struct {
	ID       string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name     string `gorm:"column:name;not null;default:null"`
	Password string `gorm:"column:password;not null;default:null"`
	Email    string `gorm:"unique;not null;type:varchar(100);default:null"`
}

func (table User20231222185637) TableName() string {
	return "users"
}
func UserUp20231222185637() error {

	if err := db.Db.Migrator().CreateTable(&User20231222185637{}); err != nil {

		logger.Info("Table Init:", err)

		return err
	}

	return nil
}
func UserDown20231222185637() error {

	if err := db.Db.Migrator().DropTable(&User20231222185637{}); err != nil {

		logger.Info("Table Drop:", err)

		return err
	}

	return nil
}

func init() {

	Migrations_Arr = append(Migrations_Arr, Migration{
		Name:   "User20231222185637",
		UpFn:   UserUp20231222185637,
		DownFn: UserDown20231222185637,
	})
	logger.Info("Table Init")

}
