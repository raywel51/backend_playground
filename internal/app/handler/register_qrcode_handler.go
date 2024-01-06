package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"os"
	"strings"
	"time"

	"playground/internal/app/helper"
	"playground/internal/app/model"
	"playground/internal/app/repository"
)

func GenQr(c *gin.Context) {

	var req model.GenQrRequest
	var err error

	s := c.Request.Header.Get("Authorization")

	token := strings.TrimPrefix(s, "Bearer ")

	if token == "" {
		c.JSON(401, gin.H{"error": "Token is missing"})
		return
	}

	if c.ContentType() == "application/x-www-form-urlencoded" {
		if err := c.ShouldBindWith(&req, binding.Form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": err.Error()})
			return
		}
	} else {
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": err.Error()})
			return
		}
	}

	if req.IdentityId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": "missing identity"})
		return
	}

	if req.IdentityType == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": "missing identity"})
		return
	}

	if req.VisitorType == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": "missing Visitor type"})
		return
	}
	visitorType, err := repository.GetVisitorInfo(int(req.VisitorType))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
		return
	}

	if req.CarType == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": "missing car type"})
	}

	var pinCode string
	for {
		pin, err := helper.GeneratePIN()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
			return
		}

		qrDao, err := repository.CheckPin(pin)
		if qrDao == nil {
			pinCode = pin
			break
		}
	}

	var primitiveId = primitive.NewObjectID()

	genQrCodeDao := model.GenQrDao{
		ID:            primitiveId,
		IdentityId:    req.IdentityId,
		IdentityType:  int(req.IdentityType),
		RealName:      req.RealName,
		FamilyName:    req.FamilyName,
		Channel:       req.Channel,
		RoomNumber:    req.RoomNumber,
		VisitorType:   int(req.VisitorType),
		ProjectCode:   req.ProjectCode,
		Address:       req.ProjectCode,
		LicensePlate:  req.LicensePlate,
		CarType:       int(req.CarType),
		RegisterTime:  time.Now(),
		CheckoutTime:  nil,
		PaymentStatus: false,
		PinCode:       pinCode,
		QrKey:         primitiveId.Hex(),
	}

	err = repository.RegisterQrCode(&genQrCodeDao)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": true,
		"data": gin.H{
			"customer_name": genQrCodeDao.RealName + " " + genQrCodeDao.FamilyName,
			"pin_code":      genQrCodeDao.PinCode,
			"qr_key":        primitiveId,
			"register_time": genQrCodeDao.RegisterTime,
			"room_number":   genQrCodeDao.RoomNumber,
			"license_plate": genQrCodeDao.LicensePlate,
			"qr_img":        os.Getenv("HOST_URI") + "/v2/gen-qr/" + primitiveId.Hex(),
			"visitorType": gin.H{
				"visitor_type_id": visitorType.IDField,
				"visitor_type_th": visitorType.ValueTH,
				"visitor_type_en": visitorType.ValueEN,
				"visitor_type_ch": visitorType.ValueCN,
				"visitor_card":    visitorType.VisitorCard,
			},
		},
		"message_en": "Registration successful",
		"message_th": "ลงทะเบียนสำเร็จ",
		"token":      token,
	})
}
