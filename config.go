// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Package goten is there to provide a way to generate
// a random token with multiple options.
package goten

// Options is to set the options for the token generator.
type Options struct {
	Uppercase bool
	Numbers   bool
	Symbols   bool
}
