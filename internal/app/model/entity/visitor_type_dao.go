package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type VisitorTypeDao struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	IDField     int                `bson:"id"`
	VisitorType int                `bson:"visitor_type"`
	ValueTH     string             `bson:"value_th"`
	ValueEN     string             `bson:"value_en"`
	ValueCN     string             `bson:"value_cn"`
	VisitorCard int                `bson:"visitor_card"`
}
