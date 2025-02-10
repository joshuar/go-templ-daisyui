// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package modal

import (
	"github.com/a-h/templ"

	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/actions/button"
	"github.com/joshuar/go-templ-daisyui/attributes"
	"github.com/joshuar/go-templ-daisyui/display/icon"
	"github.com/joshuar/go-templ-daisyui/modifiers/color"
)

type Option components.Option2[*Props]

// Props represents the properties for a DaisyUI Modal component.
type Props struct {
	*attributes.Attributes
	closeButton   templ.Component
	closeOutside  bool
	openInitially bool
}

// WithModalCloseOutside ensures the when clicked outside the modal, it will
// close it.
func WithModalCloseOutside() Option {
	return func(p *Props) {
		p.closeOutside = true
	}
}

// WithModalCloseButton inserts a button to close the modal, aligned in the
// modal with the given alignment.
func WithModalCloseButton() Option {
	return func(p *Props) {
		p.closeButton = button.Build(
			button.AsShape(button.Circle, false),
			button.WithThemeColor(color.Ghost, false),
			button.WithContent(icon.Build("fa-xmark")),
			button.WithExtraAttributes(templ.Attributes{
				"_": "on click trigger closeModal",
				// "_": "on click or keyup[key=='Esc'] take .modal-open from #" + modal.id + " wait 200ms then remove #" + modal.id,
			}),
			button.WithID("close"),
		).Show()
	}
}

func OpenModalInitially() Option {
	return func(p *Props) {
		p.openInitially = true
	}
}

func WithName(name attributes.Name) Option {
	return func(p *Props) {
		p.SetName(name)
	}
}

func WithID(id attributes.ID) Option {
	return func(p *Props) {
		p.SetID(id)
	}
}

func WithExtraAttributes(attrs templ.Attributes) Option {
	return func(p *Props) {
		p.AddAttributes(attrs)
	}
}

// Build creates a new ModalProps from the given options. This can be used
// to pre-build a modal that can be shown later.
func Build(options ...Option) *Props {
	modal := &Props{
		Attributes: attributes.New(),
	}

	for _, option := range options {
		option(modal)
	}

	return modal
}
