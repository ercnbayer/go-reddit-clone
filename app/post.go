package app

import "emreddit/db"

func ReadPost(id string) (*db.Post, error) {

	return db.ReadPost(id)
}

func CreatePost(post *db.Post) error {
	//
	return db.CreatePost(post)
}

func DeletePost(id string) (string, error) {
	return db.DeletePost(id)
}

func UpdatePost(post *db.Post, reqID string) error {

	// to do: check if reqId is permitted
	return db.UpdatePost(post)
}
