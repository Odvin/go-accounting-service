package util

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/google/uuid"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomName() string {
	return RandomString(6)
}

func RandomSurname() string {
	return RandomString(10)
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

// Random Admin status for the client
func RandomAdmStatus() string {
	adm := []string{"adm:active", "adm:blocked", "adm:suspended", "adm:processed"}
	n := len(adm)
	return adm[rand.Intn(n)]
}

// Random UUID
func RandomUUID() uuid.UUID {
	return uuid.New()
}

// Random KYC status for the client
func RandomKycStatus() string {
	kyc := []string{"kyc:unconfirmed", "kyc:confirmed", "kyc:pending", "kyc:rejected", "kyc:resubmission", "kyc:initiated"}
	n := len(kyc)
	return kyc[rand.Intn(n)]
}

// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
