// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package state

const (
	Unset State = iota
	Info
	Success
	Warning
	Error
)

// State represents a state such as alert, info, warning or error.
type State int
