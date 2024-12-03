// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package components

import (
	"strings"

	"github.com/a-h/templ"
)

// Option represents a generic option that can be applied to a component.
type Option[T any] func(T) T

// componentAttributes can be inherited by a component to be able to add customised
// attributes to the component.
type componentAttributes struct {
	attributes templ.Attributes
}

func (c *componentAttributes) AddAttributes(attrs templ.Attributes) {
	c.attributes = attrs
}

func (c *componentAttributes) Attributes() templ.Attributes {
	return c.attributes
}

// componentClasses can be inherited by a component to be able to customise the
// class of the component.
type componentClasses struct {
	classes []string
}

func (c *componentClasses) AddClasses(classes ...string) {
	c.classes = append(c.classes, classes...)
}

func (c *componentClasses) Class() string {
	return strings.Join(c.classes, " ")
}

type customisableAttributes interface {
	AddAttributes(attrs templ.Attributes)
}

type customisableClasses interface {
	AddClasses(classes ...string)
}

func AddClassesToComponent(component any, classes ...string) {
	switch component := component.(type) {
	case customisableClasses:
		component.AddClasses(classes...)
	}
}

func AddAttributesToComponent(component any, attr templ.Attributes) {
	switch component := component.(type) {
	case customisableAttributes:
		component.AddAttributes(attr)
	}
}

// WithClasses adds the given classes to the component.
func WithClasses[T any](classes ...string) Option[T] {
	return func(c T) T {
		component := &c
		AddClassesToComponent(component, classes...)
		return *component
	}
}

// WithAttributes adds the given attributes to the component.
func WithAttributes[T any](attr templ.Attributes) Option[T] {
	return func(c T) T {
		component := &c
		AddAttributesToComponent(component, attr)
		return *component
	}
}
