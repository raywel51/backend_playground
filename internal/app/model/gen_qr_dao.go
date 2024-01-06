package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type GenQrDao struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	IdentityId    string             `bson:"identity_id"`
	IdentityType  int                `bson:"identity_type"`
	RealName      string             `bson:"real_name"`
	FamilyName    string             `bson:"family_name"`
	Channel       string             `bson:"channel"`
	RoomNumber    string             `bson:"room_number"`
	VisitorType   int                `bson:"visitor_type"`
	ProjectCode   string             `bson:"project_code"`
	Address       string             `bson:"address"`
	LicensePlate  string             `bson:"license_plate"`
	CarType       int                `bson:"car_type"`
	RegisterTime  time.Time          `bson:"register_time"`
	CheckoutTime  *time.Time         `bson:"check_out_time"`
	PaymentStatus bool               `bson:"payment_status"`
	PinCode       string             `bson:"pin_code"`
	QrKey         string             `bson:"qr_key"`
}
