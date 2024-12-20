// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package components

import (
	"fmt"
	"log/slog"
)

const (
	Z0    ZIndex = iota + 1 // z-index: 0;
	Z10                     // z-index: 10;
	Z20                     // z-index: 20;
	Z30                     // z-index: 30;
	Z40                     // z-index: 40;
	Z50                     // z-index: 50;
	ZAuto                   // z-index: auto;
)

type ZIndex int

// modifierZIndex can be embedded into components to allow setting a z-index option
// on the component. The component will need to handle rendering with the
// appropriate size itself.
type modifierZIndex struct {
	zindex ZIndex
}

func (m *modifierZIndex) SetZIndex(index ZIndex) {
	m.zindex = index
}

type customisableZIndex interface {
	SetZIndex(index ZIndex)
}

// WithZIndex adds the given z-index to the component.
//
//nolint:varnamelen // the var is copied into another with an appropriate name.
func WithZIndex[T any](index ZIndex) Option[T] {
	return func(c T) T {
		component := &c

		if settable, ok := any(component).(customisableZIndex); ok {
			settable.SetZIndex(index)
		} else {
			slog.Warn("WithZIndex passed as option to component, but component does not support z-index modifier.",
				slog.String("component", fmt.Sprintf("%T", &c)))
		}

		return *component
	}
}
