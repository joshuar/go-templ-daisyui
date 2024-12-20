// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package components

import (
	"fmt"
	"log/slog"
)

// Option represents a generic option that can be applied to a component.
type Option[T any] func(T) T

// componentID is an inheritable struct for components that can have an id
// attribute.
type componentID struct {
	id string
}

func (c *componentID) SetID(id string) {
	c.id = id
}

func (c *componentID) ID() string {
	return c.id
}

type customisableID interface {
	SetID(id string)
}

// WithID allows setting an id attribute on a component.
//
//nolint:varnamelen // the var is copied into another with an appropriate name.
func WithID[T any](id string) Option[T] {
	return func(c T) T {
		component := &c

		if settable, ok := any(component).(customisableID); ok {
			settable.SetID(id)
		} else {
			slog.Warn("WithID passed as option to component, but component does not support setting ID.",
				slog.String("component", fmt.Sprintf("%T", &c)))
		}

		return *component
	}
}

// componentID is an inheritable struct for components that can have an id
// attribute.
type componentValue struct {
	value string
}

func (c *componentValue) SetValue(value string) {
	c.value = value
}

func (c *componentValue) Value() string {
	return c.value
}

type customisableValue interface {
	SetValue(value string)
	Value() string
}

// WithValue allows setting an value on a component.
//
//nolint:varnamelen // the var is copied into another with an appropriate name.
func WithValue[T any](value string) Option[T] {
	return func(c T) T {
		component := &c

		if settable, ok := any(component).(customisableValue); ok {
			settable.SetValue(value)
		} else {
			slog.Warn("WithValue passed as option to component, but component does not support setting value.",
				slog.String("component", fmt.Sprintf("%T", &c)))
		}

		return *component
	}
}
