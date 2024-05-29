package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type TokenDao struct {
	ID       primitive.ObjectID `bson:"_id"`
	Token    string             `bson:"token"`
	Expiry   *time.Time         `bson:"expiry"`
	Username string             `bson:"username"`
}
