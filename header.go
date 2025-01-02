// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	MIT

package components

const (
	H1 HeaderSize = iota // h1
	H2                   // h2
	H3                   // h3
	H4                   // h4
	H5                   // h5
	H6                   // h6
)

// HeaderSize defines the header size (H1...H6).
type HeaderSize int

// Header represents a standard HTML header (e.g., <h1>Header</h1>).
type Header struct {
	title string
	size  HeaderSize
}

// WithHeaderSize sets the header size. If not added as an option, the header
// defaults to h1.
func WithHeaderSize(size HeaderSize) Option[Header] {
	return func(h Header) Header {
		h.size = size
		return h
	}
}

// NewHeader creates a new header component with the given title and size.
func NewHeader(title string, options ...Option[Header]) Header {
	header := Header{
		title: title,
	}

	for _, option := range options {
		header = option(header)
	}

	return header
}
