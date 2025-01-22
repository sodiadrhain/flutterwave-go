# flutterwave-go
Flutterwave go library is built for accessing the Flutterwave API

DOCUMENTATION: https://developer.flutterwave.com/docs/getting-started


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
import (
	"fmt"
	"net/http"

	"github.com/sodiadrhain/flutterwave-go"
)

secretKey := "FLWSECK_TEST-SANDBOXDEMOKEY-X"

client := flutterwave.New(apiKey, http.DefaultClient)

// Create a transfer
req := &TransferRequest{
	AccountBank:   "044",
	AccountNumber: "0690000040",
	Amount:        5000.00,
	Narration:     "Test transfers",
	Currency:      "NGN",
	Reference:     "test-ref-rfxx007",
	DebitCurrency: "NGN",
	Meta:          TransferMeta{Email: "test@tesuser.com", FirstName: "test" LastName: "user"},
}

transfer, err := client.Transfer.CreateTransfer(req)

if err != nil {
    // handle error
	fmt.Println(err)
}

// Get a transfer
transfer, err := client.Transfer.GetTransfer(transfer.ID)

fmt.Println(transfer)
```
See the test files for details and more example usage.


