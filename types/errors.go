package types

import "fmt"

var ErrOverMaxRetries = NewGnsError("reached maximum retries")

type GnsError struct {
	Message string `json:"message"`
}

func (e *GnsError) Error() string {
	return e.Message
}

func NewGnsError(message string, fields ...any) *GnsError {
	return &GnsError{Message: fmt.Sprintf(message, fields...)}
}
