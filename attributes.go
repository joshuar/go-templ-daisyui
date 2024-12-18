// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package components

import (
	"sync"

	"github.com/a-h/templ"
)

// componentAttributes can be inherited by a component to be able to add customized
// attributes to the component.
type componentAttributes struct {
	attributes templ.Attributes
	mu         sync.Mutex
}

func (c *componentAttributes) SetAttribute(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, found := c.attributes[key]; found {
		c.attributes[key] = value
	}
}

// AddAttributes will ensure the component has the given attributes.
func (c *componentAttributes) AddAttributes(attrs templ.Attributes) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.attributes = attrs
}

// Attributes can be used when rendering a component to render its attributes.
func (c *componentAttributes) Attributes() templ.Attributes {
	return c.attributes
}

// customisableAttributes is an interface to detect components whose attributes
// can be customized.
type customisableAttributes interface {
	AddAttributes(attrs templ.Attributes)
}

// WithAttributes adds the given attributes to the component.
func WithAttributes[T any](attr templ.Attributes) Option[T] {
	return func(c T) T {
		// Get a pointer to the underlying component.
		component := &c
		// If the component supports customizing attributes, set the attributes
		// to the given value.
		if settable, ok := any(component).(customisableAttributes); ok {
			settable.AddAttributes(attr)
		}
		// Return the modified component.
		return *component
	}
}
