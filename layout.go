// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

package components

import (
	"fmt"
	"log/slog"
)

// Layout options.
const (
	VerticalLayout   Layout = iota + 1 // vertical
	HorizontalLayout                   // horizontal
)

type Layout int

// modifierLayout can be embedded into Components to allow setting a layout option
// on the Component. The Component will need to handle rendering with the
// appropriate size itself.
type modifierLayout struct {
	layout Layout
}

func (m *modifierLayout) SetLayout(layout Layout) {
	m.layout = layout
}

// LayoutIsSet will return true if the Component has a layout.
func (m *modifierLayout) LayoutIsSet() bool {
	return m.layout > 0
}

type customisableLayout interface {
	SetLayout(layout Layout)
}

// WithLayout adds the given layout to the Component.
//
//nolint:varnamelen // the var is copied into another with an appropriate name.
func WithLayout[T any](layout Layout) Option[T] {
	return func(c T) T {
		component := &c

		if settable, ok := any(component).(customisableLayout); ok {
			settable.SetLayout(layout)
		} else {
			slog.Warn("WithLayout passed as option to component, but component does not support layout modifier.",
				slog.String("component", fmt.Sprintf("%T", &c)))
		}

		return *component
	}
}
