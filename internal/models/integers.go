package models

type GenerateRandomIntegersRequest struct {
	MinValue      int64 `json:"min_value"`
	MaxValue      int64 `json:"max_value"`
	DesiredLength int64 `json:"desired_length"`
	Count         int64 `json:"count"`
}

type GenerateRandomIntegersResponse struct {
	RandomValues []int64 `json:"random_values"`
	Error        string  `json:"error,omitempty"`
}
