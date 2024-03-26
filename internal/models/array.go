package models

import (
	"github.com/go-playground/validator/v10"
	"testcase/internal/utils"
)

type GenerateRandomArrayRequest struct {
	MinValue      int64 `json:"min_value" validate:"required"`
	MaxValue      int64 `json:"max_value" validate:"required"`
	DesiredSize   int64 `json:"desired_size" validate:"required"`
	DesiredLength int64 `json:"desired_length" validate:"required"`
	Count         int64 `json:"count" validate:"required"`
}

func (r *GenerateRandomArrayRequest) ValidateArray() error {
	validate := validator.New()
	return utils.ValidateStruct(r, validate)
}

type GenerateRandomArrayResponse struct {
	RandomValues [][]int64 `json:"random_values"`
	Error        string    `json:"error,omitempty"`
}
