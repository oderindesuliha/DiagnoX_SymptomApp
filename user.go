package models

import "time"

type User struct {
	ID          string    `bson: "_id,omitempty" json:"id"`
	FirstName   string    `bson: "name" json:"first_name"`
	LastName    string    `bson: "last_name" json:"last_name"`
	Email       string    `bson: "email" json:"email"`
	PhoneNumber string    `bson: "phone_number" json:"phone_number"`
	Password    string    `bson: "password" json:"_"`
	CreatedAt   time.Time `bson: "created_at" json:"created_at"`
	verified    bool      `bson: "verified" json:"verified"`
}
