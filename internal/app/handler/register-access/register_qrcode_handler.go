package register_access

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"os"
	"strings"
	"time"

	"playground/internal/app/helper"
	"playground/internal/app/helper/credential-helper"
	"playground/internal/app/model/entity"
	"playground/internal/app/model/request"
	"playground/internal/app/model/response"
	"playground/internal/app/repository"
)

func GenQr(c *gin.Context) {

	var req request.GenQrRequest
	var err error

	s := c.Request.Header.Get("Authorization")

	token := strings.TrimPrefix(s, "Bearer ")

	if token == "" {
		c.JSON(401, gin.H{"error": "Token is missing"})
		return
	}

	if !credential_helper.BindRequest(c, &req) {
		return
	}

	if err := credential_helper.IsEmpty(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
		return
	}

	visitorType, err := repository.GetVisitorInfo(int(req.VisitorType))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
		return
	}

	carType, err := repository.GetCarTypeInfo(int(req.CarType))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
		return
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

	genQrCodeDao := entity.GenQrDao{
		ID:            primitiveId,
		IdentityId:    req.IdentityId,
		IdentityType:  int(req.IdentityType),
		RealName:      req.RealName,
		FamilyName:    req.FamilyName,
		Channel:       int(req.Channel),
		RoomNumber:    req.RoomNumber,
		VisitorType:   int(req.VisitorType),
		ProjectCode:   req.ProjectCode,
		Address:       req.Address,
		LicensePlate:  req.LicensePlate,
		CarType:       int(req.CarType),
		RegisterTime:  time.Now(),
		CheckoutTime:  nil,
		PaymentStatus: false,
		Approve:       false,
		PinCode:       pinCode,
		QrKey:         primitiveId.Hex(),
	}

	err = repository.InsertOneQrCode(&genQrCodeDao)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, struct {
		Status    bool            `json:"status"`
		MessageEN string          `json:"message_en"`
		MessageTH string          `json:"message_th"`
		Token     string          `json:"token"`
		Data      response.QRCode `json:"data"`
	}{
		Status:    true,
		MessageTH: "Registration successful",
		MessageEN: "Registration successful",
		Token:     token,
		Data: response.QRCode{
			CustomerName:  genQrCodeDao.RealName + " " + genQrCodeDao.FamilyName,
			LicensePlate:  genQrCodeDao.LicensePlate,
			PinCode:       genQrCodeDao.PinCode,
			QRImg:         os.Getenv("HOST_URI") + "/v2/gen-qr/" + primitiveId.Hex(),
			QRKey:         primitiveId.Hex(),
			RegisterTime:  genQrCodeDao.RegisterTime.Format("2006-01-02 15:04"),
			RoomNumber:    genQrCodeDao.RoomNumber,
			Approved:      genQrCodeDao.Approve,
			PaymentStatus: genQrCodeDao.PaymentStatus,
			CarType: gin.H{
				"car_type_id": carType.CarType,
				"car_type_th": carType.ValueTH,
				"car_type_en": carType.ValueEN,
				"car_type":    carType.ValueCN,
			},
			VisitorType: gin.H{
				"visitor_type_id": visitorType.IDField,
				"visitor_type_th": visitorType.ValueTH,
				"visitor_type_en": visitorType.ValueEN,
				"visitor_type_ch": visitorType.ValueCN,
				"visitor_card":    visitorType.VisitorCard,
			},
		},
	})
}
