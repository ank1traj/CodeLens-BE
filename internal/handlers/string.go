package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"testcase/internal/models"
	"testcase/internal/services"
)

func GenerateRandomStringsHandler(c *gin.Context) {
	var request models.GenerateRandomStringsRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.GenerateRandomString(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
