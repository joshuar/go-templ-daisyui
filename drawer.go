// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package components

import "github.com/a-h/templ"

type drawerSide struct {
	sideContent templ.Component
}

type drawerContent struct {
	mainContent templ.Component
}

type DrawerProps struct {
	HtmlAttrID
	visibility ResponsiveSize
	drawerSide
	drawerContent
}

// WithDrawerVisibility sets the size at which the drawer will always be
// visible. Below this value, the drawer will be hidden by default.
func WithDrawerVisibility(size ResponsiveSize) Option[*DrawerProps] {
	return func(d *DrawerProps) *DrawerProps {
		d.visibility = size
		return d
	}
}

// WithMainContent sets the given templ.Component as the main content of the
// drawer.
func WithMainContent(content templ.Component) Option[*DrawerProps] {
	return func(d *DrawerProps) *DrawerProps {
		d.mainContent = content
		return d
	}
}

// WithSideContent sets the given templ.Component as the side content of the
// drawer.
func WithSideContent(content templ.Component) Option[*DrawerProps] {
	return func(d *DrawerProps) *DrawerProps {
		d.sideContent = content
		return d
	}
}

// NewCard creates a new card component with the given options. The card can be
// rendered by calling the Show method.
func BuildDrawer(id string, options ...Option[*DrawerProps]) *DrawerProps {
	drawer := &DrawerProps{}

	drawer.id = id

	for _, option := range options {
		drawer = option(drawer)
	}

	return drawer
}
