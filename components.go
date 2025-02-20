// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	MIT

package components

import (
	"github.com/a-h/templ"
)

// Option represents a generic option that can be applied to a component.
type Option[T any] func(T)

// Component represents any DaisyUI component.
type Component interface {
	Show(classes ...string) templ.Component
}

// SetContent converts the content of any valid type into a templ.Component to
// be used within a Component as its content. Invalid types are silently ignored
// and the function will return nil.
func SetContent(content any) templ.Component {
	switch value := content.(type) {
	case string:
		return templ.Raw(value)
	case Component:
		return value.Show()
	case templ.Component:
		return value
	default:
		return nil
	}
}
