// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

// Package classes can be inherited by components to provide the ability to set
// custom classes on the component.
package classes

import (
	"strings"

	"github.com/a-h/templ"
)

type Class = string

// Classes holds custom classes for a component.
type Classes struct {
	classes []Class
}

// HasClasses returns a boolean indicating whether any classes have been set.
func (c *Classes) HasClasses() bool {
	return len(c.classes) > 0
}

// ShowClasses returns a space-separated string with all classes.
func (c *Classes) ShowClasses() templ.KeyValue[string, bool] {
	return templ.KV(strings.Join(c.classes, " "), c.HasClasses())
}

// AddClass will add the given class to the list of classes.
func (c *Classes) AddClass(class Class) {
	c.classes = append(c.classes, class)
}

// New initializes a Classes object for use in a Component.
func New() *Classes {
	return &Classes{}
}
