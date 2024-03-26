package services

import (
	"math/rand"
	"testcase/internal/models"
	"testcase/internal/utils"
)

func GenerateRandomArray(request models.GenerateRandomArrayRequest) (models.GenerateRandomArrayResponse, error) {
	if err := utils.ValidateArrayInputs(request.MinValue, request.MaxValue, request.DesiredSize, request.DesiredLength, request.Count); err != nil {
		return models.GenerateRandomArrayResponse{}, err
	}

	if err := utils.IsValidIntegerLength(request.MinValue, request.MaxValue, request.DesiredLength); err != nil {
		return models.GenerateRandomArrayResponse{}, err
	}

	if err := utils.DesireCountLimit(request.Count); err != nil {
		return models.GenerateRandomArrayResponse{}, err
	}

	result := make([][]int64, 0)

	for int64(len(result)) < request.Count {
		array := make([]int64, 0)
		for int64(len(array)) < request.DesiredSize {
			randomValue := rand.Int63n(request.MaxValue-request.MinValue+1) + request.MinValue
			if utils.CountIntegers(randomValue) == request.DesiredLength {
				array = append(array, randomValue)
			}
		}
		result = append(result, array)
	}

	response := models.GenerateRandomArrayResponse{RandomValues: result}

	return response, nil
}
