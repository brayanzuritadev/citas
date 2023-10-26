package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FirstName string             `bson:"firstname" json:"firstname,omitempty"`
	LastName  string             `bson:"lastname" json:"lastname,omitempty"`
	DateBirth time.Time          `bson:"dateBirth" json:"dateBirth,omitempty"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password,omitempty"`
	Avatar    string             `bson:"avatar" json:"avatar,omitempty"`
	Banner    string             `bson:"banner" json:"banner,omitempty"`
	Profile   string             `bson:"profile" json:"profile,omitempty"`
	Address   string             `bson:"address" json:"address,omitempty"`
	Web       string             `bson:"web" json:"web,omitempty"`
}
