package goten

import (
	"log"
	"testing"
)

// TestGenerate is to test the generate function and create a token
func TestGenerate(t *testing.T) {

	options := Options{
		Numbers: true,
		Symbols: false,
	}

	token := Generate(50, &options)
	log.Println(token)

}
