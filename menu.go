// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

package components

import "github.com/a-h/templ"

type Menu struct {
	title string
	componentClasses
	modifierSize
	layout Layout
	items  []templ.Component
}

// WithMenuTitle sets the menu title.
func WithMenuTitle(title string) Option[Menu] {
	return func(m Menu) Menu {
		m.title = title
		return m
	}
}

// WithMenuLayout sets how the menu items will be laid out (either horizontal
// or vertical).
func WithMenuLayout(layout Layout) Option[Menu] {
	return func(m Menu) Menu {
		m.layout = layout
		return m
	}
}

// WithMenuItems adds the given components as menu items.
func WithMenuItems(items ...templ.Component) Option[Menu] {
	return func(m Menu) Menu {
		m.items = append(m.items, items...)
		return m
	}
}

// NewMenu creates a new menu with the given options.
func NewMenu(options ...Option[Menu]) Menu {
	menu := Menu{}

	for _, option := range options {
		menu = option(menu)
	}

	return menu
}
