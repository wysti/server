package rest

import ()

type RestError struct {
	err string
}

func (re RestError) Error() string {
	return re.err
}

type ErrorCode int

const (
	UNKNOWN ErrorCode = iota
	REQUEST_PARSE
	JSON_MARSHALLING
	INVALID_PARAMETER
	GENERIC_SERVER
)

var ErrorMessages map[ErrorCode]string

//Internationalize this?
func InitializeErrors() {
	ErrorMessages = make(map[ErrorCode]string, 0)
	ErrorMessages[UNKNOWN] = "Unknown error"
	ErrorMessages[REQUEST_PARSE] = "Error parsing request"
	ErrorMessages[JSON_MARSHALLING] = "Error rendering result as json"
	ErrorMessages[INVALID_PARAMETER] = "Error, invalid parameter"
	ErrorMessages[GENERIC_SERVER] = "Generic server error"
}
