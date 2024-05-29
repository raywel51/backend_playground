package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"

	"playground/internal/app/helper/credential-helper"
	"playground/internal/app/model/request"
)

func RefreshHandler(c *gin.Context) {
	var jsonRequest MyJSONRequest

	if err := c.ShouldBindJSON(&jsonRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if jsonRequest.RefreshToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Refresh token is required"})
		return
	}

	token, err := jwt.Parse(jsonRequest.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		return credential_helper.RefreshSecret, nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	user := request.CredentialLoginRequest{Username: "exampleuser"}
	accessToken, err := credential_helper.CreateToken(user, credential_helper.SecretKey, time.Minute*15)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating access token"})
		return
	}

	// Respond with the new access token
	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}

type MyJSONRequest struct {
	RefreshToken string `json:"refresh_token"`
}
