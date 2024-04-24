package response

type ErrorResponse struct {
	Timestamp   int64 `json:"timestamp"`
	Error string `json:"error"`
}