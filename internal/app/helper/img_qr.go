package helper

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"playground/internal/app/model/entity"
	"playground/internal/app/repository"
)

func CreateQr(raw, logo string) string {

	priId := primitive.NewObjectID()
	qr := entity.QrImgDao{
		ID:         priId,
		RawData:    raw,
		Logo:       logo,
		CreateTime: time.Now(),
	}

	err := repository.InsertImageQr(qr)
	if err != nil {
		return ""
	}

	return priId.Hex()
}
