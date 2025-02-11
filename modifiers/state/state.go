// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package state

const (
	MinInvalid State = iota
	Info
	Success
	Warning
	Error
	MaxInvalid
)

// State represents a state such as alert, info, warning or error.
type State int

// Valid returns whether the state is a valid value.
func (state State) Valid() bool {
	return state > MinInvalid && state < MaxInvalid
}
