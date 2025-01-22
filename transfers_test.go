package flutterwave

import (
	"testing"
)

var (
	accountBank           = "044"
	accountNumber         = "0690000040"
	ammount               = 500.00
	narration             = "Test transfers"
	currency              = "NGN"
	reference             = "test-ref-rfxx007"
	debitCurrency         = "NGN"
	transferMetaEmail     = "test@testuser.com"
	transferMetaFirstName = "test"
	transferMetaLastName  = "user"
)

func TestCreateTransfer(t *testing.T) {
	req := &TransferRequest{
		AccountBank:   accountBank,
		AccountNumber: accountNumber,
		Amount:        float32(ammount),
		Narration:     narration,
		Currency:      currency,
		Reference:     reference,
		DebitCurrency: debitCurrency,
		Meta:          TransferMeta{Email: transferMetaEmail, FirstName: transferMetaFirstName, LastName: transferMetaLastName},
	}

	transfer, err := testClient.Transfer.CreateTransfer(req)

	if err != nil {
		t.Error(err)
	}

	if transfer.Status != "success" {
		t.Errorf("Create transfer error: Expected transfer status to be success, got %v; With message: \"%v\"", transfer.Status, transfer.Message)
	}

	if transfer.Data.ID == 0 {
		t.Error("Create transfer error:: Expected transfer id")
	}
}

func TestRetryFailedTransfer(t *testing.T) {
	transferId := 1196021

	transfer, err := testClient.Transfer.RetryFailedTransfer(transferId)

	if err != nil {
		t.Error(err)
	}

	if transfer.Status != "success" {
		t.Errorf("Rertry transfer error: Expected transfer status to be success, got %v; With message: \"%v\"", transfer.Status, transfer.Message)
	}

	if transfer.Data.ID == 0 {
		t.Error("Rertry transfer error:: Expected transfer id")
	}
}

func TestCreateBulkTransfer(t *testing.T) {
	bulkData := make([]TransferRequest, 2)
	req := &BulkTransferRequest{
		Currency:       "NGN",
		Title:          "March staff salary payment",
		ErrorReporting: false,
		BulkData: append(bulkData, TransferRequest{
			AccountBank:   accountBank,
			AccountNumber: accountNumber,
			Amount:        float32(ammount),
			Narration:     narration,
			Currency:      currency,
			Reference:     reference,
			DebitCurrency: debitCurrency,
			Meta:          TransferMeta{Email: transferMetaEmail, FirstName: transferMetaFirstName, LastName: transferMetaLastName},
		}, TransferRequest{
			AccountBank:   accountBank,
			AccountNumber: accountNumber,
			Amount:        float32(ammount),
			Narration:     narration,
			Currency:      currency,
			Reference:     reference,
			DebitCurrency: debitCurrency,
			Meta:          TransferMeta{Email: transferMetaEmail, FirstName: transferMetaFirstName, LastName: transferMetaLastName},
		}),
	}

	transfer, err := testClient.Transfer.CreateBulkTransfer(req)

	if err != nil {
		t.Error(err)
	}

	if transfer.Status != "success" {
		t.Errorf("Create bulk transfer error: Expected transfer status to be success, got %v; With message: \"%v\"", transfer.Status, transfer.Message)
	}

	if transfer.Data.ID == 0 {
		t.Error("Create bulk transfer error:: Expected transfer id")
	}
}

func TestQueryTransferFee(t *testing.T) {
	amount, currency, entity := 10000, "NGN", "account"

	fee, err := testClient.Transfer.QueryTransferFee(amount, currency, entity)

	if err != nil {
		t.Error(err)
	}

	if fee.Status != "success" {
		t.Errorf("Query transfer fee error: Expected transfer status to be success, got %v; With message: \"%v\"", fee.Status, fee.Message)
	}

}

func TestGetTransfers(t *testing.T) {
	transfers, err := testClient.Transfer.GetTransfers("page=1&status=failed&from=2023-11-18&to=2024-08-24&include_proof=false&include_provider_ref=false&include_approver_info=false&include_vat=false&include_date_completed=false&include_rate=false&include_debit_currency_amount=false")

	if err != nil {
		t.Error(err)
	}

	if transfers.Status != "success" {
		t.Errorf("Get transfers error: Expected transfer status to be success, got %v; With message: \"%v\"", transfers.Status, transfers.Message)
	}
}

func TestGetTransfer(t *testing.T) {
	transferId := 1196021

	transfer, err := testClient.Transfer.GetTransfer(transferId)

	if err != nil {
		t.Error(err)
	}

	if transfer.Status != "success" {
		t.Errorf("Get transfer error: Expected transfer status to be success, got %v; With message: \"%v\"", transfer.Status, transfer.Message)
	}

	if transfer.Data.ID == 0 {
		t.Error("Get transfer error: Expected transfer id")
	}
}

func TestGetTransferRetry(t *testing.T) {
	transferId := 1196021

	transfer, err := testClient.Transfer.GetTransferRetry(transferId)

	if err != nil {
		t.Error(err)
	}

	if transfer.Status != "success" {
		t.Errorf("Get transfer retry error: Expected transfer status to be success, got %v; With message: \"%v\"", transfer.Status, transfer.Message)
	}
}

func TestGetBulkTransfer(t *testing.T) {
	bulkId := 20514

	bulkTransfer, err := testClient.Transfer.GetTransferRetry(bulkId)

	if err != nil {
		t.Error(err)
	}

	if bulkTransfer.Status != "success" {
		t.Errorf("Get bulk transfer error: Expected transfer status to be success, got %v; With message: \"%v\"", bulkTransfer.Status, bulkTransfer.Message)
	}
}
