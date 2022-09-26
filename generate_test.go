// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package goten_test

import (
	"github.com/gowizzard/goten"
	"reflect"
	"testing"
)

// TestGenerate is to test the generate function with table
// driven tests. We simulate the generation and check the result.
func TestGenerate(t *testing.T) {

	tests := []struct {
		length  int
		options goten.Options
	}{
		{
			length: 50,
			options: goten.Options{
				Uppercase: false,
				Numbers:   false,
				Symbols:   false,
			},
		},
		{
			length: 256,
			options: goten.Options{
				Uppercase: true,
				Numbers:   false,
				Symbols:   false,
			},
		},
		{
			length: 10,
			options: goten.Options{
				Uppercase: false,
				Numbers:   true,
				Symbols:   false,
			},
		},
		{
			length: 25,
			options: goten.Options{
				Uppercase: false,
				Numbers:   false,
				Symbols:   true,
			},
		},
		{
			length: 124,
			options: goten.Options{
				Uppercase: true,
				Numbers:   true,
				Symbols:   false,
			},
		},
		{
			length: 16,
			options: goten.Options{
				Uppercase: true,
				Numbers:   false,
				Symbols:   true,
			},
		},
		{
			length: 36,
			options: goten.Options{
				Uppercase: false,
				Numbers:   true,
				Symbols:   true,
			},
		},
		{
			length: 512,
			options: goten.Options{
				Uppercase: true,
				Numbers:   true,
				Symbols:   true,
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
