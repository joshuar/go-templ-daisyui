// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

package components

import (
	"fmt"
	"log/slog"
)

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

type hasModifierStateColor interface {
	SetStateColor(state StateColor, outline bool)
}

// WithStateColor styles the component with the given color state and optionally outlined.
//
//nolint:varnamelen // the var is copied into another with an appropriate name.
func WithStateColor[T any](state StateColor, outline bool) Option[T] {
	return func(c T) T {
		component := &c

		if settable, ok := any(component).(hasModifierStateColor); ok {
			settable.SetStateColor(state, outline)
		} else {
			slog.Warn("WithState passed as option to component, but component does not support state modifier.",
				slog.String("component", fmt.Sprintf("%T", &c)))
		}

		return *component
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

type hasModifierDisabled interface {
	Disable()
}

// WithStateColor styles the component with the given color state and optionally outlined.
//
//nolint:varnamelen // the var is copied into another with an appropriate name.
func AsDisabled[T any]() Option[T] {
	return func(c T) T {
		component := &c

		if settable, ok := any(component).(hasModifierDisabled); ok {
			settable.Disable()
		} else {
			slog.Warn("AsDisabled passed as option to component, but component does not support disabled modifier.",
				slog.String("component", fmt.Sprintf("%T", &c)))
		}

		return *component
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

type hasModifierActive interface {
	Activate()
}

// WithStateColor styles the component with the given color state and optionally outlined.
//
//nolint:varnamelen // the var is copied into another with an appropriate name.
func AsActive[T any]() Option[T] {
	return func(c T) T {
		component := &c

		if settable, ok := any(component).(hasModifierActive); ok {
			settable.Activate()
		} else {
			slog.Warn("AsActive passed as option to component, but component does not support active modifier.",
				slog.String("component", fmt.Sprintf("%T", &c)))
		}

		return *component
	}
}
