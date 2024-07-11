package migration

import "emreddit/db"

type Post20240629154420 struct {
	ID          string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	OwnerID     string `gorm:"type:uuid;not null;"`
	Upvote      int    `gorm:"type:int;not null"`
	Downvote    int    `gorm:"type:int;not null"`
	Description string `gorm:"type:varchar(100);"`
}

func (table Post20240629154420) TableName() string {
	return "posts"
}
func PostUp20240629154420() error {

	return db.Db.Migrator().CreateTable(&Post20240629154420{})
}
func PostDown20240629154420() error {

	return db.Db.Migrator().DropTable(&Post20240629154420{})
}

func init() {

	Migrations_Arr = append(Migrations_Arr, Migration{
		Name:   "20240629154420Post",
		UpFn:   PostUp20240629154420,
		DownFn: PostDown20240629154420,
	})

}
