package response

type ErrorResponse struct {
	Code    int    `json:"code"`
	Timestamp   int64 `json:"timestamp"`
	Error string `json:"error"`
}