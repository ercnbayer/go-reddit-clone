package app

import "emreddit/db"

func ReadComment(id string) (*db.Comment, error) {

	return db.ReadComment(id)
}

func CreateComment(comment *db.Comment) error {
	return db.CreateComment(comment)
}

func DeleteComment(id string) (string, error) {

	return db.DeleteComment(id)
}

func UpdateComment(comment *db.Comment) error {

	return db.UpdateComment(comment)
}
