package errors

import (
	"encoding/json"
	"fmt"
)

type HTTPError struct {
	Cause   error  `json:"-"`
	Detail  string `json:"detail"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (e *HTTPError) Error() string {
	if e.Cause == nil {
		return e.Message
	}
	return e.Message + " : " + e.Cause.Error()
}

func (e *HTTPError) ResponseBody() ([]byte, error) {
	body, err := json.Marshal(e)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing response body: %v", err)
	}
	return body, nil
}

func NewHTTPError(err error, status int, message string) *HTTPError {
	httpError := &HTTPError{
		Cause:   err,
		Message: message,
		Status:  status,
	}

	if err != nil {
		httpError.Detail = err.Error()
	}

	return httpError
}
