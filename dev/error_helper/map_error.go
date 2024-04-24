package err_helper

import (
	"chatbot/dev/model/response"
	"chatbot/dev/util"
	"errors"
	"net/http"
)

func MapError(err error) (string, int, response.ErrorResponse) {
	var message string
	var code int
	var title string

	if errors.Is(err, ErrServiceRequestBody) { // ok
		title = "InvalidPayload"
		message = "invalid request body"
		code = http.StatusBadRequest
	} else if errors.Is(err, ErrServiceTimestamp) { // ok 
		title = "InvalidTimestamp"
		message = "timestamp is more than current time"
		code = http.StatusBadRequest
	} else if errors.Is(err, ErrServiceMandatoryParameter) { // ok 
		title = "InvalidMandatoryParameter"
		message = "invalid mandatory parameter"
		code = http.StatusBadRequest
	} else if errors.Is(err, ErrServiceUserNotFound) { // ok 
		title = "UserNotFound"
		message = "user not found"
		code = http.StatusBadRequest
	} else if errors.Is(err, ErrServiceUnauthorized) { // ok 
		title = "Unauthorized"
		message = "unauthorized"
		code = http.StatusBadRequest
	} else { // ok
		title = "InternalServerError"
		code = http.StatusInternalServerError
		message = "internal server error"
	}

	return title, code,  response.ErrorResponse {
		Timestamp: util.GenerateCurrentTimestamp(),
		Error: message,
	}
}

// http 400 : invalid request body
// http 400 : timestamp is more than current time
// http 400 : invalid mandatory parameter
// http 400 : user not found
// http 401 : unauthorized
// http 500 : internal server error