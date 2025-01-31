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

// modifierChecked can be inherited by Components that can have a checked state
// (i.e., toggles, checkboxes...)
type modifierChecked struct {
	binaryState
}

// Check toggles the `checked` modifier of the Component to true (checked).
func (m *modifierChecked) Check() {
	m.state = true
}

// Check toggles the `checked` modifier of the Component to false.
func (m *modifierChecked) UnCheck() {
	m.state = false
}

// IsChecked returns whether the Component is checked (true) or not (false).
func (m *modifierChecked) IsChecked() bool {
	return m.state
}

type hasModifierChecked[T any] interface {
	Check()
	UnCheck()
}

// AsChecked will apply the checked modifier to the Component.
func AsChecked[T hasModifierChecked[T]]() Option[T] {
	return func(c T) T {
		c.Check()
		return c
	}
}
