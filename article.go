// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	MIT

package components

const (
	Gray    ProseGrayScale = iota + 1 // prose-gray
	Slate                             // prose-slate
	Zinc                              // prose-zinc
	Neutral                           // prose-neutral
	Stone                             // prose-stone
)

type ProseGrayScale int

type ArticleProps struct {
	componentClasses
	grayscale ProseGrayScale
}

// WithGrayScale adds optional gray scale scheme to the prose.
// https://github.com/tailwindlabs/tailwindcss-typography/blob/main/README.md#choosing-a-gray-scale
func WithGrayScale(scale ProseGrayScale) Option[ArticleProps] {
	return func(a ArticleProps) ArticleProps {
		a.grayscale = scale
		return a
	}
}

// BuildArticle creates a new article component with the given title, content and
// options. Title is optional.
func BuildArticle(options ...Option[ArticleProps]) ArticleProps {
	article := ArticleProps{}

	for _, option := range options {
		article = option(article)
	}

	return article
}
