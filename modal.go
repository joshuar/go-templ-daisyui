// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	MIT

package components

import "github.com/a-h/templ"

// ModalProps represents the properties for a DaisyUI Modal component.
type ModalProps struct {
	htmlAttrID
	closeButton   templ.Component
	closeOutside  bool
	openInitially bool
}

// WithModalCloseOutside ensures the when clicked outside the modal, it will
// close it.
func WithModalCloseOutside() Option[*ModalProps] {
	return func(modal *ModalProps) *ModalProps {
		modal.closeOutside = true
		return modal
	}
}

// WithModalCloseButton inserts a button to close the modal, aligned in the
// modal with the given alignment.
func WithModalCloseButton() Option[*ModalProps] {
	return func(modal *ModalProps) *ModalProps {
		modal.closeButton = Button(
			WithID[*ButtonProps]("close"),
			WithButtonShape(ButtonCircle, false),
			WithColor[*ButtonProps](ColorGhost, false),
			WithButtonContent(AsIconContent("fa-xmark")),
			WithAttributes[*ButtonProps](templ.Attributes{
				"_": "on click or keyup[key=='Esc'] take .modal-open from #" + modal.id + " wait 200ms then remove #" + modal.id,
			}),
		)

		return modal
	}
}

func OpenModalInitially() Option[*ModalProps] {
	return func(modal *ModalProps) *ModalProps {
		modal.openInitially = true
		return modal
	}
}

// FromModalProps can be passed as an option to Modal, together with an existing
// ModalProps object, to render a Modal. This is useful if the ModalProps was
// pre-built before being rendered.
func FromModalProps(props *ModalProps) Option[*ModalProps] {
	return func(_ *ModalProps) *ModalProps {
		return props
	}
}

// BuildModal creates a new ModalProps from the given options. This can be used
// to pre-build a modal that can be shown later.
func BuildModal(options ...Option[*ModalProps]) *ModalProps {
	modal := &ModalProps{}

	for _, option := range options {
		modal = option(modal)
	}

	return modal
}
