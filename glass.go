// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	MIT

package components

import (
	"fmt"
	"log/slog"
)

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

type customisableGlass interface {
	SetGlass(value bool)
}

// WithGlass sets that the component should have a glass background.
//
//nolint:varnamelen // the var is copied into another with an appropriate name.
func WithGlass[T any]() Option[T] {
	return func(c T) T {
		component := &c

		if settable, ok := any(component).(customisableGlass); ok {
			settable.SetGlass(true)
		} else {
			slog.Warn("WithGlass passed as option to component, but component does not support glass modifier.",
				slog.String("component", fmt.Sprintf("%T", &c)))
		}

		return *component
	}
}
