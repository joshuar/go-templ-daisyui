// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package components

import (
	"fmt"
	"log/slog"
)

const (
	MaskSquircle          MaskStyle = iota + 1 // mask-squircle
	MaskHeart                                  // mask-heart
	MaskHexagon                                // mask-hexagon
	MaskHexagonAlt                             // mask-hexagon-2
	MaskDecagon                                // mask-decagon
	MaskPentagon                               // mask-pentagon
	MaskDiamond                                // mask-diamond
	MaskSquare                                 // mask-square
	MaskCircle                                 // mask-circle
	MaskParallelogram                          // mask-parallelogram
	MaskParallelogramAlt1                      // mask-parallelogram-2
	MaskParallelogramAlt2                      // mask-parallelogram-3
	MaskParallelogramAlt3                      // mask-parallelogram-4
	MaskStar                                   // mask-star
	MaskStarAlt                                // mask-star-2
	MaskTriangle                               // mask-triangle
	MaskTriangleAlt1                           // mask-triangle-2
	MaskTriangleAlt2                           // mask-triangle-3
	MaskTriangleAlt3                           // mask-triangle-4
	MaskHalfLeft                               // mask-half-1
	MaskHalfRight                              // mask-half-2
)

type MaskStyle int

type modifierMask struct {
	mask MaskStyle
}

func (m *modifierMask) SetMask(mask MaskStyle) {
	m.mask = mask
}

// MaskIsSet will return true if a mask has been set on the component.
func (m *modifierMask) MaskIsSet() bool {
	return m.mask > 0
}

type customisableMask interface {
	SetMask(mask MaskStyle)
}

// WithMask will add the given mask to the component.
//
//nolint:varnamelen // the var is copied into another with an appropriate name.
func WithMask[T any](mask MaskStyle) Option[T] {
	return func(c T) T {
		component := &c

		if settable, ok := any(component).(customisableMask); ok {
			settable.SetMask(mask)
		} else {
			slog.Warn("WithMask passed as option to component, but component does not support mask modifier.",
				slog.String("component", fmt.Sprintf("%T", &c)))
		}

		return *component
	}
}
