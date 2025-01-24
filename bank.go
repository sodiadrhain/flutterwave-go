package flutterwave

import (
	"fmt"
	"net/http"
)

type BankService service

type Bank struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type BankResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    []Bank `json:"data"`
}

type BankBranch struct {
	ID         int    `json:"id"`
	BranchCode string `json:"branch_code"`
	BranchName string `json:"branch_name"`
	SwiftCode  string `json:"swift_code"`
	BIC        string `json:"bic"`
	BankId     int    `json:"bank_id"`
}

type BankBranchResponse struct {
	Status  string       `json:"status"`
	Message string       `json:"message"`
	Data    []BankBranch `json:"data"`
}

// Get all banks
// https://developer.flutterwave.com/reference/get-all-banks
//
// Pass either NG, GH, KE, UG, ZA or TZ to get list of banks
// in Nigeria, Ghana, Kenya, Uganda, South Africa or Tanzania respectively
func (s *BankService) GetBanks(country string) (BankResponse, error) {
	var res BankResponse
	url := fmt.Sprintf("/banks/%s", country)
	err := s.client.makeRequest(http.MethodGet, url, nil, &res)
	return res, err
}

// Get bank branches
// https://developer.flutterwave.com/reference/get-bank-branches
//
// id is bank id
func (s *BankService) GetBankBranches(id int) (BankBranchResponse, error) {
	var res BankBranchResponse
	url := fmt.Sprintf("/banks/%d/branches", id)
	err := s.client.makeRequest(http.MethodGet, url, nil, &res)
	return res, err
}
