// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package input

const (
	TopLeft     FormControlLabelPosition = iota // Top-left form control label.
	TopRight                                    // Top-right form control label.
	BottomLeft                                  // Bottom-left form control label.
	BottomRight                                 // Bottom-right form control label.
)

// FormControlLabelPosition defines the position of a form control label.
type FormControlLabelPosition int

type FormControlProps struct {
	topLeft     string
	topRight    string
	bottomLeft  string
	bottomRight string
	enabled     bool
}

// SetFormControlLabel sets a label for the form control at the given position.
func (fc *FormControlProps) SetFormControlLabel(position FormControlLabelPosition, label string) {
	switch position {
	case TopLeft:
		fc.topLeft = label
	case TopRight:
		fc.topRight = label
	case BottomLeft:
		fc.bottomLeft = label
	case BottomRight:
		fc.bottomRight = label
	}
}

// WithFormControlLabel Option sets a form control label at the given position
// on the form control container.
func WithFormControlLabel(position FormControlLabelPosition, label string) Option {
	return func(p *Props) {
		if !p.FormControl.enabled {
			p.FormControl.enabled = true
		}

		p.FormControl.SetFormControlLabel(position, label)
	}
}

// AsFormControl will ensure the input is wrapped in a form control container.
func AsFormControl() Option {
	return func(p *Props) {
		p.FormControl.enabled = true
	}
}
