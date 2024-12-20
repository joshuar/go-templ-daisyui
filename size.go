// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

package components

import (
	"fmt"
	"log/slog"
)

const (
	XS Size = iota + 1 // Extra-small size. class: {component}-xs
	SM                 // Small size. class: {component}-sm
	MD                 // Medium size. class: {component}-md
	LG                 // Large size. class: {component}-lg
	XL                 // Extra-large size. class: {component}-xl
)

type Size int

// modifierSize can be embedded into components to allow setting a size option
// on the component. The component will need to handle rendering with the
// appropriate size itself.
type modifierSize struct {
	size Size
}

func (m *modifierSize) SetSize(size Size) {
	m.size = size
}

type customisableSize interface {
	SetSize(size Size)
}

// WithSize adds the given size to the component.
//
//nolint:varnamelen // the var is copied into another with an appropriate name.
func WithSize[T any](size Size) Option[T] {
	return func(c T) T {
		component := &c

		if settable, ok := any(component).(customisableSize); ok {
			settable.SetSize(size)
		} else {
			slog.Warn("WithSize passed as option to component, but component does not support size modifier.",
				slog.String("component", fmt.Sprintf("%T", &c)))
		}

		return *component
	}
}
