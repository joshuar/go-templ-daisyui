// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package components

import (
	"fmt"
	"log/slog"
)

const (
	ObjectNone      ObjectFit = iota + 1 // object-none
	ObjectContain                        // object-contain
	ObjectCover                          // object-cover
	ObjectFill                           // object-fill
	ObjectScaleDown                      // object-scale-down
)

type ObjectFit int

type modifierObjectFit struct {
	fit ObjectFit
}

func (m *modifierObjectFit) SetFit(fit ObjectFit) {
	m.fit = fit
}

type customisableObjectFit interface {
	SetFit(fit ObjectFit)
}

// WithObjectFit will add an object fit to the component class.
//
//nolint:varnamelen // the var is copied into another with an appropriate name.
func WithObjectFit[T any](fit ObjectFit) Option[T] {
	return func(c T) T {
		component := &c

		if settable, ok := any(component).(customisableObjectFit); ok {
			settable.SetFit(fit)
		} else {
			slog.Warn("WithObjectFit passed as option to component, but component does not support object fit modifier.",
				slog.String("component", fmt.Sprintf("%T", &c)))
		}

		return *component
	}
}
