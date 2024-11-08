// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

package components

import (
	"strings"
)

type Menu struct {
	Title string
	class []string
}

func (m Menu) Class() string {
	return strings.Join(m.class, " ")
}

type MenuOption Option[Menu]

type MenuItem struct {
	Tooltip *Tooltip
}

type MenuItemOption Option[MenuItem]

// MenuTitle sets the menu title.
func MenuTitle(title string) MenuOption {
	return func(m Menu) Menu {
		m.Title = title
		return m
	}
}

// MenuLayout sets how the menu items will be layed out.
func MenuLayout(layout Layout) MenuOption {
	return func(m Menu) Menu {
		m.class = append(m.class, layout.LayoutObject("menu"))
		return m
	}
}

// MenuSize sets the sizing of the menu.
func MenuSize(size Size) MenuOption {
	return func(m Menu) Menu {
		m.class = append(m.class, size.SizeObject("menu"))
		return m
	}
}

// MenuClasses sets additional class attributes on the menu.
func MenuClasses(classes ...string) MenuOption {
	return func(m Menu) Menu {
		m.class = append(m.class, classes...)
		return m
	}
}

// NewMenu creates a new menu with the given options.
func NewMenu(options ...MenuOption) Menu {
	menu := Menu{
		class: []string{"menu"},
	}

	for _, option := range options {
		menu = option(menu)
	}

	return menu
}

// MenuItemTooltip will wrap the menu item with the given tooltip.
func MenuItemTooltip(tip Tooltip) MenuItemOption {
	return func(mi MenuItem) MenuItem {
		mi.Tooltip = &tip
		return mi
	}
}

// NewMenuItem creates a new menu item with the given options.
func NewMenuItem(options ...MenuItemOption) MenuItem {
	item := MenuItem{}

	for _, option := range options {
		item = option(item)
	}

	return item
}
