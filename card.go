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
	CardLayoytSideResponsive                   // lg:card-side
)

// CardLayout is used to define the card style.
type CardLayout int

// cardTitle defines a Card's title.
type cardTitle struct {
	text string
	size HeaderSize
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
	componentShadow
	modifierGlass
	Actions []templ.Component
	Badges  []Badge
	title   *cardTitle
}

// WithBody sets the container for card content.
func WithBody(t templ.Component) Option[Card] {
	return func(card Card) Card {
		card.body = t
		return card
	}
}

// WithActions sets the container for card actions.
func WithActions(actions ...templ.Component) Option[Card] {
	return func(card Card) Card {
		card.Actions = append(card.Actions, actions...)
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
func WithTitle(title string, size HeaderSize) Option[Card] {
	return func(card Card) Card {
		card.title = &cardTitle{
			text: title,
			size: size,
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
func WithCardShadow(size Size) Option[Card] {
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
