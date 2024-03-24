package models

type GenerateRandomStringsRequest struct {
	Count  int64 `json:"count"`
	Length int64 `json:"length"`
}

type GenerateRandomStringsResponse struct {
	RandomValues []string `json:"random_values"`
	Error        string   `json:"error,omitempty"`
}
