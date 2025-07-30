// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

// Package dropdown defines a DaisyUI Dropdown.
//
// https://daisyui.com/components/dropdown
package dropdown

import (
	"github.com/a-h/templ"
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/attributes"
)

// Props holds common properties for all types of Dropdown components.
type Props struct {
	classes     *components.Classes
	attributes  *attributes.Attributes
	btnClasses  *components.Classes
	btnContent  any
	menuClasses *components.Classes
}

// Option is a functional option to apply to a Dropdown.
type Option components.Option[*Props]

// WithForceOpen option will ensure the dropdown will be open by default.
func WithForceOpen() Option {
	return func(p *Props) {
		p.classes.Add(ForceOpen)
	}
}

// WithHoverOpen option will ensure the dropdown will open on hover.
func WithHoverOpen() Option {
	return func(p *Props) {
		p.classes.Add(OpenOnHover)
	}
}

// WithOpenFrom option sets from where the dropdown will open.
func WithOpenFrom(open Open) Option {
	return func(p *Props) {
		p.classes.Add(open)
	}
}

// WithOpenAlignment option sets how the dropdown will be aligned when opened.
func WithOpenAlignment(align Alignment) Option {
	return func(p *Props) {
		p.classes.Add(align)
	}
}

// WithButtonContent option sets the content inside the button. If not specified a default text will be used.
func WithButtonContent(content any) Option {
	return func(p *Props) {
		p.btnContent = content
	}
}

// WithButtonClasses option adds additional class values to the dropdown button.
func WithButtonClasses(extraClasses ...components.Class) Option {
	return func(p *Props) {
		p.btnClasses.Add(extraClasses...)
	}
}

// WithMenuClasses option adds additional class values to the dropdown menu.
func WithMenuClasses(extraClasses ...components.Class) Option {
	return func(p *Props) {
		p.menuClasses.Add(extraClasses...)
	}
}

// WithExtraAttributes adds additional attributes to the component.
func WithAttributes(attrs templ.Attributes) Option {
	return func(btn *Props) {
		btn.attributes.AddAttributes(attrs)
	}
}

// WithExtraClasses adds additional classesto the component.
func WithClasses(extraClasses ...components.Class) Option {
	return func(p *Props) {
		p.classes.Add(extraClasses...)
	}
}

// Build will generate and set the properties of a dropdown with the given options.
func Build(options ...Option) *Props {
	dropdown := &Props{
		classes:     components.NewClasses(),
		btnClasses:  components.NewClasses(),
		menuClasses: components.NewClasses(),
		attributes:  attributes.New(),
	}

	for _, option := range options {
		option(dropdown)
	}

	return dropdown
}
