// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package components

import (
	"github.com/a-h/templ"
)

const (
	CardLayoutNormal         CardLayout = iota // card-normal
	CardLayoutCompact                          // card-compact
	CardLayoutSide                             // card-side
	CardLayoutSideResponsive                   // lg:card-side
)

// CardLayout is used to define the card style.
type CardLayout int

// cardTitleProps defines the properties for a Card title.
type cardTitleProps struct {
	text   string
	size   HeaderSize
	badges []templ.Component
}

type cardBodyProps struct {
	content templ.Component
	componentAttributes
}

// CardProps represents the properties for a Card.
type CardProps struct {
	Image     templ.Component
	border    bool
	layout    CardLayout
	fullImage bool
	centered  bool
	componentAttributes
	componentID
	modifierShadow
	modifierGlass
	modifierBaseColor
	TopRightActions    []templ.Component
	BottomRightActions []templ.Component
	Badges             []templ.Component
	Title              *cardTitleProps
	Body               cardBodyProps
}

// WithBody sets the container for card content.
func WithBody(content templ.Component, attributes templ.Attributes) Option[CardProps] {
	return func(card CardProps) CardProps {
		if content != nil {
			card.Body = cardBodyProps{
				content: content,
			}
			card.Body.AddAttributes(attributes)
		}
		return card
	}
}

// WithTopActions defines the actions to be placed on the top-right of the card.
func WithTopRightActions(actions ...templ.Component) Option[CardProps] {
	return func(card CardProps) CardProps {
		card.TopRightActions = append(card.TopRightActions, actions...)
		return card
	}
}

// WithTopActions defines the actions to be placed on the top-right of the card.
func WithBottomRightActions(actions ...templ.Component) Option[CardProps] {
	return func(card CardProps) CardProps {
		card.BottomRightActions = append(card.BottomRightActions, actions...)
		return card
	}
}

// WithBadges adds the provided list of badges to the Card.
func WithBadges(badges ...templ.Component) Option[CardProps] {
	return func(card CardProps) CardProps {
		card.Badges = append(card.Badges, badges...)
		return card
	}
}

// WithTitle sets the title for card body.
func WithTitle(title string, size HeaderSize, badges ...templ.Component) Option[CardProps] {
	return func(card CardProps) CardProps {
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
func WithImage(url string, options ...Option[ImageProps]) Option[CardProps] {
	return func(card CardProps) CardProps {
		card.Image = Image(url, options...)
		return card
	}
}

// WithFullImage sets the image as the card background.
func WithFullImage() Option[CardProps] {
	return func(card CardProps) CardProps {
		card.fullImage = true
		return card
	}
}

// WithCardLayout sets the card layout.
func WithCardLayout(l CardLayout) Option[CardProps] {
	return func(card CardProps) CardProps {
		card.layout = l
		return card
	}
}

// WithCenteredContent aligns any image and text content in the center of the
// card with appropriate padding.
func WithCenteredLayout() Option[CardProps] {
	return func(card CardProps) CardProps {
		card.centered = true
		return card
	}
}

// WithBorder sets a border around the card.
func WithBorder() Option[CardProps] {
	return func(card CardProps) CardProps {
		card.border = true
		return card
	}
}

// WithCardShadow adds a shadow to the card.
func WithCardShadow(size ResponsiveSize) Option[CardProps] {
	return func(card CardProps) CardProps {
		card.shadowSize = size
		return card
	}
}

// FromCardProps will set an existing CardProps as the Card properties. This is
// useful when you have built a reusable card with BuildCard and want to render
// it. Pass this option to Card with the CardProps to render.
func FromCardProps(props CardProps) Option[CardProps] {
	return func(cp CardProps) CardProps {
		return props
	}
}

// BuildCard creates a CardProps with the given options.
func BuildCard(options ...Option[CardProps]) CardProps {
	card := CardProps{}

	for _, option := range options {
		card = option(card)
	}

	return card
}
