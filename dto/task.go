package dto

type Task struct {
	UserID      uint   `json:"user_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description,omitempty"`
	Completed   bool   `json:"completed"`
}
