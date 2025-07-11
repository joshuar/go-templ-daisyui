// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

// Package card is a DaisyUI Card component.
//
// https://daisyui.com/components/card/
package card

import (
	"slices"

	"github.com/a-h/templ"
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/attributes"
	"github.com/joshuar/go-templ-daisyui/classes/shadow"
)

type (
	// Option is a functional option for a Card component.
	Option components.Option[*Props]
	// BodyOption is a functional option for a Card body.
	BodyOption components.Option[*BodyProps]
)

// Props represents a Card's properties. These are used when the Show() method
// is called to render the Card.
type Props struct {
	attributes *attributes.Attributes
	classes    *components.Classes
}

// BodyProps represents a Card's body properties. These are used when the Show()
// method is called to render the Card body.
type BodyProps struct {
	attributes *attributes.Attributes
	classes    *components.Classes
}

// WithStyle option adds the given style class to the card.
func WithStyle(style Style) Option {
	return func(p *Props) {
		p.classes.Add(style)
	}
}

// WithShadow option will apply a shadow of the given size to the Card.
func WithShadow(shadow shadow.Shadow) Option {
	return func(p *Props) {
		p.classes.Add(shadow)
	}
}

// WithLayout option sets a card layout option.
func WithLayout(layout Layout) Option {
	return func(p *Props) {
		p.classes.Add(layout)
	}
}

// WithSize option sets the size class of the card component.
func WithSize(cardSize Size) Option {
	return func(p *Props) {
		p.classes.Add(cardSize)
	}
}

func WithAttributes(attrs templ.Attributes) Option {
	return func(p *Props) {
		p.attributes.AddAttributes(attrs)
	}
}

func WithClasses(extraClasses ...components.Class) Option {
	return func(p *Props) {
		p.classes.Add(extraClasses...)
	}
}

// WithBodyExtraAttributes option sets additional attributes on the Card body.
func WithBodyAttributes(attrs templ.Attributes) BodyOption {
	return func(p *BodyProps) {
		p.attributes.AddAttributes(attrs)
	}
}

func WithBodyClasses(extraClasses ...components.Class) BodyOption {
	return func(p *BodyProps) {
		p.classes.Add(extraClasses...)
	}
}

// Build creates a Card component from the given options.
func Build(options ...Option) *Props {
	card := &Props{
		attributes: attributes.New(),
		classes:    components.NewClasses(),
	}

	for option := range slices.Values(options) {
		if option != nil {
			option(card)
		}
	}

	return card
}

// BuildBody creates a Card body from the given options.
func BuildBody(options ...BodyOption) *BodyProps {
	body := &BodyProps{
		attributes: attributes.New(),
		classes:    components.NewClasses(),
	}

	for _, option := range options {
		option(body)
	}

	return body
}
