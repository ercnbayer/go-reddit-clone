package db

import (
	"emreddit/logger"
	"errors"
)

type User struct {
	ID       string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name     string `gorm:"column:name;not null;default:null"`
	Password string `gorm:"column:password;not null;default:null"`
	Email    string `gorm:"unique;not null;type:varchar(100);default:null"`
}

// TableName overrides the table name used by User to `profiles`
func (table *User) TableName() string {
	return "users"
}

func InsertUser(user *User) error {

	if err := Db.Save(user).Error; err != nil {
		return err
	}

	logger.Info(user)

	return nil

}

func LoginUser(user *User) error {

	if err := Db.Where(" email=? AND password=? ", user.Email, user.Password).First(&user).Error; err != nil {
		logger.Info("Wrong Creds entry", err)
		return err
	}

	return nil

}

func DeleteUser(id string) (string, error) {

	var QueryResult = Db.Delete(&User{ID: id})

	if err := QueryResult.Error; err != nil {

		logger.Info("delete ", err)
		return id, err

	}
	if QueryResult.RowsAffected == 0 {

		logger.Info("USER IS NOT FOUND")

		return id, errors.New("user NOT FOUND")
	}

	return id, nil
}

func ReadUser(id string, user *User) error {

	/*if err := Db.First(&Person{ID: id}).Error; err != nil {
		return err
	}*/

	user.ID = id

	if err := Db.First(user).Error; err != nil {
		return err
	}

	return nil
}

func PatchUpdateUser(user *User) error {

	var result = Db.Updates(user)
	if err := result.Error; err != nil {

		logger.Error("err update", err)
		return err

	}

	if result.RowsAffected == 0 {
		logger.Error("err user not found")

		return errors.New("user NOT FOUND")
	}

	// Return the updated person
	return nil

}

func GetUsers() ([]User, error) {
	var users []User // creating person arr
	if err := Db.Find(&users).Error; err != nil {
		//check if err
		return nil, err
	}

	return users, nil
}
