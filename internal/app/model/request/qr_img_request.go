package request

type QrImgRequest struct {
	RawData string `form:"raw_data" json:"raw_data" bson:"raw_data"`
	Logo    string `form:"logo" json:"logo" bson:"logo"`
}
