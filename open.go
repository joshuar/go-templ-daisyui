// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package components

const (
	OpenLeft OpenFrom = iota
	OpenRight
	OpenTop
	OpenBottom
)

// OpenFrom represents the way a Component will open.
type OpenFrom int

// ModifierOpenFrom can be embedded in a component that can open/close. It
// defines the way the Component will open.
type ModifierOpenFrom struct {
	openFrom  OpenFrom
	forceOpen bool
}

// SetOpenFrom can be used on a Component to set the OpenFrom value.
func (m *ModifierOpenFrom) SetOpenFrom(from OpenFrom) {
	m.openFrom = from
}

// SetForceOpen can be used on a Component to force-open the Component.
func (m *ModifierOpenFrom) SetForceOpen(value bool) {
	m.forceOpen = value
}

// HasOpenModifier indicates that a Component has the OpenFrom Modifier.
type hasOpenFromModifier[T any] interface {
	SetOpenFrom(from OpenFrom)
	SetForceOpen(value bool)
}

// WithOpenFrom option specifies how the Component will show its
// content on opening.
func WithOpenFrom[T hasOpenFromModifier[T]](from OpenFrom) Option[T] {
	return func(c T) T {
		c.SetOpenFrom(from)
		return c
	}
}

// WithForceOpen option specifies that the component should be open by default.
func WithForceOpen[T hasOpenFromModifier[T]]() Option[T] {
	return func(c T) T {
		c.SetForceOpen(true)
		return c
	}
}
