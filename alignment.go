// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

//go:generate go run golang.org/x/tools/cmd/stringer -type=Alignment -linecomment -output alignment_generated.go
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

type Alignment int

type modifierAlignment struct {
	alignment Alignment
}

func (m *modifierAlignment) SetAlignment(alignment Alignment) {
	m.alignment = alignment
}

type customisableAlignment interface {
	SetAlignment(alignment Alignment)
}

// WithAlignment adds an alignment to the component.
func WithAlignment[T any](alignment Alignment) Option[T] {
	return func(c T) T {
		component := &c

		if settable, ok := any(component).(customisableAlignment); ok {
			settable.SetAlignment(alignment)
		} else {
			slog.Warn("WithAlignment passed as option to component, but component does not support alignment modifier.",
				slog.String("component", fmt.Sprintf("%T", settable)))
		}

		return *component
	}
}
