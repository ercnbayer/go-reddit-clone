package migration

import (
	"emreddit/db"
	"time"
)

type User20240409155741 struct {
	RefreshToken string    `gorm:"column:Refresh_Token; type:varchar(20);"`
	IsUsed       bool      `gorm:"column:Is_Used;"`
	ExpireDate   time.Time `gorm:"column:Expire_Date"`
}

func (table User20240409155741) TableName() string {
	return "users"
}
func UserUp20240409155741() error {

	err := db.Db.Migrator().AddColumn(&User20240409155741{}, "Refresh_Token")

	if err != nil {
		return err
	}

	err = db.Db.Migrator().AddColumn(&User20240409155741{}, "Is_Used")

	if err != nil {
		return err
	}
	db.Db.Migrator().AddColumn(&User20240409155741{}, "Expire_Date")
	if err != nil {
		return err
	}
	return nil
}
func UserDown20240409155741() error {

	err := db.Db.Migrator().DropColumn(&User20240409155741{}, "Refresh_Token")

	if err != nil {
		return err
	}

	err = db.Db.Migrator().DropColumn(&User20240409155741{}, "Is_Used")

	if err != nil {
		return err
	}
	db.Db.Migrator().DropColumn(&User20240409155741{}, "Expire_Date")
	if err != nil {
		return err
	}

	return nil
}

func init() {

	Migrations_Arr = append(Migrations_Arr, Migration{
		Name:   "20240409155741User",
		UpFn:   UserUp20240409155741,
		DownFn: UserDown20240409155741,
	})

}
