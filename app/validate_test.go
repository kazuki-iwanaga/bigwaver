package main

import (
	"testing"
)

func TestGenerateHMAC(t *testing.T) {
	patterns := []struct {
		secret   string
		message  string
		expected string
	}{
		// e.g. https://docs.github.com/en/webhooks/using-webhooks/validating-webhook-deliveries#testing-the-webhook-payload-validation
		{
			"It's a Secret to Everybody",
			"Hello, World!",
			"757107ea0eb2509fc211221cce984b8a37570b6d7586c22c46f4379c8b043e17",
		},
		// e.g. https://www.php.net/manual/ja/function.hash-hmac.php
		{
			"secret",
			"The quick brown fox jumped over the lazy dog.",
			"9c5c42422b03f0ee32949920649445e417b2c634050833c5165704b825c2a53b",
		},
	}

	for _, p := range patterns {
		actual := generateHMAC(p.secret, p.message)
		if actual != p.expected {
			t.Errorf(
				"generateHMAC(%s, %s): expected %s, actual %s",
				p.secret, p.message, p.expected, actual,
			)
		}
	}
}

func TestIsSameSignature(t *testing.T) {
	patterns := []struct {
		signatureA string
		signatureB string
		expected   bool
	}{
		{
			"757107ea0eb2509fc211221cce984b8a37570b6d7586c22c46f4379c8b043e17",
			"757107ea0eb2509fc211221cce984b8a37570b6d7586c22c46f4379c8b043e17",
			true,
		},
		{
			"757107ea0eb2509fc211221cce984b8a37570b6d7586c22c46f4379c8b043e17",
			"9c5c42422b03f0ee32949920649445e417b2c634050833c5165704b825c2a53b",
			false,
		},
		{
			"9c5c42422b03f0ee32949920649445e417b2c634050833c5165704b825c2a53b",
			"",
			false,
		},
	}

	for _, p := range patterns {
		actual := isSameSignature(p.signatureA, p.signatureB)
		if actual != p.expected {
			t.Errorf(
				"isSameSignature(%s, %s): expected %t, actual %t",
				p.signatureA, p.signatureB, p.expected, actual,
			)
		}
	}
}
