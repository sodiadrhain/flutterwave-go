package flutterwave

import (
	"testing"
)

func TestGetBanks(t *testing.T) {
	countryCode := "NG"
	_, err := testClient.Bank.GetBanks(countryCode)

	if err != nil {
		t.Error(err)
	}
}

func TestGetBankBranches(t *testing.T) {
	bankCode := 280
	_, err := testClient.Bank.GetBankBranches(bankCode)

	if err != nil {
		t.Error(err)
	}
}
