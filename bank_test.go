package flutterwave

import (
	"testing"
)

func TestGetBanks(t *testing.T) {
	countryCode := "NG"
	banks, err := testClient.Bank.GetBanks(countryCode)

	if err != nil {
		t.Error(err)
	}

	if banks.Status != "success" {
		t.Errorf("Get banks error: Expected transfer status to be success, got %v; With message: \"%v\"", banks.Status, banks.Message)
	}
}

func TestGetBankBranches(t *testing.T) {
	bankCode := 280
	branches, err := testClient.Bank.GetBankBranches(bankCode)

	if err != nil {
		t.Error(err)
	}

	if branches.Status != "success" {
		t.Errorf("Get bank branches error: Expected transfer status to be success, got %v; With message: \"%v\"", branches.Status, branches.Message)
	}
}
