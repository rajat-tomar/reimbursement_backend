package model

type Response struct {
	Status string      `json:"status,omitempty"`
	Errors []error     `json:"errors,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}
