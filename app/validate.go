package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
)

func generateHMAC(secret, message string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(message))
	return hex.EncodeToString(mac.Sum(nil))
}

func isSameSignature(signatureA, signatureB string) bool {
	// Use subtle.ConstantTimeCompare to prevent timing attacks
	// https://docs.github.com/en/webhooks/using-webhooks/validating-webhook-deliveries#validating-webhook-deliveries
	return subtle.ConstantTimeCompare([]byte(signatureA), []byte(signatureB)) == 1
}
