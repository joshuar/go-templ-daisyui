// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	MIT

package components

import (
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

type customID[T any] interface {
	SetID(id string)
}

// WithID allows setting an id attribute on a component.
func WithID[T customID[T]](id string) Option[T] {
	return func(c T) T {
		c.SetID(id)
		return c
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

type customValue[T any] interface {
	SetValue(value string)
	Value() string
}

// WithValue allows setting an value on a component.
func WithValue[T customValue[T]](value string) Option[T] {
	return func(c T) T {
		c.SetValue(value)
		return c
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

type customItems[T any] interface {
	SetItems(items ...templ.Component)
}

// WithValue allows setting an value on a component.
func WithItems[T customItems[T]](items ...templ.Component) Option[T] {
	return func(c T) T {
		c.SetItems(items...)
		return c
	}
}
