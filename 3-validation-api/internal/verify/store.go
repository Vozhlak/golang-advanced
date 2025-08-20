package verify

import "time"

var verificationStore = make(map[string]string)

var verificationExpiry = make(map[string]time.Time)

func storeVerificationHash(hash, email string, duration time.Duration) {
	verificationStore[hash] = email
	verificationExpiry[hash] = time.Now().Add(duration)
}

func getEmailByHash(hash string) (string, bool) {
	expiry, found := verificationExpiry[hash]
	if !found {
		return "", false
	}
	if time.Now().After(expiry) {
		deleteHash(hash)
		return "", false
	}
	email, exists := verificationStore[hash]
	return email, exists
}

func deleteHash(hash string) {
	delete(verificationStore, hash)
	delete(verificationExpiry, hash)
}
