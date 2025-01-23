package flutterwave

import (
	"fmt"
	"net/http"
)

type TransactionService service

type Transaction struct {
	ID                int32                  `json:"id,omitempty"`
	TxRef             string                 `json:"tx_ref,omitempty"`
	FlwRef            string                 `json:"flw_ref,omitempty"`
	DeviceFingerprint string                 `json:"device_fingerprint,omitempty"`
	Amount            float64                `json:"amount,omitempty"`
	Currency          string                 `json:"currency,omitempty"`
	ChargedAmount     float64                `json:"charged_amount,omitempty"`
	AppFee            float64                `json:"app_fee,omitempty"`
	MerchantFee       float64                `json:"merchant_fee,omitempty"`
	ProcessorResponse string                 `json:"processor_response,omitempty"`
	AuthModel         string                 `json:"auth_model,omitempty"`
	IP                string                 `json:"ip,omitempty"`
	Narration         string                 `json:"narration,omitempty"`
	Status            string                 `json:"status,omitempty"`
	PaymentType       string                 `json:"payment_type,omitempty"`
	CreatedAt         string                 `json:"created_at,omitempty"`
	AccountId         int                    `json:"account_id,omitempty"`
	Card              map[string]interface{} `json:"card,omitempty"`
	Meta              map[string]interface{} `json:"meta,omitempty"`
	AmountSettled     float64                `json:"amount_settled,omitempty"`
	Customer          map[string]interface{} `json:"customer,omitempty"`
}

type TransactionResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    Transaction `json:"data"`
}

type TransactionListResponse struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Data    []Transaction `json:"data"`
	Meta    MetaData      `json:"meta,omitempty"`
}

// Verify transaction status:
// https://developer.flutterwave.com/reference/verify-transaction
//
// id is the transaction id to verify
func (ts *TransactionService) VerifyTransactionById(id int) (TransactionResponse, error) {
	var res TransactionResponse
	url := fmt.Sprintf("/transactions/%d/verify", id)
	err := ts.client.NewRequest(http.MethodGet, url, nil, &res)
	return res, err
}

// Verify transaction status with reference:
// https://developer.flutterwave.com/reference/verify-transaction-with-tx_ref
//
// txRef is the transaction reference to verify
func (ts *TransactionService) VerifyTransactionByRef(txRef string) (TransactionResponse, error) {
	var res TransactionResponse
	url := fmt.Sprintf("/transactions/verify_by_reference?tx_ref=%s", txRef)
	err := ts.client.NewRequest(http.MethodGet, url, nil, &res)
	return res, err
}

// Get multiple transactions:
// https://developer.flutterwave.com/reference/get-all-transactions
//
//	'params' == 'Query Params'
//	"to pass params for" tx_ref: GetTransactions("tx_ref=test-ref")
//	"to pass params for" tx_ref & page: GetTransactions("tx_ref=test-ref&page=1")
//	"to pass params for" tx_ref & page & status: GetTransactions("tx_ref=test-ref&page=1&status=successful")
func (ts *TransactionService) GetTransactions(params string) (TransactionListResponse, error) {
	var res TransactionListResponse
	url := "/transactions"
	if params != "" {
		url = fmt.Sprintf("%s?%s", url, params)
	}
	err := ts.client.NewRequest(http.MethodGet, url, nil, &res)
	return res, err
}

// View transaction timeline:
// https://developer.flutterwave.com/reference/get-transaction-events
//
// id is the transaction id to view events for
func (ts *TransactionService) ViewTransactionTimeline(id int) (ApiResponseList, error) {
	var res ApiResponseList
	url := fmt.Sprintf("/transactions/%d/events", id)
	err := ts.client.NewRequest(http.MethodGet, url, nil, &res)
	return res, err
}

// Refund a transaction:
// https://developer.flutterwave.com/reference/transaction-refund
//
// id is the transaction id to refund
//
// body is request body to pass, example usage:
//
//	body := map[string]interface{}{"amount": 5000, "comments": "test comment"}
func (ts *TransactionService) RefundTransaction(id int, body map[string]interface{}) (ApiResponse, error) {
	var res ApiResponse
	url := fmt.Sprintf("/transactions/%d/refund", id)
	err := ts.client.NewRequest(http.MethodPost, url, body, &res)
	return res, err
}

// Fetch a refunded transaction:
// https://developer.flutterwave.com/reference/get-transaction-refunds
//
// id is the refund id to fetch
func (ts *TransactionService) FetchRefundedTransaction(id int) (ApiResponse, error) {
	var res ApiResponse
	url := fmt.Sprintf("/transactions/refunds/%d", id)
	err := ts.client.NewRequest(http.MethodGet, url, nil, &res)
	return res, err
}

// Fetch multiple refunded transactions:
// https://developer.flutterwave.com/reference/get-all-refunds
//
//	'params' == 'Query Params'
//	"to pass params for" flw_ref: FetchMultipleRefundedTransactions("flw_ref=test-ref")
//	"to pass params for" flw_ref & page: FetchMultipleRefundedTransactions("flw_ref=test-ref&page=1")
//	"to pass params for" flw_ref & page & status: FetchMultipleRefundedTransactions("flw_ref=test-ref&page=1&status=successful")
func (ts *TransactionService) FetchMultipleRefundedTransactions(params string) (ApiResponseList, error) {
	var res ApiResponseList
	url := "/refunds"
	if params != "" {
		url = fmt.Sprintf("%s?%s", url, params)
	}
	err := ts.client.NewRequest(http.MethodGet, url, nil, &res)
	return res, err
}

// Query transaction fees:
// https://developer.flutterwave.com/reference/get-transaction-fee
//
//	'params' == 'Query Params'
//	"to pass params for" amount: QueryTransactionFees("amount=5000")
//	"to pass params for" amount & currency: QueryTransactionFees("amount=5000&currency=NGN")
func (ts *TransactionService) QueryTransactionFees(params string) (ApiResponse, error) {
	var res ApiResponse
	url := "/fee"
	if params != "" {
		url = fmt.Sprintf("%s?%s", url, params)
	}
	err := ts.client.NewRequest(http.MethodGet, url, nil, &res)
	return res, err
}

// Resend failed Webhooks:
// https://developer.flutterwave.com/reference/resend-transaction-webhook
//
// id is the transaction id to send for
//
//	'params' == 'Query Params'
//	"to pass params for" amount: ResendFailedWeebhook("amount=5000")
//	"to pass params for" amount & currency: ResendFailedWeebhook("amount=5000&currency=NGN")
func (ts *TransactionService) ResendFailedWeebhook(id int, params string) (ApiResponse, error) {
	var res ApiResponse
	url := fmt.Sprintf("/transactions/%d/resend-hook", id)
	if params != "" {
		url = fmt.Sprintf("%s?%s", url, params)
	}
	err := ts.client.NewRequest(http.MethodPost, url, nil, &res)
	return res, err
}
