package flutterwave

import (
	"testing"
)

func TestVerifyTransactionById(t *testing.T) {
	transactionId := 8338738

	transaction, err := testClient.Transaction.VerifyTransactionById(transactionId)

	if err != nil {
		t.Error(err)
	}

	if transaction.Status != "success" {
		t.Errorf("Verify transaction by ID error: Expected transfer status to be success, got %v; With message: \"%v\"", transaction.Status, transaction.Message)
	}
}

func TestVerifyTransactionByRef(t *testing.T) {
	transactionRef := "MC-MC-1585230ew9v5050e8"

	transaction, err := testClient.Transaction.VerifyTransactionByRef(transactionRef)

	if err != nil {
		t.Error(err)
	}

	if transaction.Status != "success" {
		t.Errorf("Verify transaction by reference error: Expected transfer status to be success, got %v; With message: \"%v\"", transaction.Status, transaction.Message)
	}
}

func TestGetTransactions(t *testing.T) {
	transactions, err := testClient.Transaction.GetTransactions("")

	if err != nil {
		t.Error(err)
	}

	if transactions.Status != "success" {
		t.Errorf("Get transactions error: Expected transfer status to be success, got %v; With message: \"%v\"", transactions.Status, transactions.Message)
	}
}

func TestViewTransactionTimeline(t *testing.T) {
	transactionId := 8338738

	transaction, err := testClient.Transaction.ViewTransactionTimeline(transactionId)

	if err != nil {
		t.Error(err)
	}

	if transaction.Status != "success" {
		t.Errorf("View transaction events error: Expected transfer status to be success, got %v; With message: \"%v\"", transaction.Status, transaction.Message)
	}
}

func TestRefundTransaction(t *testing.T) {
	transactionId := 8338738
	body := map[string]interface{}{"amount": 5000, "comments": "test comment"}

	refund, err := testClient.Transaction.RefundTransaction(transactionId, body)

	if err != nil {
		t.Error(err)
	}

	if refund.Status != "success" {
		t.Errorf("Refund transaction error: Expected transfer status to be success, got %v; With message: \"%v\"", refund.Status, refund.Message)
	}
}

func TestFetchRefundedTransaction(t *testing.T) {
	refundId := 82537

	refund, err := testClient.Transaction.FetchRefundedTransaction(refundId)

	if err != nil {
		t.Error(err)
	}

	if refund.Status != "success" {
		t.Errorf("Fetch refunded transaction error: Expected transfer status to be success, got %v; With message: \"%v\"", refund.Status, refund.Message)
	}
}

func TestFetchMultipleRefundedTransactions(t *testing.T) {
	transactions, err := testClient.Transaction.FetchMultipleRefundedTransactions("")

	if err != nil {
		t.Error(err)
	}

	if transactions.Status != "success" {
		t.Errorf("Fetch multiple refunded transactions error: Expected transfer status to be success, got %v; With message: \"%v\"", transactions.Status, transactions.Message)
	}
}

func TestQueryTransactionFees(t *testing.T) {
	fees, err := testClient.Transaction.QueryTransactionFees("")

	if err != nil {
		t.Error(err)
	}

	if fees.Status != "success" {
		t.Errorf("Fetch multiple refunded transactions error: Expected transfer status to be success, got %v; With message: \"%v\"", fees.Status, fees.Message)
	}
}

func TestResendFailedWeebhook(t *testing.T) {
	transactionId := 8338738

	resp, err := testClient.Transaction.ResendFailedWeebhook(transactionId, "")

	if err != nil {
		t.Error(err)
	}

	if resp.Status != "success" {
		t.Errorf("Resend failed webhook error: Expected transfer status to be success, got %v; With message: \"%v\"", resp.Status, resp.Message)
	}
}
