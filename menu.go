// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

package components

type Menu struct {
	title string
	componentClasses
	componentAttributes
	componentItems
	componentHiddenBreakpoint
	componentRevealedBreakpoint
	modifierResponsiveSize
	modifierLayout
	modifierBaseColor
}

// WithMenuTitle sets the menu title.
func WithMenuTitle(title string) Option[Menu] {
	return func(m Menu) Menu {
		m.title = title
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

type DropDownMenu struct {
	Menu
	modifierBaseColor
}

// NewDropDownMenu creates a new menu wrapped in a dropdown with the given options.
func NewDropDownMenu(options ...Option[DropDownMenu]) DropDownMenu {
	menu := DropDownMenu{}

	for _, option := range options {
		menu = option(menu)
	}

	return menu
}
