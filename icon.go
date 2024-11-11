// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

//go:generate go run golang.org/x/tools/cmd/stringer -type=FAStyle -linecomment -output icon_generated.go
package components

import "strings"

const (
	Solid FAStyle = iota // fa-solid
)

type FAStyle int

type Icon struct {
	Label string
	class []string
	style FAStyle
}

func (i Icon) Class() string {
	i.class = append(i.class, i.style.String())
	return strings.Join(i.class, " ")
}

type IconOption func(Icon) Icon

func WithStyle(s FAStyle) IconOption {
	return func(i Icon) Icon {
		i.style = s
		return i
	}
}

func WithColor(c Color) IconOption {
	return func(i Icon) Icon {
		i.class = append(i.class, c.String())
		return i
	}
}

func WithLabel(l string) IconOption {
	return func(i Icon) Icon {
		i.Label = l
		return i
	}
}

func WithExtraIconClasses(classes ...string) IconOption {
	return func(i Icon) Icon {
		i.class = append(i.class, classes...)
		return i
	}
}

func NewIcon(name string, options ...IconOption) Icon {
	icon := Icon{
		class: []string{"fa-" + name},
	}

	for _, option := range options {
		icon = option(icon)
	}

	return icon
}
