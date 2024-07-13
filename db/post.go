package db

import (
	"emreddit/logger"
	"errors"
)

type Post struct {
	ID          string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	OwnerID     string
	Owner       UserEntity `gorm:"foreignKey:OwnerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Upvote      int        `gorm:"type:int;not null"`
	Downvote    int        `gorm:"type:int;not null"`
	Description string     `gorm:"type:varchar(100);"`
}

func CreatePost(post *Post) error { //inserting comment

	logger.Info(post)
	if err := Db.Save(post).Error; err != nil { //checking for errors
		return err
	}

	logger.Info(post)

	return nil

}

func DeletePost(id string) (string, error) { //delete User

	var QueryResult = Db.Delete(&Post{ID: id})

	if err := QueryResult.Error; err != nil {

		logger.Info("delete ", err)
		return id, err

	}
	if QueryResult.RowsAffected == 0 { //checking affecting row to know if any operation took effect

		logger.Info("POST IS NOT FOUND")

		return id, errors.New("POST NOT FOUND")
	}

	return id, nil
}

func ReadPost(id string) (*Post, error) {
	post := new(Post)
	post.ID = id //setting id

	if err := Db.Find(post).Error; err != nil { //checking for errors.
		return &Post{}, err
	}

	return post, nil
}

func UpdatePost(post *Post) error {
	var result = Db.Updates(post)
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

func GetAllPosts(OwnerID string) ([]Post, error) {

	var UserPosts []Post
	if err := Db.Where("owner_id=?", OwnerID).Find(&UserPosts).Error; err != nil {

		return nil, err
	}
	return UserPosts, nil
}
