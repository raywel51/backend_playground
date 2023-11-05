package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Title       string             `form:"title" json:"title" bson:"title"`
	Author      string             `form:"author" json:"author" bson:"author"`
	Publication int                `form:"publication" json:"publication" bson:"publication"`
}
