package flutterwave

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// Flutterwave api base url
const baseURL string = "https://api.flutterwave.com"

type service struct {
	client *Client
}

// Client object config
type Client struct {
	service     service
	httpClient  *http.Client
	baseURL     string
	secretKey   string
	Transfer    *TransferService
	Transaction *TransactionService
	Bank        *BankService
	Beneficiary *BeneficiaryService
}

// New initializes a new client config for communication with Flutterwave API
// using given secret key
func New(secretKey string, httpClient *http.Client) *Client {

	if httpClient == nil {
		httpClient = &http.Client{Timeout: 60 * time.Second}
	}

	client := &Client{
		httpClient: httpClient,
		secretKey:  secretKey,
		baseURL:    fmt.Sprintf("%s/v3", baseURL),
	}

	client.service.client = client
	client.Transfer = (*TransferService)(&client.service)
	client.Transaction = (*TransactionService)(&client.service)
	client.Bank = (*BankService)(&client.service)
	client.Beneficiary = (*BeneficiaryService)(&client.service)
	return client
}

// Response from the Flutterwave API
type ApiResponse struct {
	Status  string                 `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

// Response with list data from the Flutterwave API
type ApiResponseList struct {
	Status  string                   `json:"status"`
	Message string                   `json:"message"`
	Data    []map[string]interface{} `json:"data"`
	Meta    MetaData                 `json:"meta,omitempty"`
}

// Pagination meta data for paginated responses from the Flutterwave API
type PaginationMeta struct {
	Total      int `json:"total"`
	CurrenPage int `json:"current_page"`
	TotalPages int `json:"total_pages"`
	PageSize   int `json:"page_size,omitempty"`
}

// Meta data from the Flutterwave API
type MetaData struct {
	PageInfo PaginationMeta `json:"page_info,omitempty"`
}

// ErrorResponse from the Flutterwave API
type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// makeRequest makes HTTP request to the Flutterwave API
func (c *Client) makeRequest(method, urlPath string, reqBody, v interface{}) error {
	newURL := c.baseURL + urlPath
	var body io.Reader

	if reqBody != nil {
		bb, err := json.Marshal(reqBody)
		if err != nil {
			return hanldeError(UnknownError, err, nil)
		}
		body = bytes.NewReader(bb)
	}

	req, err := http.NewRequest(method, newURL, body)
	if err != nil {
		return hanldeError(RequestError, err, nil)
	}

	if method == http.MethodPost {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Authorization", "Bearer "+c.secretKey)

	req.Header.Set("accept", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return hanldeError(RequestError, err, nil)
	}

	defer res.Body.Close()

	if !strings.Contains(res.Header.Get("Content-Type"), "application/json") {
		return hanldeError(ApiError, fmt.Errorf("expected api response type of application/json, got %s", res.Header.Get("Content-Type")), nil)
	}

	if res.StatusCode >= 400 {
		return hanldeError(ApiError, nil, res)
	}

	if err := json.NewDecoder(res.Body).Decode(&v); err != nil {
		return hanldeError(UnknownError, err, nil)
	}

	return nil
}
