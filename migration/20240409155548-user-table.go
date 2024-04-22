package migration

import "emreddit/db"

type User20240409155548 struct {
	ID       string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name     string `gorm:"column:name;not null;"`
	Password string `gorm:"column:password;not null;"`
	Email    string `gorm:"unique;not null;type:varchar(100);"`
}

func (table User20240409155548) TableName() string {
	return "users"
}
func UserUp20240409155548() error {

	return db.Db.Migrator().CreateTable(&User20240409155548{})
}
func UserDown20240409155548() error {

	return db.Db.Migrator().DropTable(&User20240409155548{})
}

func init() {

	Migrations_Arr = append(Migrations_Arr, Migration{
		Name:   "20240409155548User",
		UpFn:   UserUp20240409155548,
		DownFn: UserDown20240409155548,
	})

}
