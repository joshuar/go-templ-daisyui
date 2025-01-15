// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package components

import (
	"fmt"
	"log/slog"
)

// modifierAutofocus can be inherited by Components that can use the "autofocus"
// attribute.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/autofocus
type modifierAutofocus struct {
	binaryState
}

func (m *modifierAutofocus) SetAutofocus() {
	m.state = true
}

// ObjectFitIsSet will return true if the Component has an object-fit value.
func (m *modifierAutofocus) HasAutofocus() bool {
	return m.state
}

type hasAutofocusModifier interface {
	SetAutofocus()
}

// WithObjectFit will add an object fit to the component class.
//
//nolint:varnamelen // the var is copied into another with an appropriate name.
func WithAutofocus[T any]() Option[T] {
	return func(c T) T {
		component := &c

		if settable, ok := any(component).(hasAutofocusModifier); ok {
			settable.SetAutofocus()
		} else {
			slog.Warn("WithAutofocus passed as option to component, but component does not support autofocus modifier.",
				slog.String("component", fmt.Sprintf("%T", &c)))
		}

		return *component
	}
}
