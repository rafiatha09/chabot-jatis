package request

type Request struct {
	UserID    string `json:"user_id" binding:"required"` // Mandatory
	Text      string `json:"text" binding:"required"`    // Mandatory
	Timestamp int64  `json:"timestamp,omitempty"` // Optional
}