package services

import (
	"math/rand"
	"testcase/internal/models"
	"testcase/internal/utils"
)

func GenerateRandomIntegers(request models.GenerateRandomIntegersRequest) (models.GenerateRandomIntegersResponse, error) {
	if err := utils.ValidateMinAndMax(request.MinValue, request.MaxValue); err != nil {
		return models.GenerateRandomIntegersResponse{}, err
	}

	if err := utils.IsValidIntegerLength(request.MinValue, request.MaxValue, request.DesiredLength); err != nil {
		return models.GenerateRandomIntegersResponse{}, err
	}

	if err := utils.DesireCountLimit(request.Count); err != nil {
		return models.GenerateRandomIntegersResponse{}, err
	}

	result := make([]int64, 0)
	for int64(len(result)) < request.Count {
		randomValue := rand.Int63n(request.MaxValue-request.MinValue+1) + request.MinValue
		if utils.CountIntegers(randomValue) == request.DesiredLength {
			result = append(result, randomValue)
		}
	}
	response := models.GenerateRandomIntegersResponse{RandomValues: result}
	return response, nil
}
