package goten

import (
	"math/rand"
	"strings"
	"time"
)

// To save the letters, numbers, and symbols
const (
	letters = "abcdefghijklmnopqrstuvwxyz"
	numbers = "1234567890"
	symbols = "!@#$%^&*"
)

// Generate is to generate a new token
// The function check the parameter for the token
// And set every loop anew random seed
func Generate(length int, options *Options) string {

	char := letters
	if options != nil {
		if options.Uppercase {
			char += strings.ToUpper(letters)
		}
		if options.Numbers {
			char += numbers
		}
		if options.Symbols {
			char += symbols
		}
	}

	var token string
	for i := 0; i < length; i++ {
		rand.Seed(time.Now().UnixNano())
		token += string(char[rand.Intn(len(char))])
	}

	return token

}
