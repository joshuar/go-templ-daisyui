// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

package components

import (
	"fmt"
	"log/slog"
)

const (
	StateInfo State = iota + 1
	StateSuccess
	StateWarning
	StateError
)

type State int

// modifierState can be embedded into components to allow setting the state
// color of the component. The component will need to handle rendering with the
// appropriate color itself.
type modifierState struct {
	state        State
	stateOutline bool
}

func (m *modifierState) SetState(state State, outline bool) {
	m.state = state
	m.stateOutline = outline
}

type customisableState interface {
	SetState(state State, outline bool)
}

// WithState styles the component with the given color state and optionally outlined.
//
//nolint:varnamelen // the var is copied into another with an appropriate name.
func WithState[T any](state State, outline bool) Option[T] {
	return func(c T) T {
		component := &c

		if settable, ok := any(component).(customisableState); ok {
			settable.SetState(state, outline)
		} else {
			slog.Warn("WithState passed as option to component, but component does not support state modifier.",
				slog.String("component", fmt.Sprintf("%T", &c)))
		}

		return *component
	}
}
