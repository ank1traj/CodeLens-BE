package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"testcase/internal/models"
	"testcase/internal/services"
)

func GenerateRandomIntegersHandler(c *gin.Context) {
	var request models.GenerateRandomIntegersRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.GenerateRandomIntegers(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
