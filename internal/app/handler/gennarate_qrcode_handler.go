package handler

import (
	"bytes"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/nfnt/resize"
	"github.com/skip2/go-qrcode"

	"playground/internal/app/helper"
	"playground/internal/app/model/request"
)

func ReadQrCodeHandler(c *gin.Context) {
	url := c.Param("key")

	qrCode, err := qrcode.New("https://www.raywel.ddns.net/api-go/"+url, qrcode.Medium)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate QR code"})
		return
	}

	qrImg, err := qrCode.PNG(500)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate QR code image"})
		return
	}

	// Convert QR code byte slice to an image
	qrImage, _, err := image.Decode(bytes.NewReader(qrImg))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode QR code image"})
		return
	}

	// Load logo image from project path
	logoPath := "assets/img/logo.png" // Update with your image path
	logoBytes, err := readLogoFromFile(logoPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load logo image", "err": err})
		return
	}

	// Decode logo image
	logoImg, _, err := image.Decode(bytes.NewReader(logoBytes))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode logo image", "err": err.Error()})
		return
	}

	resizedLogo := resizeImage(logoImg, 253, 253)

	// Add logo to the center of the QR code
	qrWithLogo := addLogoToQRCode(qrImage, resizedLogo)

	// Encode the final image with logo to PNG format
	var buf bytes.Buffer
	if err := png.Encode(&buf, qrWithLogo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode final image"})
		return
	}

	fmt.Printf("gennarate qrcode -> : %s\n", url)
	c.Header("Content-Type", "image/png")
	c.Data(http.StatusOK, "image/png", buf.Bytes())
}

func CreateQrCode(c *gin.Context) {
	var req request.QrImgRequest

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

	if req.RawData == "" || req.Logo == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": "Username, Password, and Email are required"})
		return
	}

	key := helper.CreateQr(req.RawData, req.Logo)
	if key != "" {
		c.JSON(http.StatusCreated, gin.H{"status": true, "url": "http://localhost:1662/v2/gen-qr/" + key})
	}
}

func GetQrCode(c *gin.Context) {

}

func RemoveQrCode(c *gin.Context) {

}

func addLogoToQRCode(qrCode, logo image.Image) image.Image {

	offsetX := (qrCode.Bounds().Dx() - logo.Bounds().Dx()) / 2
	offsetY := (qrCode.Bounds().Dy() - logo.Bounds().Dy()) / 2

	// Create a new RGBA image to draw the QR code and logo onto it
	qrWithLogo := image.NewRGBA(qrCode.Bounds())

	// Draw the QR code onto the new image
	draw.Draw(qrWithLogo, qrWithLogo.Bounds(), qrCode, image.Point{}, draw.Over)

	// Calculate the position to draw the logo
	logoPosition := image.Point{X: offsetX, Y: offsetY}

	// Draw the logo onto the new image
	draw.Draw(qrWithLogo, logo.Bounds().Add(logoPosition), logo, image.Point{}, draw.Over)

	cropPx := 45
	bounds := qrWithLogo.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, bounds, qrWithLogo, bounds.Min, draw.Src)

	// Crop the combined QR code and logo image
	croppedImage := rgba.SubImage(image.Rect(cropPx, cropPx, rgba.Bounds().Dx()-cropPx, rgba.Bounds().Dy()-cropPx)).(*image.RGBA)

	return croppedImage
}

// Function to read logo image from file
func readLogoFromFile(path string) ([]byte, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	logoBytes, err := ioutil.ReadFile(absPath)
	if err != nil {
		return nil, err
	}

	return logoBytes, nil
}

func resizeImage(img image.Image, width, height int) image.Image {
	resized := resize.Resize(uint(width), uint(height), img, resize.Bilinear)

	return resized
}
