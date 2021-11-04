//********************************************************************************************************************//
//
// This file is part of goten.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor (aka gowizzard)
//
//********************************************************************************************************************//

package goten

import (
	"fmt"
	"math/rand"
	"time"
)

// Characters for the token
const char = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOFQRSTUVWXYZ1234567890"

func init() {

	// Generate random seed, that the key is random
	// The seed based on the current time with nano seconds
	rand.Seed(time.Now().UnixNano())

}

// Generate a new funktion
func New(length int) string {

	// Variable for the user token
	var token string

	// Range for token
	for i := 0; i < length; i++ {

		// Building an string
		token = fmt.Sprintf("%s%s", token, string(char[rand.Intn(len(char))]))

	}

	// Return the token
	return token

}