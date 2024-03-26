package models

import (
	"github.com/go-playground/validator/v10"
	"testcase/internal/utils"
)

type GenerateRandomIntegersRequest struct {
	MinValue      int64 `json:"min_value" validate:"required"`
	MaxValue      int64 `json:"max_value" validate:"required"`
	DesiredLength int64 `json:"desired_length" validate:"required,gte=1,omitempty"`
	Count         int64 `json:"count" validate:"required,gte=1,omitempty"`
}

func (r *GenerateRandomIntegersRequest) ValidateInteger() error {
	validate := validator.New()
	return utils.ValidateStruct(r, validate)
}

type GenerateRandomIntegersResponse struct {
	RandomValues []int64 `json:"random_values"`
	Error        string  `json:"error,omitempty"`
}
