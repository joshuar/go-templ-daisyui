// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

package components

import (
	"fmt"
	"log/slog"
)

type BaseInputProps struct {
	modifierResponsiveSize
	componentAttributes
	componentID
	componentValue
}

type modifierBordered struct {
	modifierColor
	modifierStateColor
	bordered bool
}

func (m *modifierBordered) SetBordered(value bool) {
	m.bordered = value
}

// Bordered will return true if the Component should have a bordered class.
func (m *modifierBordered) Bordered() bool {
	return m.bordered
}

type customBordered interface {
	SetBordered(value bool)
}

// Bordered will set whether the input should have a border.
func Bordered[T any](value bool) Option[T] {
	return func(c T) T {
		component := &c

		if settable, ok := any(component).(customBordered); ok {
			settable.SetBordered(value)
		}

		slog.Warn("Bordered passed as option to component, but component does not support bordered modifier.",
			slog.String("component", fmt.Sprintf("%T", &c)))

		return *component
	}
}

type modifierPlaceholder struct {
	placeholderText string
}

func (m *modifierPlaceholder) SetPlaceholder(text string) {
	m.placeholderText = text
}

// PlaceholderIsSet will return true if the Component has placeholder text.
func (m modifierPlaceholder) PlaceholderIsSet() bool {
	return m.placeholderText != ""
}

type customPlaceholder interface {
	SetPlaceholder(text string)
}

// WithPlaceholder will set the placeholder text.
func WithPlaceholder[T any](text string) Option[T] {
	return func(c T) T {
		component := &c

		if settable, ok := any(component).(customPlaceholder); ok {
			settable.SetPlaceholder(text)
		} else {
			slog.Warn("WithState passed as option to component, but component does not support state modifier.",
				slog.String("component", fmt.Sprintf("%T", &c)))
		}

		return *component
	}
}

type modifierTextValidation struct {
	required bool
	min      int
	max      int
	pattern  string
}

func (m modifierTextValidation) IsRequired() bool {
	return m.required
}

func (m *modifierTextValidation) SetValidation() {
	m.required = true
}

type customValidation interface {
	IsRequired() bool
	SetValidation()
}

// func (m *modifierPlaceholder) SetPlaceholder(text string) {
// 	m.placeholderText = text
// }

// // PlaceholderIsSet will return true if the Component has placeholder text.
// func (m modifierPlaceholder) PlaceholderIsSet() bool {
// 	return m.placeholderText != ""
// }

// WithPlaceholder will set the placeholder text.
func WithValidationRequired[T any]() Option[T] {
	return func(c T) T {
		component := &c

		if settable, ok := any(component).(customValidation); ok {
			settable.SetValidation()
		} else {
			slog.Warn("WithValidationRequired passed as option to component, but component does not support validation.",
				slog.String("component", fmt.Sprintf("%T", &c)))
		}

		return *component
	}
}
