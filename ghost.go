// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package components

import (
	"fmt"
	"log/slog"
)

// modifierGhost is an inheritable struct for Components that can have an ghost
// modifier.
type modifierGhost struct {
	ghost bool
}

// SetGhost will set the ghost modifier of the Component to the given value.
func (m *modifierGhost) SetGhost(value bool) {
	m.ghost = value
}

// GhostIsSet returns a boolean for whether this Component has the ghost modifier.
func (m *modifierGhost) GhostIsSet() bool {
	return m.ghost
}

type customGhost interface {
	SetGhost(value bool)
}

// WithGhost allows setting the ghost modifier on an attribute.
//
//nolint:varnamelen // the var is copied into another with an appropriate name.
func AsGhost[T any](value bool) Option[T] {
	return func(c T) T {
		component := &c

		if settable, ok := any(component).(customGhost); ok {
			settable.SetGhost(value)
		} else {
			slog.Warn("WithGhost passed as option to component, but component does not support setting ghost modifier.",
				slog.String("component", fmt.Sprintf("%T", &c)))
		}

		return *component
	}
}
