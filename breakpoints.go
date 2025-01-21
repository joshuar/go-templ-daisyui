// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	MIT

package components

// componentHiddenBreakpoint is an inheritable struct that allows a Component to
// be hidden on screens equal to and above the set size. Use
// componentRevealedBreakpoint to reveal a Component above certain sizes.
type componentHiddenBreakpoint struct {
	hiddenBreakpoint ResponsiveSize
}

func (c *componentHiddenBreakpoint) SetHiddenBreakpoint(size ResponsiveSize) {
	c.hiddenBreakpoint = size
}

// HiddenBreakpointIsSet will return true if the Component has a hidden breakpoint.
func (c *componentHiddenBreakpoint) HiddenBreakpointIsSet() bool {
	return c.hiddenBreakpoint > 0
}

type customHiddenBreakpoint[T any] interface {
	SetHiddenBreakpoint(size ResponsiveSize)
}

// WithHiddenBreakpoint allows hiding a Component on screens with the given size
// or above. On smaller screens, the Component will be revealed.
//
// https://tailwindcss.com/docs/display#hidden
func WithHiddenBreakpoint[T customHiddenBreakpoint[T]](size ResponsiveSize) Option[T] {
	return func(c T) T {
		c.SetHiddenBreakpoint(size)
		return c
	}
}

// componentRevealedBreakpoint is an inheritable struct that allows a Component to
// be revealed on screens equal to and above the set size. Use
// componentHiddenBreakpoint to hide a Component above certain sizes.
type componentRevealedBreakpoint struct {
	revealedBreakpoint ResponsiveSize
}

func (c *componentRevealedBreakpoint) SetRevealedBreakpoint(size ResponsiveSize) {
	c.revealedBreakpoint = size
}

// RevealedBreakpointIsSet will return true if the Component has a revealed breakpoint.
func (c *componentRevealedBreakpoint) RevealedBreakpointIsSet() bool {
	return c.revealedBreakpoint > 0
}

type customRevealedBreakpoint[T any] interface {
	SetRevealedBreakpoint(size ResponsiveSize)
}

// WithRevealedBreakpoint allows revealing a Component on screens with the given size
// or above. On smaller screens, the Component will be hidden.
//
// https://tailwindcss.com/docs/display#hidden
func WithRevealedBreakpoint[T customRevealedBreakpoint[T]](size ResponsiveSize) Option[T] {
	return func(c T) T {
		c.SetRevealedBreakpoint(size)
		return c
	}
}
