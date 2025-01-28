package flutterwave

import (
	"fmt"
	"testing"
)

func createBeneficiary() (BeneficiaryResponse, error) {
	req := &BeneficiaryRequest{
		AccountBank:     accountBank,
		AccountNumber:   generateAccountNumber(),
		BeneficiaryName: fmt.Sprintf("%s %s", transferMetaFirstName, transferMetaLastName),
		Currency:        "GHS",
		BankName:        "ACCESS BANK GHANA",
	}

	return testClient.Beneficiary.CreateBeneficiary(req)
}

func TestCreateBeneficiary(t *testing.T) {
	_, err := createBeneficiary()
	if err != nil {
		t.Error(err)
	}
}

func TestGetBeneficiaries(t *testing.T) {
	_, err := testClient.Beneficiary.GetBeneficiaries("page=1")
	if err != nil {
		t.Error(err)
	}
}

func TestGetBeneficiary(t *testing.T) {
	ben, err := createBeneficiary()
	if err != nil {
		t.Error(err)
	}

	_, err = testClient.Beneficiary.GetBeneficiary(ben.Data.ID)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteBeneficiary(t *testing.T) {
	ben, err := createBeneficiary()
	if err != nil {
		t.Error(err)
	}

	_, err = testClient.Beneficiary.DeleteBeneficiary(ben.Data.ID)
	if err != nil {
		t.Error(err)
	}
}
