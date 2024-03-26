package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"testcase/internal/models"
	"testcase/internal/services"
)

func GenerateRandomArrayHandler(c *gin.Context) {
	var request models.GenerateRandomArrayRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := request.ValidateArray(); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	result, err := services.GenerateRandomArray(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
