package controller

type ErrorResponse struct {
	Reason  string `json:"reason"`
	Message string `json:"message"`
}
