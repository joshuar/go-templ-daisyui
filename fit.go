// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

//go:generate go run golang.org/x/tools/cmd/stringer -type=ObjectFit -linecomment -output fit_generated.go
package components

const (
	ObjectNone      ObjectFit = iota + 1 // object-none
	ObjectContain                        // object-contain
	ObjectCover                          // object-cover
	ObjectFill                           // object-fill
	ObjectScaleDown                      // object-scale-down
)

type ObjectFit int
