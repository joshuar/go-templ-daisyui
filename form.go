// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	MIT

package components

import "github.com/a-h/templ"

type FormProps struct {
	componentAttributes
	componentID
	components []templ.Component
}

// WithFormComponents adds the given components to the form. They will be
// rendered in the order given.
func WithFormComponents(components ...templ.Component) Option[*FormProps] {
	return func(form *FormProps) *FormProps {
		form.components = append(form.components, components...)
		return form
	}
}

// BuildForm creates FormProps that can be used to render a form.
func BuildForm(options ...Option[*FormProps]) *FormProps {
	form := &FormProps{}

	for _, option := range options {
		form = option(form)
	}

	return form
}
