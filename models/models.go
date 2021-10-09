package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Creating Users and Posts struct
type Users struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Password string             `json:"pass,omitempty" bson:"pass,omitempty"`
	Posts    []string 		    `json:"posts" bson:"posts"`
}

type Posts struct {
	Id        primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Caption   string              `json:"caption,omitempty" bson:"caption,omitempty"`
	Img       string              `json:"img,omitempty" bson:"img,omitempty"`
	Timestamp primitive.Timestamp `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
}
