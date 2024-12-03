// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

package components

type Menu struct {
	Title string
	componentClasses
}

func (m *Menu) String() string {
	return "menu"
}

type MenuItem struct {
	Tooltip *Tooltip
}

// MenuTitle sets the menu title.
func MenuTitle(title string) Option[Menu] {
	return func(m Menu) Menu {
		m.Title = title
		return m
	}
}

// MenuLayout sets how the menu items will be layed out.
func MenuLayout(layout Layout) Option[Menu] {
	return func(m Menu) Menu {
		m.AddClasses(layout.LayoutObject("menu"))
		return m
	}
}

// NewMenu creates a new menu with the given options.
func NewMenu(options ...Option[Menu]) Menu {
	menu := Menu{}
	menu = WithClasses[Menu](menu.String())(menu)

	for _, option := range options {
		menu = option(menu)
	}

	return menu
}

// MenuItemTooltip will wrap the menu item with the given tooltip.
func MenuItemTooltip(tip Tooltip) Option[MenuItem] {
	return func(mi MenuItem) MenuItem {
		mi.Tooltip = &tip
		return mi
	}
}

// NewMenuItem creates a new menu item with the given options.
func NewMenuItem(options ...Option[MenuItem]) MenuItem {
	item := MenuItem{}

	for _, option := range options {
		item = option(item)
	}

	return item
}
