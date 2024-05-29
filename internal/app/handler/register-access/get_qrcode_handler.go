package register_access

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"playground/internal/app/model/response"
	"playground/internal/app/repository"
)

func GetQrCode(c *gin.Context) {
	qrCodes, err := repository.FindHistoryQrCode()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var responseDataList []response.QRCode

	for idx, qrCode := range qrCodes {
		data := response.QRCode{
			ID:           idx + 1,
			CustomerName: qrCode.RealName + " " + qrCode.FamilyName,
			LicensePlate: qrCode.LicensePlate,
			PinCode:      qrCode.PinCode,
			QRImg:        os.Getenv("HOST_URI") + "/api/v1/gen-qr/" + qrCode.QrKey,
			QRKey:        qrCode.QrKey,
			RegisterTime: qrCode.RegisterTime.Format("2006-01-02 15:04"),
			RoomNumber:   qrCode.RoomNumber,
			CarType:      gin.H{"car_type": qrCode.CarType},
			VisitorType:  gin.H{"visitor_card": qrCode.VisitorType},
		}

		responseDataList = append(responseDataList, data)
	}

	c.JSON(http.StatusOK, struct {
		Status    bool              `json:"status"`
		Size      int               `json:"size"`
		MessageEN string            `json:"message_en"`
		MessageTH string            `json:"message_th"`
		Data      []response.QRCode `json:"data"`
	}{
		Status:    true,
		MessageEN: "",
		MessageTH: "",
		Size:      len(qrCodes),
		Data:      responseDataList,
	})
}
