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

type Option components.Option[*Props]

// Props represents the properties for a DaisyUI Modal component.
type Props struct {
	*attributes.Attributes
	closeButton   bool
	closeOutside  bool
	openInitially bool
	actions       []*button.Props
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

// OpenInitially option sets the modal to be open when first shown.
func OpenInitially() Option {
	return func(p *Props) {
		p.openInitially = true
	}
}

// WithActions option adds the given buttons as actions in the Modal (displayed in the bottom right below the content).
func WithActions(buttons ...*button.Props) Option {
	return func(p *Props) {
		p.actions = buttons
	}
}

func WithName(name attributes.Name) Option {
	return func(p *Props) {
		p.SetName(name)
	}
}

func WithID(id string) Option {
	return func(p *Props) {
		p.SetID(id)
	}
}

func WithExtraAttributes(attrs templ.Attributes) Option {
	return func(p *Props) {
		p.AddAttributes(attrs)
	}
}

// Build generates Modal properties from the given options.
func Build(options ...Option) *Props {
	modal := &Props{
		Attributes: attributes.New(),
	}

	for _, option := range options {
		option(modal)
	}

	return modal
}
