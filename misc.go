package langfuse

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"
)

// ErrorResponse is the single-error response format.
type ErrorResponse struct {
	Error string `json:"error"`
}

// Err returns an error if the response contains an error message.
func (t ErrorResponse) Err() error {
	if len(t.Error) == 0 {
		return nil
	}
	return errors.New(t.Error)
}

// ErrorsResponse is the multi-error response format.
type ErrorsResponse struct {
	Errors []*APIError `json:"errors"`
}

// APIError is a single error entry.
type APIError struct {
	Field   *string `json:"field,omitempty"`
	Message *string `json:"message,omitempty"`
}

// Errs returns an aggregated error from the list.
func (t ErrorsResponse) Errs() error {
	s := []string{}
	for _, err := range t.Errors {
		var msg strings.Builder
		if err.Field != nil {
			msg.WriteString("field: ")
			msg.WriteString(*err.Field)
			msg.WriteString(", ")
		}
		if err.Message != nil {
			msg.WriteString("message: ")
			msg.WriteString(*err.Message)
		}
		s = append(s, msg.String())
	}

	if len(s) == 0 {
		return nil
	}

	return errors.New(strings.Join(s, ", "))
}

type statusCodeError struct {
	Code   int
	Status string
}

func (t statusCodeError) Error() string {
	return fmt.Sprintf("langfuse server error: %s", t.Status)
}

func (t statusCodeError) HTTPStatusCode() int {
	return t.Code
}

func checkStatusCode(resp *http.Response, d debug) error {
	if resp.StatusCode/100 == 2 {
		return nil
	}

	if err := logResponse(resp, d); err != nil {
		return err
	}

	// {"errors": [{"field": "...", "message": "..."}]}
	errorsResponse := new(ErrorsResponse)
	if err := newJSONParser(errorsResponse)(resp); err == nil {
		if errorsResponse.Errs() != nil {
			return errorsResponse.Errs()
		}
	}

	// {"error": "..."}
	errorResponse := new(ErrorResponse)
	if err := newJSONParser(errorResponse)(resp); err == nil {
		if errorResponse.Err() != nil {
			return errorResponse.Err()
		}
	}

	return statusCodeError{Code: resp.StatusCode, Status: resp.Status}
}

type responseParser func(*http.Response) error

func newJSONParser(dst interface{}) responseParser {
	return func(resp *http.Response) error {
		return json.NewDecoder(resp.Body).Decode(dst)
	}
}

func logResponse(resp *http.Response, d debug) error {
	if d.Debug() {
		text, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return err
		}
		d.Debugln(string(text))
	}

	return nil
}
