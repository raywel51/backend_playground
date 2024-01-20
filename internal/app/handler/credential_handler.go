package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"

	"playground/internal/app/helper"
	credential_helper "playground/internal/app/helper/credential-helper"
	"playground/internal/app/model/entity"
	"playground/internal/app/model/request"
	"playground/internal/app/repository"
)

func UserLogin(c *gin.Context) {
	var req request.CredentialLoginRequest

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

	if req.Username == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": "Username and Password are required"})
		return
	}

	user, err := repository.SelectOneUserByUsername(req.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid credentials"})
		return
	}

	if !helper.VerifyPassword(req.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"status": false, "message": "Invalid password"})
		return
	}

	accessToken, err := credential_helper.CreateToken(req, credential_helper.SecretKey, time.Minute*15)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating access token"})
		return
	}

	expiryTime := time.Now().Add(12 * time.Hour)

	token := entity.TokenDao{
		ID:       primitive.NewObjectID(),
		Token:    accessToken,
		Expiry:   &expiryTime,
		Username: req.Username,
	}

	err = repository.InsertOrUpdateToken(&token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Login successful",
		"token":   accessToken,
	})
}

func UserRegister(c *gin.Context) {
	var req request.CredentialRegisterRequest

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

	if req.Username == "" || req.Password == "" || req.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": "Username, Password, and Email are required"})
		return
	}

	userChecker, _ := repository.SelectOneUserByUsername(req.Username)
	if userChecker != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Username already exists in the system."})
		return
	}

	emailChecker, _ := repository.SelectOneUserByEmail(req.Email)
	if emailChecker != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Email already exists in the system."})
		return
	}

	hashedString := helper.GetHashing(req.Password) // get password with sha256 algorithm

	user := entity.UserDao{
		ID:       primitive.NewObjectID(),
		Username: req.Username,
		Password: hashedString,
		Email:    req.Email,
	}

	err := repository.InsertUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "UserDao registered successfully", "user": user})
}

func UserReadAll(c *gin.Context) {
	users, err := repository.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": true, "count": len(users), "data": users})
}

func UserReadOneById(c *gin.Context) {
	id := c.Param("id")

	users, err := repository.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"status": true, "data": users})
}

func UserDeleteById(c *gin.Context) {
	id := c.Param("id")

	err := repository.DeleteUserById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"status": true})
}
