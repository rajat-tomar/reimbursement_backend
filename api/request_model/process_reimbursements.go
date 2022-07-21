package request_model

type ProcessReimbursementsRequest struct {
	Status string `json:"status"`
	UserId int    `json:"user_id"`
}
