package models

type Transaction struct {
	ID     int          `json:"id" gorm:"primaryKey:autoIncrement"`
	Image  string       `json:"image" gorm:"type: varchar(255)"`
	UserID int          `json:"user_id"`
	User   UserResponse `json:"user"`
}

type TransactionResponse struct {
	ID     int          `json:"id"`
	Image  string       `json:"image"`
	User   UserResponse `json:"user"`
	UserID int          `json:"_"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}
