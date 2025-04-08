// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

// Package card is a DaisyUI Card component.
//
// https://daisyui.com/components/card/
package card

import (
	"github.com/a-h/templ"
	components "github.com/joshuar/go-templ-daisyui"
	"github.com/joshuar/go-templ-daisyui/attributes"
	"github.com/joshuar/go-templ-daisyui/classes"
	"github.com/joshuar/go-templ-daisyui/display/image"
	"github.com/joshuar/go-templ-daisyui/modifiers/color"
	"github.com/joshuar/go-templ-daisyui/modifiers/shadow"
	"github.com/joshuar/go-templ-daisyui/modifiers/size"
)

const (
	LayoutNormal  Layout = iota + 1 // card-normal
	LayoutCompact                   // card-compact
	LayoutSide                      // card-side
)

// Layout defines a Card layout.
type Layout int

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
	*classes.Classes
	body            *BodyProps
	border          bool
	glass           bool
	layout          Layout
	baseColor       color.BaseColor
	shadow          shadow.ShadowSize
	fullImage       bool
	image           *image.Props
	centeredContent bool
	size            size.ResponsiveSize
}

// BodyProps represents a Card's body properties. These are used when the Show()
// method is called to render the Card body.
type BodyProps struct {
	title   templ.Component
	content templ.Component
	*ActionsProps
	*attributes.Attributes
	*classes.Classes
}

// ActionsProps is the container for Card actions.
type ActionsProps struct {
	actions []templ.Component
}

// Bordered option will ensure the Card has a border.
func Bordered() Option {
	return func(p *Props) {
		p.border = true
	}
}

// Glass option will ensure the Card is displayed with a glass background.
func Glass() Option {
	return func(p *Props) {
		p.glass = true
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
		p.layout = layout
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
		p.image = image.Build(url, options...)
	}
}

// WithImageBackground option will set the image as the background of the Card.
// Using this option implies usage of WithImage option to specify the image,
// else it will have no effect.
func WithFullImage() Option {
	return func(p *Props) {
		p.fullImage = true
	}
}

// WithShadow option will apply a shadow of the given size to the Card.
func WithShadow(shadowSize shadow.ShadowSize) Option {
	return func(p *Props) {
		p.shadow = shadowSize
	}
}

// WithText will set the text to display within the Badge.
func WithSize(cardSize size.ResponsiveSize) Option {
	return func(p *Props) {
		p.size = cardSize
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

func WithExtraClasses(extraClasses ...classes.Class) Option {
	return func(p *Props) {
		for _, class := range extraClasses {
			p.AddClass(class)
		}
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

func WithBodyExtraClasses(extraClasses ...classes.Class) BodyOption {
	return func(p *BodyProps) {
		for _, class := range extraClasses {
			p.AddClass(class)
		}
	}
}

// Build creates a Card component from the given options.
func Build(options ...Option) *Props {
	card := &Props{
		Attributes: attributes.New(),
		Classes:    classes.New(),
	}

	for _, option := range options {
		option(card)
	}

	return card
}

// BuildBody creates a Card body from the given options.
func BuildBody(options ...BodyOption) *BodyProps {
	body := &BodyProps{
		ActionsProps: &ActionsProps{},
		Attributes:   attributes.New(),
		Classes:      classes.New(),
	}

	for _, option := range options {
		option(body)
	}

	return body
}
