package app

import (
	"emreddit/db"
)

func GetUser(id string) (*db.UserEntity, error) {
	var user *db.UserEntity = new(db.UserEntity)
	err := db.ReadUser(id, user)
	if err != nil {
		return nil, err
	}
	return user, err
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
