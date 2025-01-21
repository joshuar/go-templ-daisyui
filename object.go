// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package components

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

// ObjectFitIsSet will return true if the Component has an object-fit value.
func (m *modifierObjectFit) ObjectFitIsSet() bool {
	return m.fit > 0
}

type customObjectFit[T any] interface {
	SetFit(fit ObjectFit)
}

// WithObjectFit will add an object fit to the component class.
func WithObjectFit[T customObjectFit[T]](fit ObjectFit) Option[T] {
	return func(c T) T {
		c.SetFit(fit)
		return c
	}
}
