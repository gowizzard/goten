package goten

import (
	"math/rand"
	"time"
)

// To save the letters, numbers, and symbols
const (
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOFQRSTUVWXYZ"
	numbers = "1234567890"
	symbols = "!@#$%^&*"
)

// Options is to set the options for the token generator
type Options struct {
	Numbers bool
	Symbols bool
}

// Generate is to generate a new token
// The function check the parameter for the token
// And set every loop anew random seed
func Generate(length int, options *Options) string {

	char := letters
	if options != nil {
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
