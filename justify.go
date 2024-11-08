// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

//go:generate go run golang.org/x/tools/cmd/stringer -type=JustifyContent -linecomment -output justify_generated.go
package components

const (
	JustifyNormal  JustifyContent = iota //
	JustifyStart                         // justify-start
	JustifyEnd                           // justify-end
	JustifyCenter                        // justify-center
	JustifyBetween                       // justify-between
	JustifyAround                        // justify-around
	JustifyEvenly                        // justify-evenly
	JustifyStretch                       // justify-stretch
)

type JustifyContent int
