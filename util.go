package flutterwave

import (
	"crypto/rand"
	"fmt"
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
