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

// cardTitle defines a Card's title.
type cardTitle struct {
	text   string
	size   HeaderSize
	badges []Badge
}

// Card represents a DaisyUI Card component.
// https://daisyui.com/components/card/
type Card struct {
	body      templ.Component
	Image     *Image
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
	Badges             []Badge
	title              *cardTitle
}

// WithBody sets the container for card content.
func WithBody(t templ.Component) Option[Card] {
	return func(card Card) Card {
		card.body = t
		return card
	}
}

// WithTopActions defines the actions to be placed on the top-right of the card.
func WithTopRightActions(actions ...templ.Component) Option[Card] {
	return func(card Card) Card {
		card.TopRightActions = append(card.TopRightActions, actions...)
		return card
	}
}

// WithTopActions defines the actions to be placed on the top-right of the card.
func WithBottomRightActions(actions ...templ.Component) Option[Card] {
	return func(card Card) Card {
		card.BottomRightActions = append(card.BottomRightActions, actions...)
		return card
	}
}

// WithBadges adds the provided list of badges to the Card.
func WithBadges(badges ...Badge) Option[Card] {
	return func(card Card) Card {
		card.Badges = append(card.Badges, badges...)
		return card
	}
}

// WithTitle sets the title for card body.
func WithTitle(title string, size HeaderSize, badges ...Badge) Option[Card] {
	return func(card Card) Card {
		card.title = &cardTitle{
			text: title,
			size: size,
		}

		if len(badges) > 0 {
			card.title.badges = badges
		}

		return card
	}
}

// WithImage sets the figure displayed in the card body.
func WithImage(i Image) Option[Card] {
	return func(card Card) Card {
		card.Image = &i
		return card
	}
}

// WithFullImage sets the image as the card background.
func WithFullImage() Option[Card] {
	return func(card Card) Card {
		card.fullImage = true
		return card
	}
}

// WithCardLayout sets the card layout.
func WithCardLayout(l CardLayout) Option[Card] {
	return func(card Card) Card {
		card.layout = l
		return card
	}
}

// WithCenteredContent aligns any image and text content in the center of the
// card with appropriate padding.
func WithCenteredLayout() Option[Card] {
	return func(card Card) Card {
		card.centered = true
		return card
	}
}

// WithBorder sets a border around the card.
func WithBorder() Option[Card] {
	return func(card Card) Card {
		card.border = true
		return card
	}
}

// WithCardShadow adds a shadow to the card.
func WithCardShadow(size ResponsiveSize) Option[Card] {
	return func(card Card) Card {
		card.shadowSize = size
		return card
	}
}

// NewCard creates a new card component with the given options. The card can be
// rendered by calling the Show method.
func NewCard(options ...Option[Card]) Card {
	card := Card{}

	for _, option := range options {
		card = option(card)
	}

	return card
}
