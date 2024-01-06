package app

import "emreddit/db"

type UserUpdatePayload struct { //payload for updating user
	ID       string
	Name     string `validate:"omitempty,required"` /* if struct field not empty validate */
	Password string `validate:"omitempty,required"`
	Email    string `validate:"omitempty,required,email"`
}

func GetUser(id string) error {
	var user db.UserEntity
	return db.ReadUser(id, &user)
}
