// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

//go:generate go run golang.org/x/tools/cmd/stringer -type=Size -linecomment -output size_generated.go
package components

import (
	"fmt"
)

// Size options. MD is default in DaisyUI so it goes first.
const (
	MD Size = iota // md
	XS             // xs
	SM             // sm
	LG             // lg
	XL             // xl
)

type Size int

// customisableSize is an inheritable struct for components that can have a size
// property.
type customisableSize interface {
	fmt.Stringer
}

// WithSize adds the given size to the component.
func WithSize[T any](size Size) Option[T] {
	return func(c T) T {
		if component, ok := any(&c).(customisableSize); ok {
			c = WithClasses[T](component.String() + "-" + size.String())(c)
		}

		return c
	}
}
