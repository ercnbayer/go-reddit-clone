package app

import "emreddit/db"

func ReadComment(id string) (*db.Comment, error) {

	return db.ReadComment(id)
}

func CreateComment(comment *db.Comment) error {

	return db.CreateComment(comment)
}

func DeleteComment(id string, reqID string) (string, error) {

	// TO DO ADD IF reqId permitted to do the operation

	return db.DeleteComment(id)
}

func UpdateComment(comment *db.Comment, id string) error {

	// TO DO ADD IF id permitedd to do the operation

	return db.UpdateComment(comment)
}
