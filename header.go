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
	componentClasses
	Size HeaderSize
}

// WithHeaderSize sets the header size. If not added as an option, the header
// defaults to h1.
func WithHeaderSize(size HeaderSize) Option[Header] {
	return func(h Header) Header {
		h.Size = size
		return h
	}
}

// NewHeader creates a new header component with the given title and size.
func NewHeader(title string, options ...Option[Header]) Header {
	header := Header{
		Title: title,
	}

	for _, option := range options {
		header = option(header)
	}

	return header
}
