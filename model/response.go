package model

type Response struct {
	Errors []error     `json:"errors,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}
