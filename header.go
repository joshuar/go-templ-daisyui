// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

//go:generate go run golang.org/x/tools/cmd/stringer -type=HeaderSize -linecomment -output header_generated.go
package components

const (
	H1 HeaderSize = iota // h1
	H2                   // h2
	H3                   // h3
	H4                   // h4
	H5                   // h5
	H6                   // h6
)

type HeaderSize int

type Header struct {
	Title string
	Size  HeaderSize
}

// NewHeader creates a new header component with the given title and size.
func NewHeader(title string, size HeaderSize) Header {
	return Header{
		Title: title,
		Size:  size,
	}
}
