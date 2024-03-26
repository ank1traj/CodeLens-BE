package services

import (
	"math/rand"
	"testcase/internal/models"
	"testcase/internal/utils"
	"time"
)

func GenerateRandomString(request models.GenerateRandomStringsRequest) (models.GenerateRandomStringsResponse, error) {
	if err := utils.ValidateStringInput(request.Length, request.Count); err != nil {
		return models.GenerateRandomStringsResponse{}, err
	}

	if err := utils.DesireCountLimit(request.Count); err != nil {
		return models.GenerateRandomStringsResponse{}, err
	}


	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())

	result := make([]string, 0)
	for i := int64(0); i < request.Count; i++ {
		// Generate the random string of length request.Length
		var strBuilder string
		for j := int64(0); j < request.Length; j++ {
			randomValue := charSet[rand.Intn(len(charSet))]
			strBuilder += string(randomValue)
		}
		result = append(result, strBuilder)
	}

	response := models.GenerateRandomStringsResponse{RandomValues: result}
	return response, nil
}
