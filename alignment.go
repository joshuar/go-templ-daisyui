// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

//go:generate go run golang.org/x/tools/cmd/stringer -type=Alignment -linecomment -output alignment_generated.go
package components

const (
	AlignEnd Alignment = iota
	AlignTop
	AlignBottom
	AlignMiddle
	AlignLeft
	AlignRight
)

type Alignment int
