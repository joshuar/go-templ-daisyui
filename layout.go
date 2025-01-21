// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

package components

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

type customLayout[T any] interface {
	SetLayout(layout Layout)
}

// WithLayout adds the given layout to the Component.
func WithLayout[T customLayout[T]](layout Layout) Option[T] {
	return func(c T) T {
		c.SetLayout(layout)
		return c
	}
}
