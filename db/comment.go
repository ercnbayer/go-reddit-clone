package db

import (
	"emreddit/logger"
	"errors"
)

type Comment struct {
	ID          string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	OwnerID     string `gorm:"type:uuid; not null;"`
	PostID      string `gorm:"type:uuid;not null;"`
	Description string `gorm:"type:varchar(100);"`
}

func CreateComment(comment *Comment) error {

	if err := Db.Save(comment).Error; err != nil { //checking for errors
		return err
	}

	logger.Info(comment)

	return nil
}

func DeleteComment(id string) (string, error) {

	var QueryResult = Db.Delete(&Comment{ID: id})

	if err := QueryResult.Error; err != nil {

		logger.Info("delete ", err)
		return id, err

	}
	if QueryResult.RowsAffected == 0 { //checking affecting row to know if any operation took effect

		logger.Info("COMMENT IS NOT FOUND")

		return id, errors.New("COMMENT IS NOT FOUND")
	}

	return id, nil
}

func ReadComment(id string) (*Comment, error) {

	comment := new(Comment)
	comment.ID = id //setting id

	if err := Db.First(comment).Error; err != nil { //checking for errors.
		return nil, err
	}

	return comment, nil
}
func UpdateComment(comment *Comment) error {
	var result = Db.Updates(comment)
	if err := result.Error; err != nil { //checking for errors

		logger.Error("err  post update", err)
		return err

	}

	if result.RowsAffected == 0 { //check if any operation affects table
		logger.Error(" err : post not found")

		return errors.New("post NOT FOUND")
	}

	// Return the updated person
	return nil
}
