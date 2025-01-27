// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	MIT

package components

const (
	TextBase TextSize = iota
	TextSM
	TextLG
	TextXL
	Text2XL
)

const (
	TextNormal TextWeight = iota
	TextThin
	TextExtralight
	TextLight
	TextMedium
	TextSemibold
	TextBold
	TextExtrabold
	TextBlack
)

type (
	// Text Size
	//
	// https://tailwindcss.com/docs/font-size
	TextSize int
	// Text Weight
	//
	// https://tailwindcss.com/docs/font-weight
	TextWeight int
)

type TextProps struct {
	size   TextSize
	weight TextWeight
	italic bool
	modifierColor
}

// WithTextSize option sets the font size.
//
// https://tailwindcss.com/docs/font-size
func WithTextSize(size TextSize) Option[*TextProps] {
	return func(text *TextProps) *TextProps {
		text.size = size
		return text
	}
}

// WithTextWeight option sets the font weight.
//
// https://tailwindcss.com/docs/font-weight
func WithTextWeight(weight TextWeight) Option[*TextProps] {
	return func(text *TextProps) *TextProps {
		text.weight = weight
		return text
	}
}

// AsItalicText option will render the text as italic.
//
// https://tailwindcss.com/docs/font-style
func AsItalicText() Option[*TextProps] {
	return func(text *TextProps) *TextProps {
		text.italic = true
		return text
	}
}

func BuildText(options ...Option[*TextProps]) *TextProps {
	text := &TextProps{}

	for _, option := range options {
		text = option(text)
	}

	return text
}
