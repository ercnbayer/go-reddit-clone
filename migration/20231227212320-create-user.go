package migration

import (
	"emreddit/db"
)

type User20231227212320 struct {
	ID       string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name     string `gorm:"column:name;not null;"`
	Password string `gorm:"column:password;not null;"`
	Email    string `gorm:"unique;not null;type:varchar(100);"`
}

func (table User20231227212320) TableName() string {
	return "users"
}
func UserUp20231227212320() error {

	return db.Db.Migrator().CreateTable(&User20231227212320{})
}
func UserDown20231227212320() error {

	return db.Db.Migrator().DropTable(&User20231227212320{})
}

func init() {

	Migrations_Arr = append(Migrations_Arr, Migration{
		Name:   "User20231227212320",
		UpFn:   UserUp20231227212320,
		DownFn: UserDown20231227212320,
	})

}
