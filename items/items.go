// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

// Package items defines an inheritable struct that Components with children can
// use to store their children.
package items

import (
	"slices"

	"github.com/a-h/templ"
	components "github.com/joshuar/go-templ-daisyui"
)

// Items represents children of a Component, such as list items or badges.
type Items struct {
	Items   []templ.Component
	classes *components.Classes
}

// AppendItem will append the given item to the list of existing items.
func (i *Items) AppendItem(item templ.Component) {
	i.Items = append(i.Items, item)
}

// PrependItem will prepend the given item to the list of existing items.
func (i *Items) PrependItem(item templ.Component) {
	i.Items = slices.Insert(i.Items, 0, item)
}

// ReplaceItems will overwrite any existing items with the given list of items.
func (i *Items) ReplaceItems(items ...templ.Component) {
	i.Items = items
}

// GetItems returns the list of items.
func (i *Items) GetItems() []templ.Component {
	return i.Items
}

// SetClasses sets the classes specified on each item.
func (i *Items) SetClasses(classes ...components.Class) {
	i.classes.Add(classes...)
}

// New initializes and returns an Attributes struct that can be embedded in
// other components.
func New() *Items {
	return &Items{
		classes: components.NewClasses(),
	}
}
