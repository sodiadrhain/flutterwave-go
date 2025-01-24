package flutterwave

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const (
	ApiError     = "API_ERROR"
	RequestError = "REQUEST_ERROR"
	UnknownError = "UNKNOWN_ERROR"
)

// This contains the general errors
type Errors struct {
	Name    string       `json:"name"`
	Err     string       `json:"error"`
	Details ErrorDetails `json:"details"`
}

type ErrorDetails struct {
	StatusCode int      `json:"status_code"`
	Status     string   `json:"status"`
	URL        *url.URL `json:"url"`
	Message    string   `json:"message"`
}

func hanldeError(errType string, err error, resp *http.Response) *Errors {
	var (
		errorResp ErrorResponse
		details   ErrorDetails
		errorMsg  string
	)

	if err != nil {
		errorMsg = err.Error()
	}

	if resp != nil {
		respBody, _ := io.ReadAll(resp.Body)
		_ = json.Unmarshal(respBody, &errorResp)

		details = ErrorDetails{
			StatusCode: resp.StatusCode,
			Status:     errorResp.Status,
			URL:        resp.Request.URL,
			Message:    errorResp.Message,
		}

		errorMsg = errorResp.Message
	}

	return &Errors{
		Name:    errType,
		Err:     errorMsg,
		Details: details,
	}
}

func (err *Errors) Error() string {
	res, _ := json.Marshal(err)
	return string(res)
}
