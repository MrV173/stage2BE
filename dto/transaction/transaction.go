package transactiondto

import "landtick/models"

type CreateTransactionRequest struct {
	Image  string `json:"image" form:"image"`
	UserID int    `json:"user_id" form:"user_id"`
}

type TransactionResponse struct {
	Image  string              `json:"image"`
	UserID int                 `json:"user_id"`
	User   models.UserResponse `json:"user"`
}
