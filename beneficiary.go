package flutterwave

import (
	"fmt"
	"net/http"
)

type BeneficiaryService service

type Beneficiary struct {
	ID            int                    `json:"id,omitempty"`
	AccountNumber string                 `json:"account_number,omitempty"`
	BankCode      string                 `json:"bank_code,omitempty"`
	FullName      string                 `json:"full_name,omitempty"`
	CreatedAt     string                 `json:"created_at,omitempty"`
	Meta          map[string]interface{} `json:"meta,omitempty"`
	BankName      string                 `json:"bank_name,omitempty"`
}

type BeneficiaryResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    Beneficiary `json:"data"`
}

type BeneficiaryListResponse struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Data    []Beneficiary `json:"data"`
	Meta    MetaData      `json:"meta,omitempty"`
}

type BeneficiaryRequest struct {
	AccountBank     string                 `json:"account_bank,omitempty"`
	AccountNumber   string                 `json:"account_number,omitempty"`
	Currency        string                 `json:"currency,omitempty"`
	BeneficiaryName string                 `json:"beneficiary_name,omitempty"`
	BankName        string                 `json:"name,omitempty"`
	Meta            map[string]interface{} `json:"meta,omitempty"`
}

type DeleteBeneficiaryResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"omitempty"`
}

// Create a beneficiary:
// https://developer.flutterwave.com/reference/create-a-beneficiary
//
// req receives the request body a type of BeneficiaryRequest
func (s *BeneficiaryService) CreateBeneficiary(req *BeneficiaryRequest) (BeneficiaryResponse, error) {
	var res BeneficiaryResponse
	err := s.client.makeRequest(http.MethodPost, "/beneficiaries", req, &res)
	return res, err
}

// Get all beneficiaries:
// https://developer.flutterwave.com/reference/list-all-beneficiaries
//
//	'params' == 'Query Params'
//	"to pass params for" page: GetBeneficiaries("page=1")
func (s *BeneficiaryService) GetBeneficiaries(params string) (BeneficiaryListResponse, error) {
	var res BeneficiaryListResponse
	url := "/beneficiaries"
	if params != "" {
		url = fmt.Sprintf("%s?%s", url, params)
	}
	err := s.client.makeRequest(http.MethodGet, url, nil, &res)
	return res, err
}

// Get a beneficiary:
// https://developer.flutterwave.com/reference/fetch-a-beneficiary
//
// id is the beneficiary id to fetch
func (s *BeneficiaryService) GetBeneficiary(id int) (BeneficiaryResponse, error) {
	var res BeneficiaryResponse
	url := fmt.Sprintf("/beneficiaries/%d", id)
	err := s.client.makeRequest(http.MethodGet, url, nil, &res)
	return res, err
}

// Delete a beneficiary:
// https://developer.flutterwave.com/reference/delete-a-beneficiary
//
// id is the beneficiary id to delete
func (s *BeneficiaryService) DeleteBeneficiary(id int) (DeleteBeneficiaryResponse, error) {
	var res DeleteBeneficiaryResponse
	url := fmt.Sprintf("/beneficiaries/%d", id)
	err := s.client.makeRequest(http.MethodDelete, url, nil, &res)
	return res, err
}
