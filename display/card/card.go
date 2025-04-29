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
	"github.com/joshuar/go-templ-daisyui/display/image"
	"github.com/joshuar/go-templ-daisyui/modifiers/color"
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
	*attributes.Attributes
	classes         *components.Classes
	body            *BodyProps
	baseColor       color.BaseColor
	image           *image.Props
	centeredContent bool
}

// BodyProps represents a Card's body properties. These are used when the Show()
// method is called to render the Card body.
type BodyProps struct {
	title   templ.Component
	content templ.Component
	*ActionsProps
	*attributes.Attributes
	classes *components.Classes
}

// ActionsProps is the container for Card actions.
type ActionsProps struct {
	actions []templ.Component
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

// WithBaseColor option sets a base color on the card.
func WithBaseColor(baseColor color.BaseColor) Option {
	return func(p *Props) {
		p.baseColor = baseColor
	}
}

// WithLayout option sets a card layout option.
func WithLayout(layout Layout) Option {
	return func(p *Props) {
		p.classes.Add(layout)
	}
}

// WithCenteredContent option aligns the content in the Card as centered.
func WithCenteredContent() Option {
	return func(p *Props) {
		p.centeredContent = true
	}
}

// WithImage option sets the image to be displayed in the Card.
func WithImage(url string, options ...image.Option) Option {
	return func(p *Props) {
		if url == "" {
			return
		}
		p.image = image.Build(url, options...)
	}
}

// WithSize option sets the size class of the card component.
func WithSize(cardSize Size) Option {
	return func(p *Props) {
		p.classes.Add(cardSize)
	}
}

// WithBodyOptions option sets the options for the Card body.
func WithBodyOptions(options ...BodyOption) Option {
	return func(p *Props) {
		p.body = BuildBody(options...)
	}
}

// WithTitle option sets a title in the Card body. This can be text, another DaisyUI
// component or a custom templ.Component. Other values are ignored and will
// result in an empty title.
func WithTitle(title any) BodyOption {
	return func(p *BodyProps) {
		p.title = components.SetContent(title)
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

func WithExtraClasses(extraClasses ...components.Class) Option {
	return func(p *Props) {
		p.classes.Add(extraClasses...)
	}
}

// WithContent option sets the Card body content. This can be text, another DaisyUI
// component or a custom templ.Component. Other values are ignored and will
// result in an empty title.
func WithContent(content any) BodyOption {
	return func(p *BodyProps) {
		p.content = components.SetContent(content)
	}
}

// WithActions option sets the actions to display in the bottom right of the
// Card body.
func WithActions(actions ...templ.Component) BodyOption {
	return func(p *BodyProps) {
		p.actions = actions
	}
}

// WithBodyID option sets the id attribute on the Card body.
func WithBodyID(id string) BodyOption {
	return func(p *BodyProps) {
		p.SetID(id)
	}
}

// WithBodyExtraAttributes option sets additional attributes on the Card body.
func WithBodyExtraAttributes(attrs templ.Attributes) BodyOption {
	return func(p *BodyProps) {
		p.AddAttributes(attrs)
	}
}

func WithBodyExtraClasses(extraClasses ...components.Class) BodyOption {
	return func(p *BodyProps) {
		p.classes.Add(extraClasses...)
	}
}

// Build creates a Card component from the given options.
func Build(options ...Option) *Props {
	card := &Props{
		Attributes: attributes.New(),
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
		ActionsProps: &ActionsProps{},
		Attributes:   attributes.New(),
		classes:      components.NewClasses(),
	}

	for _, option := range options {
		option(body)
	}

	return body
}
