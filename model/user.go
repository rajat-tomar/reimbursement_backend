package model

type User struct {
	Id    int    `json:"id,omitempty"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
