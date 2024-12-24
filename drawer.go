// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package components

import "github.com/a-h/templ"

type Drawer struct {
	componentID
	visibility  Size
	mainContent templ.Component
	sideContent templ.Component
	modifierZIndex
}

// WithDrawerVisibility sets the size at which the drawer will always be
// visible. Below this value, the drawer will be hidden by default.
func WithDrawerVisibility(size Size) Option[Drawer] {
	return func(d Drawer) Drawer {
		d.visibility = size
		return d
	}
}

// WithMainContent sets the given templ.Component as the main content of the
// drawer.
func WithMainContent(content templ.Component) Option[Drawer] {
	return func(d Drawer) Drawer {
		d.mainContent = content
		return d
	}
}

// WithSideContent sets the given templ.Component as the side content of the
// drawer.
func WithSideContent(content templ.Component) Option[Drawer] {
	return func(d Drawer) Drawer {
		d.sideContent = content
		return d
	}
}

// NewCard creates a new card component with the given options. The card can be
// rendered by calling the Show method.
func NewDrawer(id string, options ...Option[Drawer]) Drawer {
	drawer := Drawer{}

	drawer.id = id

	for _, option := range options {
		drawer = option(drawer)
	}

	return drawer
}