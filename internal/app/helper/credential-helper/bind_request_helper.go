package credential_helper

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func BindRequest(c *gin.Context, req interface{}) bool {
	if c.ContentType() == "application/x-www-form-urlencoded" {
		if err := c.ShouldBindWith(req, binding.Form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": err.Error()})
			return false
		}
	} else {
		if err := c.ShouldBindJSON(req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": err.Error()})
			return false
		}
	}

	return true
}
