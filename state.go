// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

package components

type binaryState struct {
	state bool
}

// Toggle can be used to toggle the state of the Component.
func (m *binaryState) Toggle() {
	m.state = !m.state
}

const (
	StateInfo StateColor = iota + 1
	StateSuccess
	StateWarning
	StateError
)

type StateColor int

// modifierStateColor can be embedded into components to allow setting the state
// color of the component. The component will need to handle rendering with the
// appropriate color itself.
type modifierStateColor struct {
	stateColor        StateColor
	stateColorOutline bool
}

func (m *modifierStateColor) SetStateColor(state StateColor, outline bool) {
	m.stateColor = state
	m.stateColorOutline = outline
}

// StateIsSet will return true if the Component has a size.
func (m *modifierStateColor) StateColorIsSet() bool {
	return m.stateColor > 0
}

type hasModifierStateColor[T any] interface {
	SetStateColor(state StateColor, outline bool)
}

// WithStateColor styles the component with the given color state and optionally outlined.
func WithStateColor[T hasModifierStateColor[T]](state StateColor, outline bool) Option[T] {
	return func(c T) T {
		c.SetStateColor(state, outline)
		return c
	}
}

type modifierDisabled struct {
	binaryState
}

// Disable  can be used to disable the Component.
func (m *modifierDisabled) Disable() {
	m.state = true
}

func (m *modifierDisabled) IsDisabled() bool {
	return m.state
}

type hasModifierDisabled[T any] interface {
	Disable()
}

// WithStateColor styles the component with the given color state and optionally outlined.
func AsDisabled[T hasModifierDisabled[T]]() Option[T] {
	return func(c T) T {
		c.Disable()
		return c
	}
}

type modifierActive struct {
	binaryState
}

// Activate can be used to activate the Component.
func (m *modifierActive) Activate() {
	m.state = true
}

func (m *modifierActive) IsActive() bool {
	return m.state
}

type hasModifierActive[T any] interface {
	Activate()
}

// WithStateColor styles the component with the given color state and optionally outlined.
func AsActive[T hasModifierActive[T]]() Option[T] {
	return func(c T) T {
		c.Activate()
		return c
	}
}
