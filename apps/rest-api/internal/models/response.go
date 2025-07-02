package models

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type HealthCheckResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}
