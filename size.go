// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

//go:generate go run golang.org/x/tools/cmd/stringer -type=Size -linecomment -output size_generated.go
package components

// Size options. MD is default in DaisyUI so it goes first.
const (
	MD Size = iota // md
	XS             // xs
	SM             // sm
	LG             // lg
	XL             // xl
)

type Size int

// WithSize adds the given size to the component.
func WithSize[T any](size Size) Option[T] {
	return func(c T) T {
		// Get a pointer to the underlying component.
		component := &c
		// Apply the given classes to the component.
		addSizeToComponent(component, size)
		// Return the modified component.
		return *component
	}
}

// customisableSize is an interface to detect components whose classes
// can be sized.
type customisableSize interface {
	String() string
	customisableClasses
}

func addSizeToComponent(component any, size Size) {
	switch component := component.(type) {
	case customisableSize:
		component.AddClasses(component.String() + "-" + size.String())
	}
}
