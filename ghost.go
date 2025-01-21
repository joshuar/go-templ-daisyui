// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package components

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

type customGhost[T any] interface {
	SetGhost(value bool)
}

// WithGhost allows setting the ghost modifier on an attribute.
func AsGhost[T customGhost[T]](value bool) Option[T] {
	return func(c T) T {
		c.SetGhost(value)
		return c
	}
}
