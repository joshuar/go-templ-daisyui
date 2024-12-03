// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package components

import (
	"strings"
)

// componentClasses can be inherited by a component to be able to customize the
// class of the component.
type componentClasses struct {
	classes []string
}

// AddClasses will add the given classes to the component.
func (c *componentClasses) AddClasses(classes ...string) {
	c.classes = append(c.classes, classes...)
}

// Class will produce a space-delimeted string with all the classes assigned to
// the component.
func (c *componentClasses) Class() string {
	return strings.Join(c.classes, " ")
}

// WithClasses adds the given classes to the component.
func WithClasses[T any](classes ...string) Option[T] {
	return func(c T) T {
		// Get a pointer to the underlying component.
		component := &c
		// Apply the given classes to the component.
		addClassesToComponent(component, classes...)
		// Return the modified component.
		return *component
	}
}

// customisableClasses is an interface to detect components whose classes
// can be customized.
type customisableClasses interface {
	AddClasses(classes ...string)
}

func addClassesToComponent(component any, classes ...string) {
	if component, ok := component.(customisableClasses); ok {
		component.AddClasses(classes...)
	}
}
