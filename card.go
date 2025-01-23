// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package components

import (
	"github.com/a-h/templ"
)

const (
	CardLayoutNormal         CardLayout = iota // card-normal
	CardLayoutSide                             // card-side
	CardLayoutSideResponsive                   // lg:card-side
)

const (
	ImageTop ImageLayout = iota
	ImageBottom
	ImageOverlay
)

// CardLayout is used to define the card style.
type (
	CardLayout  int
	ImageLayout int
)

// cardTitleProps defines the properties for a Card title.
type cardTitleProps struct {
	text   string
	size   HeaderSize
	badges []templ.Component
}

type cardImageProps struct {
	image  templ.Component
	layout ImageLayout
}

type CardBodyProps struct {
	Content templ.Component
	borderContent
	componentAttributes
}

// WithTopRightActions defines the actions to position at the top-right of the
// Card body.
func WithTopRightActions(actions ...templ.Component) Option[*CardBodyProps] {
	return func(cardBody *CardBodyProps) *CardBodyProps {
		cardBody.SetBorderContent(TopRight, ComponentArray(actions...))
		return cardBody
	}
}

// WithBottomRightActions defines the actions to position at the bottom-right of the
// Card body.
func WithBottomRightActions(actions ...templ.Component) Option[*CardBodyProps] {
	return func(cardBody *CardBodyProps) *CardBodyProps {
		cardBody.SetBorderContent(BottomRight, ComponentArray(actions...))
		return cardBody
	}
}

// WithTopLeftActions defines the actions to position at the top-left of the
// Card body.
func WithTopLeftActions(actions ...templ.Component) Option[*CardBodyProps] {
	return func(cardBody *CardBodyProps) *CardBodyProps {
		cardBody.SetBorderContent(TopLeft, ComponentArray(actions...))
		return cardBody
	}
}

// WithBottomLeftActions defines the actions to position at the bottom-left of the
// Card body.
func WithBottomLeftActions(actions ...templ.Component) Option[*CardBodyProps] {
	return func(cardBody *CardBodyProps) *CardBodyProps {
		cardBody.SetBorderContent(BottomLeft, ComponentArray(actions...))
		return cardBody
	}
}

// CardProps represents the properties for a Card.
type CardProps struct {
	border   bool
	layout   CardLayout
	image    *cardImageProps
	centered bool
	compact  bool
	componentAttributes
	htmlAttrID
	modifierShadow
	modifierGlass
	modifierBaseColor
	Title *cardTitleProps
	Body  *CardBodyProps
}

// WithBody defines the content and options for the Card body.
func WithBody(content templ.Component, bodyOptions ...Option[*CardBodyProps]) Option[*CardProps] {
	return func(card *CardProps) *CardProps {
		card.Body = &CardBodyProps{
			Content: content,
		}

		for _, option := range bodyOptions {
			card.Body = option(card.Body)
		}

		return card
	}
}

// WithTitle sets the title for card body.
func WithTitle(title string, size HeaderSize, badges ...templ.Component) Option[*CardProps] {
	return func(card *CardProps) *CardProps {
		card.Title = &cardTitleProps{
			text: title,
			size: size,
		}

		if len(badges) > 0 {
			card.Title.badges = badges
		}

		return card
	}
}

// WithImage sets the figure displayed in the card body.
func WithImage(url string, layout ImageLayout, options ...Option[*ImageProps]) Option[*CardProps] {
	return func(card *CardProps) *CardProps {
		card.image = &cardImageProps{
			image:  Image(url, options...),
			layout: layout,
		}

		return card
	}
}

// WithCardLayout sets the card layout.
func WithCardLayout(l CardLayout) Option[*CardProps] {
	return func(card *CardProps) *CardProps {
		card.layout = l
		return card
	}
}

// WithCenteredContent aligns any image and text content in the center of the
// card with appropriate padding.
func WithCenteredLayout() Option[*CardProps] {
	return func(card *CardProps) *CardProps {
		card.centered = true
		return card
	}
}

// WithBorder sets a border around the card.
func WithBorder() Option[*CardProps] {
	return func(card *CardProps) *CardProps {
		card.border = true
		return card
	}
}

// WithCardShadow adds a shadow to the card.
func WithCardShadow(size ResponsiveSize) Option[*CardProps] {
	return func(card *CardProps) *CardProps {
		card.shadowSize = size
		return card
	}
}

// WithCompactCardBody reduces padding in the card body.
func WithCompactCardBody() Option[*CardProps] {
	return func(card *CardProps) *CardProps {
		card.compact = true
		return card
	}
}

// FromCardProps will set an existing CardProps as the Card properties. This is
// useful when you have built a reusable card with BuildCard and want to render
// it. Pass this option to Card with the CardProps to render.
func FromCardProps(props *CardProps) Option[*CardProps] {
	return func(_ *CardProps) *CardProps {
		return props
	}
}

// BuildCard creates a CardProps with the given options.
func BuildCard(options ...Option[*CardProps]) *CardProps {
	card := &CardProps{
		Body: &CardBodyProps{},
	}

	for _, option := range options {
		card = option(card)
	}

	return card
}
