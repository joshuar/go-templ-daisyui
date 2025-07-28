// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

// Package modal represents a DaisyUI Modal component.
//
// https://daisyui.com/components/modal/
package modal

import (
	"github.com/a-h/templ"

	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/actions/button"
	"github.com/joshuar/go-templ-daisyui/attributes"
)

// Option is a functional option for the modal.
type Option components.Option[*Props]

// Props represents the properties for a DaisyUI Modal component.
type Props struct {
	attributes   *attributes.Attributes
	closeButton  bool
	closeOutside bool
	actions      []*button.Props
	classes      *components.Classes
}

// WithCloseOutside option ensures the when clicked outside the modal, it will
// close it.
func WithCloseOutside() Option {
	return func(p *Props) {
		p.closeOutside = true
	}
}

// WithCloseButton option inserts a button to close the modal in the top-right corner.
func WithCloseButton() Option {
	return func(p *Props) {
		p.closeButton = true
	}
}

// WithActions option adds the given buttons as actions in the Modal (displayed in the bottom right below the content).
func WithActions(buttons ...*button.Props) Option {
	return func(p *Props) {
		p.actions = buttons
	}
}

// WithAttributes option sets additional attributes on the modal.
func WithAttributes(attrs templ.Attributes) Option {
	return func(p *Props) {
		p.attributes.AddAttributes(attrs)
	}
}

// WithClasses option sets additional classes on the modal.
func WithClasses(classes ...components.Class) Option {
	return func(p *Props) {
		p.classes.Add(classes...)
	}
}

// Build generates Modal properties from the given options.
func Build(id string, options ...Option) *Props {
	modal := &Props{
		attributes: attributes.New(),
		classes:    components.NewClasses(),
	}

	modal.attributes.SetID(id)

	for _, option := range options {
		option(modal)
	}

	return modal
}
