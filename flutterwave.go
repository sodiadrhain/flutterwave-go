package flutterwave

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// Flutterwave api base url
const baseURL string = "https://api.flutterwave.com"

type service struct {
	client *Client
}

// Client object config
type Client struct {
	service    service
	httpClient *http.Client
	baseURL    string
	secretKey  string
	Transfer   *TransferService
}

// New initializes a new client config for communication with Flutterwave API
// using given secret key
func New(secretKey string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{Timeout: 60 * time.Second}
	}

	baseUrl := fmt.Sprintf("%s/v3", baseURL)
	client := &Client{
		httpClient: httpClient,
		secretKey:  secretKey,
		baseURL:    baseUrl,
		Transfer:   &TransferService{},
	}

	client.service.client = client
	client.Transfer = (*TransferService)(&client.service)
	return client
}

type ApiResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    MetaData    `json:"meta,omitempty"`
}

// Pagination meta data for paginated responses from the Flutterwave API
type PaginationMeta struct {
	Total      int `json:"total"`
	CurrenPage int `json:"current_page"`
	TotalPages int `json:"total_pages"`
}

// Meta data from the Flutterwave API
type MetaData struct {
	PageInfo PaginationMeta `json:"page_info,omitempty"`
}

func getGetTestKey() string {
	key := os.Getenv("FLW_SECK_KEY")

	if len(key) == 0 {
		key = "FLWSECK_TEST-SANDBOXDEMOKEY-X"
	}

	return key
}

// Make HTTP request to the Flutterwave API
func (c *Client) makeRequest(method, urlPath string, reqBody, v interface{}) error {
	newURL := c.baseURL + urlPath
	var body io.Reader

	if reqBody != nil {
		bb, err := json.Marshal(reqBody)
		if err != nil {
			return err
		}
		body = bytes.NewReader(bb)
	}

	req, err := http.NewRequest(method, newURL, body)
	if err != nil {
		return err
	}

	if method == http.MethodPost {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Authorization", "Bearer "+c.secretKey)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	return nil
}
