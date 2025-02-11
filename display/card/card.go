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
	"github.com/joshuar/go-templ-daisyui/display/image"
	"github.com/joshuar/go-templ-daisyui/modifiers/color"
	"github.com/joshuar/go-templ-daisyui/modifiers/shadow"
)

const (
	LayoutNormal  Layout = iota + 1 // card-normal
	LayoutCompact                   // card-compact
	LayoutSide                      // card-side
)

type Layout int

type Option components.Option2[*Props]

type Props struct {
	*attributes.Attributes
	body            *BodyProps
	border          bool
	glass           bool
	layout          Layout
	baseColor       color.BaseColor
	shadow          shadow.ShadowSize
	fullImage       bool
	image           *image.Props
	centeredContent bool
}

// BodyProps is the container for the Card content.
type BodyProps struct {
	title   templ.Component
	content templ.Component
	*ActionsProps
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

// WithTitle option sets the Card title. This can be text, another DaisyUI
// component or a custom templ.Component. Other values are ignore and will
// result in an empty title.
func WithTitle(title any) Option {
	return func(p *Props) {
		p.body.title = components.SetContent(title)
	}
}

// WithContent option sets the Card body content.
func WithContent(content templ.Component) Option {
	return func(p *Props) {
		p.body.content = content
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

// WithActions option sets the actions to display in the bottom right of the
// Card body.
func WithActions(actions ...templ.Component) Option {
	return func(p *Props) {
		p.body.actions = actions
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

func Build(options ...Option) *Props {
	card := &Props{
		Attributes: attributes.New(),
		body: &BodyProps{
			ActionsProps: &ActionsProps{},
		},
	}

	for _, option := range options {
		option(card)
	}

	return card
}
