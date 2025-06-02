// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

// Package attributes can be inherited by components to provide the ability to
// set custom attributes on a component.
package attributes

import (
	"fmt"
	"maps"
	"sync"

	"github.com/a-h/templ"
)

// Attributes represents additional named and custom attributes on an element.
type Attributes struct {
	attributes templ.Attributes
	mu         sync.Mutex
}

// ID represents an id attribute in a HTML element.
type ID string

// Target returns the id attribute as a target (i.e., for htmx requests). This
// is the base id string with a "#" prefix.
func (a ID) Target() string {
	return fmt.Sprintf("#%s", a)
}

// String returns the id attribute as a string.
func (a ID) String() string {
	return string(a)
}

type (
	// Value is the value of the element.
	Value = string
	// Name is the name of the element.
	Name = string
)

// SetAttribute will set the attribute with the given key to the given value.
func (a *Attributes) GetAttribute(key string) string {
	a.mu.Lock()
	defer a.mu.Unlock()

	if attribute, found := a.attributes[key]; found {
		if value, ok := attribute.(string); ok {
			return value
		}
	}

	return ""
}

// SetAttribute will set the attribute with the given key to the given value.
func (a *Attributes) SetAttribute(key string, value any) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if _, found := a.attributes[key]; !found {
		a.attributes[key] = value
	}
}

// UnsetAttribute will unset the attribute with the given key.
func (a *Attributes) UnsetAttribute(key string) {
	a.mu.Lock()
	defer a.mu.Unlock()

	delete(a.attributes, key)
}

// AddAttributes will ensure the component has the given attributes. Any
// existing attributes are merged with the given set of attributes.
func (a *Attributes) AddAttributes(attrs templ.Attributes) {
	// Ignore if no attributes are passed in.
	if attrs == nil {
		return
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	if a.attributes != nil {
		maps.Copy(a.attributes, attrs)
	} else {
		a.attributes = attrs
	}
}

// ShowAttributes will return any set attributes. This can be used directly in
// templ code to display attributes:
//
// {{ component.ShowAttributes()... }}
//
// If no attributes are set, nothing will be rendered.
func (a *Attributes) ShowAttributes() templ.Attributes {
	return a.attributes
}

// SetID will set the id attribute for the Component.
func (a *Attributes) SetID(id string) {
	a.SetAttribute("id", id)
}

func (a *Attributes) GetID() string {
	return a.GetAttribute("id")
}

func (a *Attributes) GetIDTarget() string {
	id := ID(a.GetAttribute("id"))
	return id.Target()
}

// SetValue will set the value attribute for the Component. Passing an empty
// string is ignored and *will not* reset/wipe an existing value. To remove an
// existing value, call the UnsetValue method.
func (a *Attributes) SetValue(value Value) {
	if value != "" {
		a.SetAttribute("value", value)
	}
}

// UnsetValue will reset/wipe the value attribute on the Component.
func (a *Attributes) UnsetValue() {
	a.UnsetAttribute("value")
}

// SetName will set the name attribute for the Component.
func (a *Attributes) SetName(name Name) {
	a.SetAttribute("name", name)
}

// Checked will set  checked="checked" for the Component.
func (a *Attributes) Checked() {
	a.SetAttribute("checked", "checked")
}

// UnChecked will remove the checked attribute for the Component.
func (a *Attributes) UnChecked() {
	a.UnsetAttribute("checked")
}

// New initializes and returns an Attributes struct that can be embedded in
// other components.
func New() *Attributes {
	return &Attributes{
		attributes: make(templ.Attributes),
	}
}
