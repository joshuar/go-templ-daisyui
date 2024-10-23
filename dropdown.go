// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: MIT

//go:generate go run golang.org/x/tools/cmd/stringer -type=DropDownOpen,DropDownStyle -linecomment -output dropdown_generated.go
package main

import "strings"

const (
	DropDownTop    DropDownOpen = iota // dropdown-top
	DropDownBottom                     // dropdown-bottom
	DropDownLeft                       // dropdown-left
	DropDownRight                      // dropdown-right
)

type DropDownOpen int

const (
	DropDownButton DropDownStyle = iota
	DropDownCard
	DropDownCustom
)

type DropDownStyle int

type DropDown struct {
	Label string
	class []string
	Style DropDownStyle
}

func (d DropDown) Class() string {
	return strings.Join(d.class, " ")
}

type DropDownOption func(DropDown) DropDown

func WithOpen(o DropDownOpen) DropDownOption {
	return func(d DropDown) DropDown {
		d.class = append(d.class, o.String())
		return d
	}
}

func WithAlignEnd() DropDownOption {
	return func(d DropDown) DropDown {
		d.class = append(d.class, "dropdown-end")
		return d
	}
}

func WithHoverOpen() DropDownOption {
	return func(d DropDown) DropDown {
		d.class = append(d.class, "dropdown-hover")
		return d
	}
}

func WithForceOpen() DropDownOption {
	return func(d DropDown) DropDown {
		d.class = append(d.class, "dropdown-open")
		return d
	}
}

func AsButton(label string) DropDownOption {
	return func(d DropDown) DropDown {
		d.Style = DropDownButton
		d.Label = label

		return d
	}
}

func AsCard() DropDownOption {
	return func(d DropDown) DropDown {
		d.Style = DropDownCard
		return d
	}
}

func AsCustom() DropDownOption {
	return func(d DropDown) DropDown {
		d.Style = DropDownCustom
		return d
	}
}

func NewDropDown(options ...DropDownOption) DropDown {
	dropdown := DropDown{}

	for _, option := range options {
		dropdown = option(dropdown)
	}

	return dropdown
}
