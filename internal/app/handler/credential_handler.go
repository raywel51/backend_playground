package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"

	"playground/internal/app/helper"
	"playground/internal/app/model"
	"playground/internal/app/repository"
)

func UserLogin(c *gin.Context) {
	var req model.CredentialLoginRequest

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

	// Here you should verify the password against the hashed password in your User model
	// For example:
	if !verifyPassword(req.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"status": false, "message": "Invalid password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true})
}

func verifyPassword(inputPassword, storedPassword string) bool {
	key := helper.GetHashing(inputPassword)
	fmt.Println(key)
	return key == storedPassword
}

func UserRegister(c *gin.Context) {
	var req model.CredentialRegisterRequest

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

	// Create a user object
	user := model.User{
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

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "user": user})
}
