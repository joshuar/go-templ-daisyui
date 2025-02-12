// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

// Package items defines an inheritable struct that Components with children can
// use to store their children.
package items

import (
	"github.com/a-h/templ"
)

// Items represents children of a Component, such as list items or badges.
type Items struct {
	Items []templ.Component
}

// AddItem will append the given item to the list of existing items.
func (i *Items) AddItem(item templ.Component) {
	i.Items = append(i.Items, item)
}

// ReplaceItems will overwrite any existing items with the given list of items.
func (i *Items) ReplaceItems(items ...templ.Component) {
	i.Items = items
}

// GetItems returns the list of items.
func (i *Items) GetItems() []templ.Component {
	return i.Items
}

// New initializes and returns an Attributes struct that can be embedded in
// other components.
func New() *Items {
	return &Items{}
}
