// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package color

const (
	// bg-base-100.
	Base100 BaseColor = iota
	// bg-base-200.
	Base200
	// bg-base-300.
	Base300
)

// BaseColor represents a `bg-base-*` value.
type BaseColor int
