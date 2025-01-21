// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

package components

const (
	ColorNeutral Color = iota + 1
	ColorPrimary
	ColorSecondary
	ColorAccent
	ColorGhost
)

const (
	ColorBgBase100 BaseColor = iota
	ColorBgBase200
	ColorBgBase300
)

// Color represents a DaisyUI color.
// https://daisyui.com/docs/colors/
type Color int

type BaseColor int

// modifierColor can be embedded into components to allow setting the color of
// the component. The component will need to handle rendering with the
// appropriate color itself.
type modifierColor struct {
	color        Color
	colorOutline bool
}

func (m *modifierColor) SetColor(color Color, outline bool) {
	m.color = color
	m.colorOutline = outline
}

// ColorIsSet will return true if the component has a color.
func (m *modifierColor) ColorIsSet() bool {
	return m.color > 0
}

type customColor[T any] interface {
	SetColor(color Color, outline bool)
}

// WithColor styles the component with the given Color and optionally outlined.
func WithColor[T customColor[T]](color Color, outline bool) Option[T] {
	return func(c T) T {
		c.SetColor(color, outline)
		return c
	}
}

// modifierBaseColor can be embedded into components to allow setting the base
// color (`bg-base-*`) of the component. If an explicit value is not set with
// `WithBaseColor`, `bg-base-100` will be used on the component.
type modifierBaseColor struct {
	baseColor BaseColor
}

func (m *modifierBaseColor) SetBaseColor(color BaseColor) {
	m.baseColor = color
}

type customBaseColor[T any] interface {
	SetBaseColor(color BaseColor)
}

// WithBaseColor styles the component with the given base color (`bg-base-*`).
func WithBaseColor[T customBaseColor[T]](color BaseColor) Option[T] {
	return func(c T) T {
		c.SetBaseColor(color)
		return c
	}
}
