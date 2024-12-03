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

type componentAlignment interface {
	fmt.Stringer
}

// WithAlignment adds an alignment to the component.
func WithAlignment[T any](alignment Alignment) Option[T] {
	return func(c T) T {
		if component, ok := any(&c).(componentAlignment); ok {
			c = WithClasses[T](component.String() + "-" + alignment.String())(c)
		}

		return c
	}
}
