package api

import db "emreddit/db"

type UserPayload struct { //payload for registerUser
	ID       string
	Name     string `validate:"required"`
	Password string `validate:"required"`
	Email    string `validate:"required,email"`
}

type UserUpdatePayload struct { //payload for updating user
	ID       string
	Name     string `validate:"omitempty,required"`
	Password string `validate:"omitempty,required"`
	Email    string `validate:"omitempty,required,email"`
}

type UserLoginPayload struct { // payload for Login User
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

func mapUserUpdatePayloadToDbUser(user *UserUpdatePayload, dbUser *db.User) {

	dbUser.ID = user.ID
	dbUser.Name = user.Name
	dbUser.Email = user.Email
	dbUser.Password = user.Password
}

func mapUserLoginPayloadToDbUser(user *UserLoginPayload, dbUser *db.User) {

	dbUser.Email = user.Email
	dbUser.Password = user.Password

}

func mapUserPayloadToDbUserCreate(user *UserPayload, dbUser *db.User) {

	dbUser.Name = user.Name
	dbUser.Email = user.Email
	dbUser.Password = user.Password

}
