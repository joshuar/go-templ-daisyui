// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

package components

import (
	"fmt"
	"log/slog"
)

// Size options.
const (
	XS Size = iota + 1 // xs
	SM                 // sm
	MD                 // md
	LG                 // lg
	XL                 // xl
)

type Size int

// modifierSize can be embedded into components to allow setting a size option
// on the component. The component will need to handle rendering with the
// appropriate size itself.
type modifierSize struct {
	size Size
}

func (m *modifierSize) SetSize(size Size) {
	m.size = size
}

// customisableSize is an inheritable struct for components that can have a size
// property.
type customisableSize interface {
	SetSize(size Size)
}

// WithSize adds the given size to the component.
func WithSize[T any](size Size) Option[T] {
	return func(c T) T {
		component := &c

		if settable, ok := any(component).(customisableSize); ok {
			settable.SetSize(size)
		} else {
			slog.Warn("WithSize passed as option to component, but component does not support size modifier.",
				slog.String("component", fmt.Sprintf("%T", &c)))
		}

		return *component
	}
}
