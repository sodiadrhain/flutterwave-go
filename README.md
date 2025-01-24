# flutterwave-go
Flutterwave go library is built for accessing the Flutterwave API.

It is currenctly designed for the V3 APIs.

DOCUMENTATION: https://developer.flutterwave.com/docs


# Getting Started
- See references: 
 	- https://developer.flutterwave.com/docs/getting-started
	- https://developer.flutterwave.com/docs/authentication


# Installation
To install the flutterwave-go package, [Having go installed and setup locally](https://golang.org/):
1. Run command to install flutterwave-go
```sh
$ go get -u github.com/sodiadrhain/flutterwave-go
```
2. Import to your code:
```sh
import "github.com/sodiadrhain/flutterwave-go"
```

# Usage
```go
package main

import (
	"fmt"
	"net/http"

	"github.com/sodiadrhain/flutterwave-go"
)

func main() {
	secretKey := "FLWSECK_TEST-SANDBOXDEMOKEY-X"

	client := flutterwave.New(secretKey, http.DefaultClient)

	// Create a transfer
	req := &flutterwave.TransferRequest{
		AccountBank:   "044",
		AccountNumber: "0690000040",
		Amount:        5000.00,
		Narration:     "Test transfers",
		Currency:      "NGN",
		Reference:     "test-ref-fesaill07",
		DebitCurrency: "NGN",
		Meta:          flutterwave.TransferMeta{Email: "test@tesuser.com", FirstName: "test", LastName: "user"},
	}

	transfer, err := client.Transfer.CreateTransfer(req)

	if err != nil {
		fmt.Println(err)
	}

	// Get a transfer
	transfer, err = client.Transfer.GetTransfer(transfer.Data.ID)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(transfer)
}
```
See the test files for details and more example usage.


