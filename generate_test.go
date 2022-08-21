package goten_test

import (
	"github.com/gowizzard/goten"
	"reflect"
	"testing"
)

// TestGenerate is to test the generate function
// We simulate the generation and check the result
func TestGenerate(t *testing.T) {

	tests := []struct {
		length  int
		options goten.Options
	}{
		{
			length: 50,
			options: goten.Options{
				Numbers: false,
				Symbols: false,
			},
		},
		{
			length: 256,
			options: goten.Options{
				Numbers: true,
				Symbols: false,
			},
		},
		{
			length: 10,
			options: goten.Options{
				Numbers: false,
				Symbols: true,
			},
		},
		{
			length: 25,
			options: goten.Options{
				Numbers: true,
				Symbols: true,
			},
		},
	}

	for _, value := range tests {

		token := goten.Generate(value.length, &value.options)
		length := reflect.ValueOf(token).Len()

		if !reflect.DeepEqual(value.length, length) {
			t.Fatalf("expected: %d, got %d", value.length, length)
		}

	}

}
