package app

import (
	"emreddit/db"
)

func GetUser(id string) (error, *db.UserEntity) {
	var user *db.UserEntity
	err := db.ReadUser(id, user)
	if err != nil {
		return err, nil
	}
	return nil, user
}
func UpdateUser(dbUser *db.UserEntity) error {

	return db.UpdateUser(dbUser)
}

func GetUsers() ([]db.UserEntity, error) {
	return db.GetUsers()

}

func DeleteUser(id string) (string, error) {

	return db.DeleteUser(id)
}
