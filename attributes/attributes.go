// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package attributes

import (
	"maps"
	"sync"

	"github.com/a-h/templ"
)

// Attributes represents additional named and custom attributes on an element.
type Attributes struct {
	attributes templ.Attributes
	mu         sync.Mutex
}

type (
	// ID is the id of the element.
	ID string
	// Value is the value of the element.
	Value string
	// Name is the name of the element.
	Name string
)

func (id ID) String() string {
	return string(id)
}

func (value Value) String() string {
	return string(value)
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
func (a *Attributes) SetID(id ID) {
	a.SetAttribute("id", string(id))
}

// SetValue will set the value attribute for the Component.
func (a *Attributes) SetValue(value Value) {
	a.SetAttribute("value", string(value))
}

// SetName will set the name attribute for the Component.
func (a *Attributes) SetName(name Name) {
	a.SetAttribute("name", string(name))
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

// type Option = opts.Option[Attributes]

// func WithID(id string) Option { return opts.Opt[Attributes, string]("id", id) }
