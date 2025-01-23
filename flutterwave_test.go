package flutterwave

var testClient *Client

func init() {
	secretKey := getGetTestKey()
	testClient = New(secretKey, nil)
}
