package migration

import (
	"emreddit/db"
	"time"
)

type RefreshToken20240426195151 struct {
	ID         string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	IsUsed     bool      `gorm:"column:is_used;not null;default:false"`
	ExpireDate time.Time `gorm:"column:expire_date;"`
	UserID     string    `gorm:"type:uuid;not null;"`
}

func (table RefreshToken20240426195151) TableName() string {
	return "refresh_tokens"
}
func RefreshTokenUp20240426195151() error {

	return db.Db.Migrator().CreateTable(&RefreshToken20240426195151{})
}
func RefreshTokenDown20240426195151() error {

	return db.Db.Migrator().DropTable(&RefreshToken20240426195151{})
}

func init() {

	Migrations_Arr = append(Migrations_Arr, Migration{
		Name:   "20240426195151RefreshToken",
		UpFn:   RefreshTokenUp20240426195151,
		DownFn: RefreshTokenDown20240426195151,
	})

}
