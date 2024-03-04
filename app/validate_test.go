package main

import (
	"testing"
)

func TestGenerateHMAC(t *testing.T) {
	t.Helper()

	cases := map[string]struct {
		secret   string
		message  string
		expected string
	}{
		// e.g. https://docs.github.com/en/webhooks/using-webhooks/validating-webhook-deliveries#testing-the-webhook-payload-validation
		"e.g. GitHub Webhook Docs": {
			secret:   "It's a Secret to Everybody",
			message:  "Hello, World!",
			expected: "757107ea0eb2509fc211221cce984b8a37570b6d7586c22c46f4379c8b043e17",
		},
		// e.g. https://www.php.net/manual/ja/function.hash-hmac.php
		"e.g. PHP Docs (hash-hmac)": {
			secret:   "secret",
			message:  "The quick brown fox jumped over the lazy dog.",
			expected: "9c5c42422b03f0ee32949920649445e417b2c634050833c5165704b825c2a53b",
		},
	}

	for _, c := range cases {
		actual := generateHMAC(c.secret, c.message)
		if actual != c.expected {
			t.Errorf(
				"generateHMAC(%s, %s): expected %s, actual %s",
				c.secret, c.message, c.expected, actual,
			)
		}
	}
}

func TestIsSameSignature(t *testing.T) {
	cases := map[string]struct {
		signatureA string
		signatureB string
		expected   bool
	}{
		"same": {
			signatureA: "757107ea0eb2509fc211221cce984b8a37570b6d7586c22c46f4379c8b043e17",
			signatureB: "757107ea0eb2509fc211221cce984b8a37570b6d7586c22c46f4379c8b043e17",
			expected:   true,
		},
		"different": {
			signatureA: "757107ea0eb2509fc211221cce984b8a37570b6d7586c22c46f4379c8b043e17",
			signatureB: "9c5c42422b03f0ee32949920649445e417b2c634050833c5165704b825c2a53b",
			expected:   false,
		},
	}

	for _, c := range cases {
		actual := isSameSignature(c.signatureA, c.signatureB)
		if actual != c.expected {
			t.Errorf(
				"isSameSignature(%s, %s): expected %t, actual %t",
				c.signatureA, c.signatureB, c.expected, actual,
			)
		}
	}
}
