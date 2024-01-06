package model

type GenQrRequest struct {
	IdentityId   string `form:"identity_id" json:"identity_id" bson:"identity_id"`
	IdentityType int16  `form:"identity_type" json:"identity_type" bson:"identity_type"`
	RealName     string `form:"real_name" json:"real_name" bson:"real_name"`
	FamilyName   string `form:"family_name" json:"family_name" bson:"family_name"`
	Channel      string `form:"channel" json:"channel" bson:"channel"`
	RoomNumber   string `form:"room_number" json:"room_number" bson:"room_number"`
	VisitorType  int8   `form:"visitor_type" json:"visitor_type" bson:"visitor_type"`
	ProjectCode  string `form:"project_code" json:"project_code" bson:"project_code"`
	Address      string `form:"address" json:"address" bson:"address"`
	LicensePlate string `form:"license_plate" json:"license_plate" bson:"license_plate"`
	CarType      int8   `form:"car_type" json:"car_type" bson:"car_type"`
}
