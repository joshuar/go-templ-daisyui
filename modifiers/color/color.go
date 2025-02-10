// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

// Color represents a DaisyUI color.
//
// https://daisyui.com/docs/colors/
package color

import "github.com/joshuar/go-templ-daisyui/modifiers/state"

const (
	Unset ThemeColor = iota
	Neutral
	Primary
	Secondary
	Accent
	Ghost
	// bg-base-100.
	Base100 BaseColor = iota
	// bg-base-200.
	Base200
	// bg-base-300.
	Base300
)

type (
	// ThemeColor represents a theme color.
	ThemeColor int
	// BaseColor represents a `bg-base-*` value.
	BaseColor int
	// StateColor represents a state as a color.
	StateColor state.State
)

// Colors can be embedded into components to allow setting the color of
// the component. The component will need to handle rendering with the
// appropriate color itself.
type Colors struct {
	theme *themeColorModifier
	state *stateModifier
	base  *baseModifier
}

type themeColorModifier struct {
	color   ThemeColor
	outline bool
}

// SetColor will set a theme color and optional outline.
func (m *Colors) SetColor(themeColor ThemeColor, outline bool) {
	m.theme = &themeColorModifier{
		color:   themeColor,
		outline: outline,
	}
}

// GetColor will get the theme color.
func (m *Colors) GetColor() ThemeColor {
	if m.theme != nil {
		return m.theme.color
	}

	return Unset
}

// ColorOutline returns whether the theme color should be outline style or not.
func (m *Colors) ColorOutline() bool {
	if m.theme == nil {
		return false
	}
	// Neutral and Ghost colors don't have an outline variant.
	if m.theme.outline && (m.theme.color != Neutral && m.theme.color != Ghost) {
		return true
	}

	return false
}

type stateModifier struct {
	state   StateColor
	outline bool
}

// SetState will set a state color and optional outline.
func (m *Colors) SetState(stateColor StateColor, outline bool) {
	m.state = &stateModifier{
		state:   stateColor,
		outline: outline,
	}
}

// GetState will get the state color.
func (m *Colors) GetState() state.State {
	if m.state != nil {
		return state.State(m.state.state)
	}

	return state.Unset
}

func (m *Colors) StateOutline() bool {
	if m.state != nil {
		return m.state.outline
	}

	return false
}

type baseModifier struct {
	base BaseColor
}

// SetBase will set a base color.
func (m *Colors) SetBase(baseColor BaseColor) {
	m.base = &baseModifier{
		base: baseColor,
	}
}

// GetBase returns the base color. If no base is set, it will return Base100
// color.
func (m *Colors) GetBase() BaseColor {
	if m.base != nil {
		return m.base.base
	}

	return Base100
}

// ColorIsSet will return true if the component has a color.
func (m *Colors) ColorIsSet() bool {
	if m.theme != nil {
		return m.theme.color > 0
	}

	return false
}

// ColorIsSet will return true if the component has a color.
func (m *Colors) StateIsSet() bool {
	if m.state != nil {
		return m.state.state > 0
	}

	return false
}

// ColorIsSet will return true if the component has a color.
func (m *Colors) BaseIsSet() bool {
	if m.base != nil {
		return m.base.base > 0
	}

	return false
}
