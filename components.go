// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package components

import (
	"fmt"
	"log/slog"

	"github.com/a-h/templ"
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

// componentItems is an inheritable struct for components that have "children"
// items, that could be any other component.
type componentItems struct {
	items []templ.Component
}

func (c *componentItems) SetItems(items ...templ.Component) {
	c.items = append(c.items, items...)
}

type customisableItems interface {
	SetItems(items ...templ.Component)
}

// WithValue allows setting an value on a component.
//
//nolint:varnamelen // the var is copied into another with an appropriate name.
func WithItems[T any](items ...templ.Component) Option[T] {
	return func(c T) T {
		component := &c

		if settable, ok := any(component).(customisableItems); ok {
			settable.SetItems(items...)
		} else {
			slog.Warn("WithItems passed as option to component, but component does not support setting items.",
				slog.String("component", fmt.Sprintf("%T", &c)))
		}

		return *component
	}
}

// componentHiddenBreakpoint is an inheritable struct that allows a Component to
// be hidden on screens equal to and above the set size. Use
// componentRevealedBreakpoint to reveal a Component above certain sizes.
type componentHiddenBreakpoint struct {
	hiddenBreakpoint Size
}

func (c *componentHiddenBreakpoint) SetHiddenBreakpoint(size Size) {
	c.hiddenBreakpoint = size
}

// HiddenBreakpointIsSet will return true if the Component has a hidden breakpoint.
func (c *componentHiddenBreakpoint) HiddenBreakpointIsSet() bool {
	return c.hiddenBreakpoint > 0
}

type customisableHiddenBreakpoint interface {
	SetHiddenBreakpoint(size Size)
}

// WithHiddenBreakpoint allows hiding a Component on screens with the given size
// or above. On smaller screens, the Component will be revealed.
//
// https://tailwindcss.com/docs/display#hidden
//
//nolint:varnamelen // the var is copied into another with an appropriate name.
func WithHiddenBreakpoint[T any](size Size) Option[T] {
	return func(c T) T {
		component := &c

		if settable, ok := any(component).(customisableHiddenBreakpoint); ok {
			settable.SetHiddenBreakpoint(size)
		} else {
			slog.Warn("WithVisibility passed as option to component, but component does not support setting visibility.",
				slog.String("component", fmt.Sprintf("%T", &c)))
		}

		return *component
	}
}

// componentRevealedBreakpoint is an inheritable struct that allows a Component to
// be revealed on screens equal to and above the set size. Use
// componentHiddenBreakpoint to hide a Component above certain sizes.
type componentRevealedBreakpoint struct {
	revealedBreakpoint Size
}

func (c *componentRevealedBreakpoint) SetRevealedBreakpoint(size Size) {
	c.revealedBreakpoint = size
}

// RevealedBreakpointIsSet will return true if the Component has a revealed breakpoint.
func (c *componentRevealedBreakpoint) RevealedBreakpointIsSet() bool {
	return c.revealedBreakpoint > 0
}

type customisableRevealedBreakpoint interface {
	SetRevealedBreakpoint(size Size)
}

// WithRevealedBreakpoint allows revealing a Component on screens with the given size
// or above. On smaller screens, the Component will be hidden.
//
// https://tailwindcss.com/docs/display#hidden
//
//nolint:varnamelen // the var is copied into another with an appropriate name.
func WithRevealedBreakpoint[T any](size Size) Option[T] {
	return func(c T) T {
		component := &c

		if settable, ok := any(component).(customisableRevealedBreakpoint); ok {
			settable.SetRevealedBreakpoint(size)
		} else {
			slog.Warn("WithVisibility passed as option to component, but component does not support setting visibility.",
				slog.String("component", fmt.Sprintf("%T", &c)))
		}

		return *component
	}
}
