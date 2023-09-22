package main

import (
	"math/rand"
	"strings"
	"time"
)

const (
	// Set of characters to use in random URL generation
	charSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	// Length of the random URL path
	urlLength = 6
)

func GenerateRandomURL() string {
	// Create a new source and randomizer
	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)

	result := make([]byte, urlLength)

	for i := range result {
		result[i] = charSet[randomizer.Intn(len(charSet))]
	}

	return string(result)
}

func SafeRedirectURL(input string) string {
	// Check if the input starts with http:// or https://
	if !strings.HasPrefix(input, "http://") && !strings.HasPrefix(input, "https://") {
		// Default to HTTPS
		return "https://" + input
	}
	return input
}
