package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type QrImgDao struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	RawData    string             `form:"raw_data" json:"raw_data" bson:"raw_data"`
	Logo       string             `form:"logo" json:"logo" bson:"logo"`
	CreateTime time.Time          `form:"create_time" json:"create_time" bson:"create_time"`
}
