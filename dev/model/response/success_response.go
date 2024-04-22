package response

type SuccessResponse struct {
	Mid 			string    	`json:"mid"`
	Text			string   	`json:"text"`
	Timestamp       int64 		`json:"timestamp"`
}