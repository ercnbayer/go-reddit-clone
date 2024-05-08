package db

import (
	"emreddit/logger"
	"errors"
	"time"
)

type RefreshToken struct {
	ID         string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	IsUsed     bool      `gorm:"column:is_used;not null;default:false"`
	ExpireDate time.Time `gorm:"column:expire_date;"`
	UserID     string    `gorm:"type:uuid;not null;"`
}

func (table *RefreshToken) TableName() string {
	return "refresh_tokens" //set TableName to operate
}

func CreateToken(token *RefreshToken) error { //inserting token

	var expire_date = time.Hour * 24 * 30
	token.ExpireDate = time.Now().Add(expire_date)
	if err := Db.Save(token).Error; err != nil { //checking for errors
		return err
	}

	logger.Info(token)

	return nil

}

func ReadToken(id string) (RefreshToken, error) {

	token := RefreshToken{}
	token.ID = id
	if err := Db.Where("Expire_Date > ?", time.Now()).Where(&RefreshToken{ID: id}).First(&token).Error; err != nil {

		return RefreshToken{}, err

	}

	return token, nil

}

func UpdateToken(token *RefreshToken) error {
	var result = Db.Updates(token)
	if err := result.Error; err != nil { //checking for errors

		logger.Error("err update", err)
		return err

	}

	if result.RowsAffected == 0 { //check if any operation affects table
		logger.Error("err user not found")

		return errors.New("user NOT FOUND")
	}
	return nil
}
