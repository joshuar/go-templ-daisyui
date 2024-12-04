// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

//go:generate go run golang.org/x/tools/cmd/stringer -type=Alignment -linecomment -output alignment_generated.go
package components

import "fmt"

const (
	AlignEnd    Alignment = iota // end
	AlignTop                     // top
	AlignBottom                  // bottom
	AlignMiddle                  // middle
	AlignLeft                    // left
	AlignRight                   // right
)

type Alignment int

type customisableAlignment interface {
	fmt.Stringer
}

// WithAlignment adds an alignment to the component.
func WithAlignment[T any](alignment Alignment) Option[T] {
	return func(c T) T {
		if settable, ok := any(&c).(customisableAlignment); ok {
			c = WithClasses[T](settable.String() + "-" + alignment.String())(c)
		}

		return c
	}
}
