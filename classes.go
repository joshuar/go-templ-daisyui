// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package components

import (
	"slices"
	"strings"

	"github.com/a-h/templ"
)

// Class represents a class value of a component.
type Class string

func (c Class) String() string {
	return string(c)
}

// Show returns a templ.KV that will conditionally render the class if it is not an empty string.
func (c Class) Show() templ.KeyValue[string, bool] {
	return templ.KV(c.String(), c != "")
}

// Classes holds custom classes for a component.
type Classes struct {
	values []Class
}

// NewClasses initializes a Classes object for use in a Component.
func NewClasses(classes ...Class) *Classes {
	c := &Classes{}

	// Add any classes specified.
	if len(classes) > 0 {
		c.Add(classes...)
	}

	return &Classes{}
}

func (c *Classes) String() string {
	values := make([]string, 0, len(c.values))
	for class := range slices.Values(c.values) {
		values = append(values, class.String())
	}
	return strings.Join(values, " ")
}

// Show returns a templ.KV that will conditionally render the class list, if classes have been set.
func (c *Classes) Show() templ.KeyValue[string, bool] {
	return templ.KV(c.String(), c.HasClasses())
}

// HasClasses returns a boolean indicating whether any classes have been set.
func (c *Classes) HasClasses() bool {
	return len(c.values) > 0
}

// Add will add the given class to the list of classes.
func (c *Classes) Add(classes ...Class) {
	c.values = append(c.values, classes...)
}
