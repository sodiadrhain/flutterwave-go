package flutterwave

import (
	"strings"
	"testing"
)

var (
	accountBank           = "0477"
	accountNumber         = "0690000040"
	ammount               = 500.00
	narration             = "Test transfers"
	currency              = "NGN"
	debitCurrency         = "NGN"
	transferMetaEmail     = "test@testuser.com"
	transferMetaFirstName = "test"
	transferMetaLastName  = "user"
)

func createTransfer() (TransferResponse, error) {
	req := &TransferRequest{
		AccountBank:   accountBank,
		AccountNumber: accountNumber,
		Amount:        float64(ammount),
		Narration:     narration,
		Currency:      currency,
		Reference:     generateReference(),
		DebitCurrency: debitCurrency,
		Meta:          TransferMeta{Email: transferMetaEmail, FirstName: transferMetaFirstName, LastName: transferMetaLastName},
	}

	return testClient.Transfer.CreateTransfer(req)
}

func TestCreateTransfer(t *testing.T) {
	_, err := createTransfer()
	if err != nil {
		t.Error(err)
	}
}

func TestRetryFailedTransfer(t *testing.T) {
	// fetch failed transfers
	transfers, err := testClient.Transfer.GetTransfers("status=failed")
	if err != nil {
		t.Error(err)
	}

	transferId := transfers.Data[0].ID

	_, err = testClient.Transfer.RetryFailedTransfer(transferId)

	if err != nil {
		errorMsg := err.(*Errors).Err
		if !strings.Contains(errorMsg, "Cannot retry") {
			t.Error(err)
		}
	}
}

func createBulkTransfer() (TransferResponse, error) {
	bulkData := make([]TransferRequest, 2)
	req := &BulkTransferRequest{
		Currency:       "NGN",
		Title:          "March staff salary payment",
		ErrorReporting: false,
		BulkData: append(bulkData, TransferRequest{
			AccountBank:   accountBank,
			AccountNumber: accountNumber,
			Amount:        float64(ammount),
			Narration:     narration,
			Currency:      currency,
			Reference:     generateReference(),
			DebitCurrency: debitCurrency,
			Meta:          TransferMeta{Email: transferMetaEmail, FirstName: transferMetaFirstName, LastName: transferMetaLastName},
		}, TransferRequest{
			AccountBank:   accountBank,
			AccountNumber: accountNumber,
			Amount:        float64(ammount),
			Narration:     narration,
			Currency:      currency,
			Reference:     generateReference(),
			DebitCurrency: debitCurrency,
			Meta:          TransferMeta{Email: transferMetaEmail, FirstName: transferMetaFirstName, LastName: transferMetaLastName},
		}),
	}

	return testClient.Transfer.CreateBulkTransfer(req)
}

func TestCreateBulkTransfer(t *testing.T) {
	_, err := createBulkTransfer()
	if err != nil {
		t.Error(err)
	}
}

func TestQueryTransferFee(t *testing.T) {
	amount, currency, entity := 10000, "NGN", "account"

	_, err := testClient.Transfer.QueryTransferFee(amount, currency, entity)

	if err != nil {
		t.Error(err)
	}
}

func TestGetTransfers(t *testing.T) {
	_, err := testClient.Transfer.GetTransfers("page=1&status=failed&from=2023-11-18&to=2024-08-24&include_proof=false&include_provider_ref=false&include_approver_info=false&include_vat=false&include_date_completed=false&include_rate=false&include_debit_currency_amount=false")

	if err != nil {
		t.Error(err)
	}
}

func TestGetTransfer(t *testing.T) {
	transfer, err := createTransfer()
	if err != nil {
		t.Error(err)
	}
	transferId := transfer.Data.ID

	_, err = testClient.Transfer.GetTransfer(transferId)

	if err != nil {
		t.Error(err)
	}
}

func TestGetTransferRetry(t *testing.T) {
	transfer, err := createTransfer()
	if err != nil {
		t.Error(err)
	}
	transferId := transfer.Data.ID

	_, err = testClient.Transfer.GetTransferRetry(transferId)

	if err != nil {
		t.Error(err)
	}
}

func TestGetBulkTransfer(t *testing.T) {
	bulkTransfer, err := createBulkTransfer()
	if err != nil {
		t.Error(err)
	}

	_, err = testClient.Transfer.GetBulkTransfer(bulkTransfer.Data.ID)
	if err != nil {
		t.Error(err)
	}
}
