// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	MIT

package components

// modifierGlass can be embedded into components to allow setting a glass
// background on the component. The component will need to handle rendering
// based on the glass value itself.
type modifierGlass struct {
	glass bool
}

func (m *modifierGlass) SetGlass(value bool) {
	m.glass = value
}

func (m *modifierGlass) UseGlass() bool {
	return m.glass
}

type customGlass[T any] interface {
	SetGlass(value bool)
}

// WithGlass sets that the component should have a glass background.
func WithGlass[T customGlass[T]]() Option[T] {
	return func(c T) T {
		c.SetGlass(true)
		return c
	}
}
