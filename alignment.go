// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

package components

import (
	"fmt"
	"log/slog"
)

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

type hasAlignmentModifier interface {
	SetAlignment(alignment Alignment)
}

// WithAlignment adds an alignment to the component.
func WithAlignment[T any](alignment Alignment) Option[T] {
	return func(c T) T {
		component := &c

		if settable, ok := any(component).(hasAlignmentModifier); ok {
			settable.SetAlignment(alignment)
		} else {
			slog.Warn("WithAlignment passed as option to component, but component does not support alignment modifier.",
				slog.String("component", fmt.Sprintf("%T", settable)))
		}

		return *component
	}
}
