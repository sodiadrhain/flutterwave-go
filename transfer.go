package flutterwave

import (
	"fmt"
	"net/http"
)

type TransferService service

type Transfer struct {
	ID               int32        `json:"id,omitempty"`
	AccountNumber    string       `json:"account_number,omitempty"`
	BankCode         string       `json:"bank_code,omitempty"`
	FullName         string       `json:"full_name,omitempty"`
	CreatedAt        string       `json:"created_at,omitempty"`
	Currency         string       `json:"currency,omitempty"`
	DebitCurrency    string       `json:"debit_currency,omitempty"`
	Amount           float64      `json:"amount,omitempty"`
	Fee              float64      `json:"fee,omitempty"`
	Status           string       `json:"status,omitempty"`
	Reference        string       `json:"reference,omitempty"`
	Meta             TransferMeta `json:"meta,omitempty"`
	Narration        string       `json:"narration,omitempty"`
	Approver         string       `json:"approver,omitempty"`
	CompleteMessage  string       `json:"complete_message,omitempty"`
	RequiresApproval int          `json:"requires_approval,omitempty"`
	IsApproved       int          `json:"is_approved,omitempty"`
	BankName         string       `json:"bank_name,omitempty"`
	Vat              float64      `json:"vat,omitempty"`
	Rate             float64      `json:"rate,omitempty"`
}

type TransferMeta struct {
	Sender                        string `json:"sender,omitempty"`
	SenderAddress                 string `json:"sender_address,omitempty"`
	SenderCity                    string `json:"sender_city,omitempty"`
	SenderCountry                 string `json:"sender_country,omitempty"`
	SenderIdNumber                string `json:"sender_id_number,omitempty"`
	SenderIdType                  string `json:"sender_id_type,omitempty"`
	SenderIdExpiry                string `json:"sender_id_expiry,omitempty"`
	SenderMobileNumber            string `json:"sender_mobile_number,omitempty"`
	SenderEmailAddress            string `json:"sender_email_address,omitempty"`
	SenderOccupation              string `json:"sender_occupation,omitempty"`
	SenderBeneficiaryRelationship string `json:"sender_beneficiary_relationship,omitempty"`
	FirstName                     string `json:"first_name,omitempty"`
	LastName                      string `json:"last_name,omitempty"`
	Email                         string `json:"email,omitempty"`
	MobileNumber                  string `json:"mobile_number,omitempty"`
	BeneficiaryMobileNumber       string `json:"beneficiary_mobile_number,omitempty"`
	RecipientAddress              string `json:"recipient_address,omitempty"`
	BeneficiaryState              string `json:"beneficiary_state,omitempty"`
	BeneficiaryCountry            string `json:"beneficairy_country,omitempty"`
	TransferPuurpose              string `json:"transfer_purpose,omitempty"`
	RoutingNumber                 string `json:"routing_number,omitempty"`
	MerchantName                  string `json:"merchant_name,omitempty"`
	AccountNumber                 string `json:"account_number,omitempty"`
}

type TransferRequest struct {
	AccountBank           string       `json:"account_bank,omitempty"`
	AccountNumber         string       `json:"account_number,omitempty"`
	Amount                float64      `json:"amount,omitempty"`
	Currency              string       `json:"currency,omitempty"`
	DebitSubAccount       string       `json:"debit_sub_account,omitempty"`
	Beneficiary           string       `json:"beneficiary,omitempty"`
	Reference             string       `json:"reference,omitempty"`
	DebitCurrency         string       `json:"debit_currency,omitempty"`
	DestinationBranchCode string       `json:"destination_branch_code,omitempty"`
	CallBackUrl           string       `json:"call_back_url,omitempty"`
	Narration             string       `json:"narration,omitempty"`
	Meta                  TransferMeta `json:"meta,omitempty"`
}

type TransferResponse struct {
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Data    Transfer `json:"data"`
}

type BulkTransferRequest struct {
	Currency       string            `json:"currency,omitempty"`
	Title          string            `json:"title,omitempty"`
	ErrorReporting bool              `json:"error_reporting,omitempty"`
	BulkData       []TransferRequest `json:"bulk_data,omitempty"`
}

type TransferListResponse struct {
	Status  string     `json:"status"`
	Message string     `json:"message"`
	Data    []Transfer `json:"data"`
	Meta    MetaData   `json:"meta,omitempty"`
}

// Create a transfer:
// https://developer.flutterwave.com/reference/create-a-transfer
//
// req receives the request body a type of TransferRequest
func (ts *TransferService) CreateTransfer(req *TransferRequest) (TransferResponse, error) {
	var res TransferResponse
	err := ts.client.NewRequest(http.MethodPost, "/transfers", req, &res)
	return res, err
}

// Retry a failed transfer:
// https://developer.flutterwave.com/reference/transfer-retry
//
// id is the transfer id to retry
func (ts *TransferService) RetryFailedTransfer(id int) (TransferResponse, error) {
	var res TransferResponse
	url := fmt.Sprintf("/transfers/%d/retries", id)
	err := ts.client.NewRequest(http.MethodPost, url, nil, &res)
	return res, err
}

// Create bulk transfer:
// https://developer.flutterwave.com/reference/create-bulk-transfer
func (ts *TransferService) CreateBulkTransfer(req *BulkTransferRequest) (TransferResponse, error) {
	var res TransferResponse
	err := ts.client.NewRequest(http.MethodPost, "/bulk-transfers", req, &res)
	return res, err
}

// Query transfer fee:
// https://developer.flutterwave.com/reference/get-transfer-fee
//
// entity == type, this is the type of transfer you want to get the fee for.
// expected values are mobilemoney and account.
func (ts *TransferService) QueryTransferFee(amount int, curreny string, entity string) (ApiResponseList, error) {
	var res ApiResponseList
	url := fmt.Sprintf("/transfers/fee?amount=%d&currency=%s&type=%s", amount, curreny, entity)
	err := ts.client.NewRequest(http.MethodGet, url, nil, &res)
	return res, err
}

// Get all transfers:
// https://developer.flutterwave.com/reference/get-all-transfers
//
//	'params' == 'Query Params'
//	"to pass params for" reference: GetTransfers("reference=test-ref")
//	"to pass params for" reference & page: GetTransfers("reference=test-ref&page=1")
//	"to pass params for" reference & page & status: GetTransfers("reference=test-ref&page=1&status=successful")
func (ts *TransferService) GetTransfers(params string) (TransferListResponse, error) {
	var res TransferListResponse
	url := "/transfers"
	if params != "" {
		url = fmt.Sprintf("/transfers?%s", params)
	}
	err := ts.client.NewRequest(http.MethodGet, url, nil, &res)
	return res, err
}

// Get a transfer:
// https://developer.flutterwave.com/reference/get-a-transfer
//
// id is the transfer id to fetch
func (ts *TransferService) GetTransfer(id int) (TransferResponse, error) {
	var res TransferResponse
	url := fmt.Sprintf("/transfers/%d", id)
	err := ts.client.NewRequest(http.MethodGet, url, nil, &res)
	return res, err
}

// Get a transfer retry:
// https://developer.flutterwave.com/reference/get-a-transfer-retry
//
// id is the transfer id to fetch retries for
func (ts *TransferService) GetTransferRetry(id int) (ApiResponseList, error) {
	var res ApiResponseList
	url := fmt.Sprintf("/transfers/%d/retries", id)
	err := ts.client.NewRequest(http.MethodGet, url, nil, &res)
	return res, err
}

// Get a bulk transfer:
// https://developer.flutterwave.com/reference/get-a-bulk-transfer
//
// id is the bulk transfer id to fetch
func (ts *TransferService) GetBulkTransfer(id int) (ApiResponseList, error) {
	var res ApiResponseList
	url := fmt.Sprintf("/transfers?batch_id=%d", id)
	err := ts.client.NewRequest(http.MethodGet, url, nil, &res)
	return res, err
}
