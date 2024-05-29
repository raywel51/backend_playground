package response

import (
	"github.com/gin-gonic/gin"
)

type QRCode struct {
	ID            int    `json:"id,omitempty"`
	CustomerName  string `json:"customer_name,omitempty"`
	LicensePlate  string `json:"license_plate,omitempty"`
	PinCode       string `json:"pin_code,omitempty"`
	QRImg         string `json:"qr_img,omitempty"`
	QRKey         string `json:"qr_key,omitempty"`
	RegisterTime  string `json:"register_time,omitempty"`
	RoomNumber    string `json:"room_number,omitempty"`
	Approved      bool   `json:"approved,omitempty"`
	PaymentStatus bool   `json:"payment_status,omitempty"`
	CarType       gin.H  `json:"car_type,omitempty"`
	VisitorType   gin.H  `json:"visitor_type,omitempty"`
}
