// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package state

//go:generate go tool golang.org/x/tools/cmd/stringer -type=State -linecomment -output state.gen.go

const (
	MinInvalid State = iota
	Info             // info
	Success          // success
	Warning          // warning
	Error            // error
	MaxInvalid
)

// State represents a state such as alert, info, warning or error.
type State int

// Valid returns whether the state is a valid value.
func (state State) Valid() bool {
	return state > MinInvalid && state < MaxInvalid
}

// AsState attempts to match a string representing a state to the appropriate state value.
func AsState[T ~string](value T) State {
	switch string(value) {
	case Info.String():
		return Info
	case Success.String():
		return Success
	case Warning.String():
		return Warning
	case Error.String():
		return Error
	}
	return 0
}
