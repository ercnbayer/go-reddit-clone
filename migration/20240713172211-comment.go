package migration

import "emreddit/db"

type Comment20240713172211 struct {
	ID          string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	OwnerID     string `gorm:"type:uuid; not null;"`
	PostID      string `gorm:"type:uuid;not null;"`
	Description string `gorm:"type:varchar(100);"`
}

func (table Comment20240713172211) TableName() string {
	return "comments"
}
func CommentUp20240713172211() error {

	return db.Db.Migrator().CreateTable(&Comment20240713172211{})
}
func CommentDown20240713172211() error {

	return db.Db.Migrator().DropTable(&Comment20240713172211{})
}

func init() {

	Migrations_Arr = append(Migrations_Arr, Migration{
		Name:   "20240713172211Comment",
		UpFn:   CommentUp20240713172211,
		DownFn: CommentDown20240713172211,
	})

}
