// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	MIT

package components

const (
	PositionStatic Position = iota + 1
	PositionFixed
	PositionAbsolute
	PositionRelative
	PositionSticky
)

// Position represents a position class.
//
// https://tailwindcss.com/docs/position
type Position int

// modifierPosition can be inherited by Components that support positioning
// within the document.
type modifierPosition struct {
	position Position
}

func (attr *modifierPosition) SetPosition(value Position) {
	attr.position = value
}

func (attr *modifierPosition) GetPosition() Position {
	return attr.position
}

func (attr *modifierPosition) PositionIsSet() bool {
	return attr.position > 0
}

type hasPosition[T any] interface {
	SetPosition(value Position)
	GetPosition() Position
}

// WithPositioning allows setting a position value for the Component to control
// its placement within the document.
//
// https://tailwindcss.com/docs/position
func WithPositioning[T hasPosition[T]](value Position) Option[T] {
	return func(c T) T {
		c.SetPosition(value)
		return c
	}
}
