package flutterwave

import (
	"testing"
)

func TestVerifyTransactionById(t *testing.T) {
	transactionId := 8338738

	_, err := testClient.Transaction.VerifyTransactionById(transactionId)

	if err != nil {
		t.Error(err)
	}
}

func TestVerifyTransactionByRef(t *testing.T) {
	transactionRef := "MC-MC-1585230ew9v5050e8"

	_, err := testClient.Transaction.VerifyTransactionByRef(transactionRef)

	if err != nil {
		t.Error(err)
	}
}

func TestGetTransactions(t *testing.T) {
	_, err := testClient.Transaction.GetTransactions("")

	if err != nil {
		t.Error(err)
	}
}

func TestViewTransactionTimeline(t *testing.T) {
	transactionId := 8338738

	_, err := testClient.Transaction.ViewTransactionTimeline(transactionId)

	if err != nil {
		t.Error(err)
	}
}

func TestRefundTransaction(t *testing.T) {
	transactionId := 8338738
	body := map[string]interface{}{"amount": 5000, "comments": "test comment"}

	_, err := testClient.Transaction.RefundTransaction(transactionId, body)

	if err != nil {
		t.Error(err)
	}
}

func TestFetchRefundedTransaction(t *testing.T) {
	refundId := 82537

	_, err := testClient.Transaction.FetchRefundedTransaction(refundId)

	if err != nil {
		t.Error(err)
	}
}

func TestFetchMultipleRefundedTransactions(t *testing.T) {
	_, err := testClient.Transaction.FetchMultipleRefundedTransactions("")

	if err != nil {
		t.Error(err)
	}
}

func TestQueryTransactionFees(t *testing.T) {
	_, err := testClient.Transaction.QueryTransactionFees("")

	if err != nil {
		t.Error(err)
	}
}

func TestResendFailedWeebhook(t *testing.T) {
	transactionId := 8338738

	_, err := testClient.Transaction.ResendFailedWeebhook(transactionId, "")

	if err != nil {
		t.Error(err)
	}
}
