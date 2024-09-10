package models

// Expense represents an expense in the application.
type Expense struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	PaidBy      string  `json:"paid_by"`
	GroupID     int     `json:"group_id"`
}
