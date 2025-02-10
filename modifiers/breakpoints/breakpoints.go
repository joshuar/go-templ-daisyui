// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	MIT

// Package breakpoints contains objects and methods to allow components to set
// hidden and revealed breakpoints.
package breakpoints

import "github.com/joshuar/go-templ-daisyui/modifiers/size"

// Breakpoints is an inheritable struct that allows a Component to
// be hidden or revealed on screens equal to and above the set sizes.
type Breakpoints struct {
	hidden   size.ResponsiveSize
	revealed size.ResponsiveSize
}

// SetHidden sets the size breakpoint from which the component will be hidden.
func (m *Breakpoints) SetHiddenBreakpoint(from size.ResponsiveSize) {
	m.hidden = from
}

// HiddenBreakpointSet returns whether a hidden breakpoint has been set.
func (m *Breakpoints) HiddenBreakpointSet() bool {
	return m.hidden > 0
}

// SetRevealed sets the size breakpoint from which the component will be revealed.
func (m *Breakpoints) SetRevealedBreakpoint(from size.ResponsiveSize) {
	m.revealed = from
}

// RevealedBreakpointSet returns whether a revealed breakpoint has been set.
func (m *Breakpoints) RevealedBreakpointSet() bool {
	return m.revealed > 0
}
