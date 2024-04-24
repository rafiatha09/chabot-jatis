package err_helper

import "errors"

var (
	ErrServiceRequestBody = errors.New("service: invalid request body")
	ErrServiceTimestamp = errors.New("service: timestamp is more than current time")
	ErrServiceMandatoryParameter = errors.New("service: invalid mandatory parameter")
	ErrServiceUserNotFound = errors.New("service: user not found")
	ErrServiceUnauthorized = errors.New("service: unauthorized")
)

// http 400 : invalid request body
// http 400 : timestamp is more than current time
// http 400 : invalid mandatory parameter
// http 400 : user not found
// http 401 : unauthorized