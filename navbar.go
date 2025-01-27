// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	MIT

package components

import "github.com/a-h/templ"

type NavBarProps struct {
	modifierZIndex
	modifierPosition
	modifierBaseColor
	htmlAttrID
	start  templ.Component
	center templ.Component
	end    templ.Component
}

func NavBarStart(component templ.Component) Option[*NavBarProps] {
	return func(navbar *NavBarProps) *NavBarProps {
		navbar.start = component
		return navbar
	}
}

func NavBarCenter(component templ.Component) Option[*NavBarProps] {
	return func(navbar *NavBarProps) *NavBarProps {
		navbar.center = component
		return navbar
	}
}

func NavBarEnd(component templ.Component) Option[*NavBarProps] {
	return func(navbar *NavBarProps) *NavBarProps {
		navbar.end = component
		return navbar
	}
}

func BuildNavBar(options ...Option[*NavBarProps]) *NavBarProps {
	navbar := &NavBarProps{}

	for _, option := range options {
		navbar = option(navbar)
	}

	return navbar
}
