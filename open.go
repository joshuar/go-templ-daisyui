// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package components

import (
	"fmt"
	"log/slog"
)

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
type HasOpenFromModifier interface {
	SetOpenFrom(from OpenFrom)
	SetForceOpen(value bool)
}

// WithOpenFrom option specifies how the Component will show its
// content on opening.
//
//nolint:varnamelen // the var is copied into another with an appropriate name.
func WithOpenFrom[T any](from OpenFrom) Option[T] {
	return func(c T) T {
		component := &c

		if settable, ok := any(component).(HasOpenFromModifier); ok {
			settable.SetOpenFrom(from)
		} else {
			slog.Warn("WithOpenFrom passed as option to component, but component does not support opening.",
				slog.String("component", fmt.Sprintf("%T", c)))
		}

		return *component
	}
}

// WithForceOpen option specifies that the component should be open by default.
//
//nolint:varnamelen // the var is copied into another with an appropriate name.
func WithForceOpen[T any]() Option[T] {
	return func(c T) T {
		component := &c

		if settable, ok := any(component).(HasOpenFromModifier); ok {
			settable.SetForceOpen(true)
		} else {
			slog.Warn("WithForceOpen passed as option to component, but component does not support opening.",
				slog.String("component", fmt.Sprintf("%T", &c)))
		}

		return *component
	}
}
