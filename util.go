package flutterwave

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"time"
)

func getGetTestKey() string {
	key := os.Getenv("FLW_SECK_KEY")

	if len(key) == 0 {
		key = "FLWSECK_TEST-SANDBOXDEMOKEY-X"
	}

	return key
}

func generateReference() string {
	timestamp := time.Now().UnixNano() // Get current timestamp in nanoseconds
	randomBytes := make([]byte, 4)
	_, err := rand.Read(randomBytes) // Generate random bytes
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("test-ref-%x-%x", timestamp, randomBytes) // Combine timestamp and random bytes
}

// Function to generate a random number with a given digit length
func generateRandomNumber(length int) string {
	if length <= 0 {
		return "" // Return 0 for invalid lengths
	}

	min := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(length-1)), nil)
	max := new(big.Int).Sub(new(big.Int).Mul(min, big.NewInt(10)), big.NewInt(1))

	n, err := rand.Int(rand.Reader, new(big.Int).Sub(max, min))
	if err != nil {
		panic(err)
	}

	return new(big.Int).Add(min, n).String()
}

func generateAccountNumber() string {
	return generateRandomNumber(10)
}
