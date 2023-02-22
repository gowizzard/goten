// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package goten

import (
	"math/rand"
	"strings"
	"time"
)

// Options is to set the options for the token generator.
type Options struct {
	Uppercase bool
	Numbers   bool
	Symbols   bool
}

// letters, numbers, symbols are to save the information.
const (
	letters = "abcdefghijklmnopqrstuvwxyz"
	numbers = "1234567890"
	symbols = "!@#$%^&*"
)

// Generate is to generate a new token. The function check the
// parameter for the token and set every loop a new random
// seed based on the actual unix time number of nanoseconds.
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
		rand.NewSource(time.Now().UnixNano())
		token += string(char[rand.Intn(len(char))])
	}

	return token

}
