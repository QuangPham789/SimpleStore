package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random String
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomUsername() string {
	return RandomString(6)
}

func RandomPassword() int64 {
	return RandomInt(100000, 999999)
}

func RandomProductName() string {
	return RandomString(8)
}

func RandomProductPrice() int64 {
	return RandomInt(100000, 999999)
}

func RandomProductCode() string {
	return RandomString(4)
}

// func HashPassword(password string) (string, error) {
// 	bytePassword := []byte(password)

// 	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)

// 	if err != nil {
// 		return "", err
// 	}

// 	return string(passwordHash), nil
// }
