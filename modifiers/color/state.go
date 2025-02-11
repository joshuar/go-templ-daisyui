// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package color

import "github.com/joshuar/go-templ-daisyui/modifiers/state"

// StateColor represents a state as a color.
type StateColor = state.State

// StateColorProps contains properties for setting a state color.
type StateColorProps struct {
	state   StateColor
	outline bool
}

func (p *StateColorProps) SetStateColor(stateColor StateColor) {
	p.state = stateColor
}

func (p *StateColorProps) SetOutline(value bool) {
	p.outline = value
}

// ValidStateColor returns true if the state color is a valid state value.
func (p *StateColorProps) ValidStateColor() bool {
	return p.state.Valid()
}

func (p *StateColorProps) GetStateColor() StateColor {
	return p.state
}

func (p *StateColorProps) OutlineState() bool {
	return p.outline
}

func NewStateColor(color StateColor, outline bool) *StateColorProps {
	return &StateColorProps{
		state:   color,
		outline: outline,
	}
}
