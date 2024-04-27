package db

import (
	"emreddit/logger"
	"time"
)

type RefreshToken struct {
	ID         string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	IsUsed     bool      `gorm:"column:Is_Used;not null;default:false"`
	ExpireDate time.Time `gorm:"column:Expire_Date;"`
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
	if err := Db.First(&token).Error; err != nil {

		return RefreshToken{}, err

	}

	return token, nil

}
