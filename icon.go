// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

package components

const (
	Solid   FAStyle = iota // fa-solid
	Regular                // fa-regular
)

// FAStyle is a valid FontAwesome Icon style.
type FAStyle int

type Icon struct {
	name  string
	label string
	style FAStyle
}

// WithStyle assigns the given FontAwesome style to the Icon.
func WithStyle(s FAStyle) Option[Icon] {
	return func(i Icon) Icon {
		i.style = s
		return i
	}
}

// WithLabel ensures the Icon has the given label.
func WithLabel(l string) Option[Icon] {
	return func(i Icon) Icon {
		i.label = l
		return i
	}
}

// NewIcon creates an Icon component with the given options.
func NewIcon(name string, options ...Option[Icon]) Icon {
	icon := Icon{
		name: name,
	}

	for _, option := range options {
		icon = option(icon)
	}

	return icon
}
