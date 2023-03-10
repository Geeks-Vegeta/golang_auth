package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Email    string             `json:"email,omitempty" validate:"required"`
	Password string             `json:"password,omitempty" validate:"required"`
}
