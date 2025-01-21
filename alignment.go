// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

package components

const (
	AlignEnd    Alignment = iota // end
	AlignTop                     // top
	AlignBottom                  // bottom
	AlignMiddle                  // middle
	AlignLeft                    // left
	AlignRight                   // right
)

// Alignment represents a component alignment modifier.
type Alignment int

// modifierAlignment can be inherited by a Component for setting alignment.
type modifierAlignment struct {
	alignment Alignment
}

func (m *modifierAlignment) SetAlignment(alignment Alignment) {
	m.alignment = alignment
}

type hasAlignmentModifier[T any] interface {
	SetAlignment(alignment Alignment)
}

// WithAlignment adds an alignment to the component.
func WithAlignment[T hasAlignmentModifier[T]](alignment Alignment) Option[T] {
	return func(c T) T {
		c.SetAlignment(alignment)
		return c
	}
}
