package models

import (
	"github.com/go-playground/validator/v10"
	"testcase/internal/utils"
)

type GenerateRandomStringsRequest struct {
	Count  int64 `json:"count" validate:"required"`
	Length int64 `json:"length" validate:"required"`
}

func (r *GenerateRandomStringsRequest) ValidateString() error {
	validate := validator.New()
	return utils.ValidateStruct(r, validate)
}

type GenerateRandomStringsResponse struct {
	RandomValues []string `json:"random_values"`
	Error        string   `json:"error,omitempty"`
}
