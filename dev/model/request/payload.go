package request

type Request struct {
	UserID    string    `json:"user_id"`
	Text      string    `json:"text"`
	Timestamp int64 	`json:"timestamp"`
}