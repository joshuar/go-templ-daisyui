// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

// Color represents a DaisyUI color.
//
// https://daisyui.com/docs/colors/
package color

const (
	InvalidMinColor ThemeColor = iota
	Neutral
	Primary
	Secondary
	Accent
	Ghost
	InvalidMaxColor
)

// ThemeColor represents a theme color.
type ThemeColor int

// ThemeColorProps contains properties for setting a theme color.
type ThemeColorProps struct {
	theme   ThemeColor
	outline bool
}

func (p *ThemeColorProps) SetThemeColor(themeColor ThemeColor) {
	p.theme = themeColor
}

func (p *ThemeColorProps) SetOutline(value bool) {
	p.outline = value
}

// ValidThemeColor returns true if the color is a valid color value.
func (p *ThemeColorProps) ValidThemeColor() bool {
	return p.theme > InvalidMinColor && p.theme < InvalidMaxColor
}

func (p *ThemeColorProps) GetThemeColor() ThemeColor {
	return p.theme
}

func (p *ThemeColorProps) OutlineTheme() bool {
	return p.outline
}

func NewThemeColor(color ThemeColor, outline bool) *ThemeColorProps {
	return &ThemeColorProps{
		theme:   color,
		outline: outline,
	}
}
